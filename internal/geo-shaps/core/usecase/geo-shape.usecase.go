package usecase

import (
	"context"
	"digikalajet/internal/geo-shaps/adapter/elasticsearch"
	"digikalajet/internal/geo-shaps/core/domain"
	"digikalajet/internal/geo-shaps/core/ports"
	fake_polygon "digikalajet/pkg/fake-polygon"
	"digikalajet/pkg/monitoring"
	utilsElasticsearch "digikalajet/utils/elasticsearch"
	"fmt"
	"log"
	"math/rand"
	"sync"
	"time"
)

var (
	localMonitoring   []domain.LocalMonitoring
	elasticMonitoring []domain.ElasticMonitoring
	elasticLoadTest   []domain.ElasticLoadTest
)

type GeoShapeUsecase struct {
	geoShapeElasticsearch ports.GeoShapeAdapterInterface
	excelUsecase          *ExcelUsecase
}

func NewGeoShapeUsecase() *GeoShapeUsecase {
	return &GeoShapeUsecase{
		geoShapeElasticsearch: elasticsearch.NewGeoShapeElasticSearch(),
		excelUsecase:          NewExcelUsecase(),
	}
}

func (u *GeoShapeUsecase) MakeIndex(ctx context.Context) {
	u.geoShapeElasticsearch.MakeGeoShapeV1Index(ctx)
}

func (u *GeoShapeUsecase) MakeFakePolygons(_ context.Context, loadTest domain.FakePolygon) {
	fake_polygon.Make(loadTest.Counter)
}

func (u *GeoShapeUsecase) InsertFakePolygonsToElastic(ctx context.Context) {
	polygons, err := fake_polygon.GetPolygons()
	if err != nil {
		log.Fatal(err)
		return
	}

	var geoShapes []domain.GeoShapeV1Index

	for _, polygon := range polygons {
		rand.Seed(time.Now().UnixNano())

		randomShopIDNumber := rand.Intn(1000) + 1
		randomPolygonIDNumber := rand.Intn(100) + 1
		randomRadiusBaseNumber := rand.Intn(2)
		var radiusBase bool
		if randomRadiusBaseNumber == 1 {
			radiusBase = true
		} else {
			radiusBase = false

		}

		geoShapes = append(geoShapes, domain.GeoShapeV1Index{
			Location: domain.GeoShapeV1IndexLocation{
				Type:        "polygon",
				Coordinates: [][][]float64{polygon.Coordinates},
			},
			ShopID:     uint(randomShopIDNumber),
			PolygonID:  uint(randomPolygonIDNumber),
			RadiusBase: radiusBase,
		})
	}

	err = u.geoShapeElasticsearch.InsertGeoShape(ctx, geoShapes)
	if err != nil {
		log.Println(err)
		return
	}
}

func (u *GeoShapeUsecase) LoadTest(ctx context.Context, loadTest domain.LoadTest) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	err := u.collectLocalMetrics()
	if err != nil {
		fmt.Printf("Error during localMetrics collection: %v\n", err)
	}
	err = u.collectElasticMetrics()
	if err != nil {
		fmt.Printf("Error during localMetrics collection: %v\n", err)
	}
	time.Sleep(10 * time.Second)
	var wg sync.WaitGroup
	wg.Add(3)

	go u.monitoring("local", ctx, &wg, u.collectLocalMetrics)
	go u.monitoring("elastic", ctx, &wg, u.collectElasticMetrics)
	go u.elasticLoadTest(ctx, loadTest, cancel, &wg)

	wg.Wait()
	time.Sleep(10 * time.Second)

	err = u.collectLocalMetrics()
	if err != nil {
		fmt.Printf("Error during localMetrics collection: %v\n", err)
	}
	err = u.collectElasticMetrics()
	if err != nil {
		fmt.Printf("Error during localMetrics collection: %v\n", err)
	}

	wg.Add(3)
	go u.excelUsecase.GenerateLocalMonitoringExcel(localMonitoring, &wg)
	go u.excelUsecase.GenerateElasticMonitoringExcel(elasticMonitoring, &wg)
	go u.excelUsecase.GenerateElasticLoadTestExcel(elasticLoadTest, &wg)
	wg.Wait()
}

func (u *GeoShapeUsecase) collectLocalMetrics() error {
	m, err := monitoring.Run()
	if err != nil {
		return err
	}
	localMonitoring = append(localMonitoring, domain.LocalMonitoring{
		Timestamp:   time.Now().Format("2006-01-02 15:04:05.000000"),
		CpuUsage:    m.CpuUsage,
		MemoryUsage: m.MemoryUsage,
		DiskUsage:   m.DiskUsage,
		NetSent:     m.NetSent,
		NetRecv:     m.NetRecv,
	})
	return nil
}

func (u *GeoShapeUsecase) collectElasticMetrics() error {
	usage, err := u.geoShapeElasticsearch.GetSystemResourceUsage(context.Background())
	if err != nil {
		return err
	}

	for node, data := range usage.Nodes {
		for _, threadPoolFs := range data.Fs.Data {
			elasticMonitoring = append(elasticMonitoring, domain.ElasticMonitoring{
				Timestamp: time.Now().Format("2006-01-02 15:04:05.000000"),
				Node:      node,
				Memory: domain.ElasticMonitoringMemory{
					TotalInBytes: data.Os.Mem.TotalInBytes,
					FreeInBytes:  data.Os.Mem.FreeInBytes,
					UsedInBytes:  data.Os.Mem.UsedInBytes,
					FreePercent:  data.Os.Mem.FreePercent,
					UsedPercent:  data.Os.Mem.UsedPercent,
				},
				Disk: domain.ElasticMonitoringDisk{
					LeastUsedDiskPercent:  data.Fs.LeastUsageEstimate.UsedDiskPercent,
					LeastTotalInBytes:     data.Fs.LeastUsageEstimate.TotalInBytes,
					LeastAvailableInBytes: data.Fs.LeastUsageEstimate.AvailableInBytes,
					MostUsedDiskPercent:   data.Fs.MostUsageEstimate.UsedDiskPercent,
					MostTotalInBytes:      data.Fs.MostUsageEstimate.TotalInBytes,
					MostAvailableInBytes:  data.Fs.MostUsageEstimate.AvailableInBytes,
				},
				ThreadPoolFs: domain.ElasticMonitoringThreadPoolFs{
					TotalInBytes:     threadPoolFs.TotalInBytes,
					FreeInBytes:      threadPoolFs.FreeInBytes,
					AvailableInBytes: threadPoolFs.AvailableInBytes,
				},
				CPU: domain.ElasticMonitoringCPU{
					Percent: data.Os.Cpu.Percent,
				},
				JVM: domain.ElasticMonitoringJVM{
					HeapUsedInBytes:         data.Jvm.Mem.HeapUsedInBytes,
					HeapUsedPercent:         data.Jvm.Mem.HeapUsedPercent,
					HeapCommittedInBytes:    data.Jvm.Mem.HeapCommittedInBytes,
					HeapMaxInBytes:          data.Jvm.Mem.HeapMaxInBytes,
					NonHeapUsedInBytes:      data.Jvm.Mem.NonHeapUsedInBytes,
					NonHeapCommittedInBytes: data.Jvm.Mem.NonHeapCommittedInBytes,
				},
			})
		}
	}

	return nil
}

func (u *GeoShapeUsecase) monitoring(name string, ctx context.Context, wg *sync.WaitGroup, collectMetrics func() error) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := collectMetrics(); err != nil {
				fmt.Printf("Error during %sMetrics collection: %v\n", name, err)
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

func (u *GeoShapeUsecase) elasticLoadTest(ctx context.Context, loadTest domain.LoadTest, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()

	var (
		requestWg sync.WaitGroup
		start     = time.Now()
	)

	for i := 0; i < loadTest.Counter; i++ {
		requestWg.Add(1)
		go func(j int) {
			defer requestWg.Done()
			startReq := time.Now()

			err := u.geoShapeElasticsearch.FindByQuery(ctx, utilsElasticsearch.Search{
				Query: utilsElasticsearch.SearchQuery{
					GeoShape: utilsElasticsearch.SearchQueryGeoShape{
						Location: utilsElasticsearch.SearchQueryGeoShapeLocation{
							Shape: utilsElasticsearch.SearchQueryGeoShapeLocationShape{
								Type:        "point",
								Coordinates: []float64{35.738314, 51.169862},
							},
							Relation: "within",
						},
					},
				},
			})

			if err != nil {
				log.Println(err)
			}
			endReq := time.Now()
			elasticLoadTest = append(elasticLoadTest, domain.ElasticLoadTest{
				Start:         startReq,
				End:           endReq,
				RequestNumber: j,
				Status:        err == nil,
			})
		}(i)
		time.Sleep(time.Millisecond)
	}

	requestWg.Wait()
	elapsed := time.Since(start)
	log.Printf("len elasticLoadTest = %d, elapsed = %s", len(elasticLoadTest), elapsed)

	cancel()
}
