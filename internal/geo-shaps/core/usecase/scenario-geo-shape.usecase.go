package usecase

import (
	"bytes"
	"context"
	"digikalajet/internal/geo-shaps/adapter/elasticsearch"
	"digikalajet/internal/geo-shaps/core/domain"
	"digikalajet/internal/geo-shaps/core/ports"
	fake_polygon "digikalajet/pkg/fake-polygon"
	"digikalajet/pkg/monitoring"
	utilsElasticsearch "digikalajet/utils/elasticsearch"
	"digikalajet/utils/viper"
	"fmt"
	"io"
	"log"
	"net/http"
	"sync"
	"time"
)

var (
	scenarioLocalMonitoring   map[string][]domain.LocalMonitoring
	scenarioElasticMonitoring map[string][]domain.ElasticMonitoring
	scenarioElasticLoadTests  map[string][]domain.ElasticLoadTest
	localMonitorings          []domain.LocalMonitoring
	elasticMonitorings        []domain.ElasticMonitoring
	elasticLoadTests          []domain.ElasticLoadTest
)

type ScenarioGeoShapeUsecase struct {
	geoShapeElasticsearch ports.GeoShapeAdapterInterface
	scenarioExcelUsecase  *ScenarioExcelUsecase
}

func NewScenarioGeoShapeUsecase() *ScenarioGeoShapeUsecase {
	return &ScenarioGeoShapeUsecase{
		geoShapeElasticsearch: elasticsearch.NewGeoShapeElasticSearch(),
		scenarioExcelUsecase:  NewScenarioExcelUsecase(),
	}
}

func (u *ScenarioGeoShapeUsecase) ScenarioLoadTest(ctx context.Context) {
	scenarios := viper.C().LoadTest
	scenarioLocalMonitoring = make(map[string][]domain.LocalMonitoring)
	scenarioElasticMonitoring = make(map[string][]domain.ElasticMonitoring)
	scenarioElasticLoadTests = make(map[string][]domain.ElasticLoadTest)

	for name, scenario := range scenarios {
		log.Println("start", name)

		err := u.collectLocalMetrics()
		if err != nil {
			log.Printf("Error during localMetrics collection: %v\n", err)
		}
		err = u.collectElasticMetrics()
		if err != nil {
			log.Printf("Error during localMetrics collection: %v\n", err)
		}
		time.Sleep(10 * time.Second)

		u.runLoadTest(ctx, domain.ScenarioLoadTest{
			WholeTime:    scenario.WholeTime,
			Interval:     scenario.Interval,
			RequestCount: scenario.RequestCount,
		})

		time.Sleep(10 * time.Second)
		err = u.collectLocalMetrics()
		if err != nil {
			log.Printf("Error during localMetrics collection: %v\n", err)
		}
		err = u.collectElasticMetrics()
		if err != nil {
			log.Printf("Error during localMetrics collection: %v\n", err)
		}

		scenarioLocalMonitoring[name] = localMonitorings
		scenarioElasticMonitoring[name] = elasticMonitorings
		scenarioElasticLoadTests[name] = elasticLoadTests
		localMonitorings = []domain.LocalMonitoring{}
		elasticMonitorings = []domain.ElasticMonitoring{}
		elasticLoadTests = []domain.ElasticLoadTest{}

		log.Println("end", name)
		//time.Sleep(time.Minute)
	}
	for name := range scenarios {
		u.scenarioExcelUsecase.GenerateLocalMonitoringExcel(name, scenarioLocalMonitoring[name])
		u.scenarioExcelUsecase.GenerateElasticMonitoringExcel(name, scenarioElasticMonitoring[name])
		u.scenarioExcelUsecase.GenerateElasticLoadTestExcel(name, scenarioElasticLoadTests[name])
	}
	log.Println("DONE!")
}

func (u *ScenarioGeoShapeUsecase) runLoadTest(ctx context.Context, scenario domain.ScenarioLoadTest) {
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(3)

	go u.monitoring("local", ctx, &wg, u.collectLocalMetrics)
	go u.monitoring("elastic", ctx, &wg, u.collectElasticMetrics)
	go u.elasticLoadTest(ctx, scenario, cancel, &wg)

	wg.Wait()
}

func (u *ScenarioGeoShapeUsecase) collectLocalMetrics() error {
	m, err := monitoring.Run()
	if err != nil {
		return err
	}
	localMonitorings = append(localMonitorings, domain.LocalMonitoring{
		Timestamp:   time.Now().Format("2006-01-02 15:04:05.000000"),
		CpuUsage:    m.CpuUsage,
		MemoryUsage: m.MemoryUsage,
		DiskUsage:   m.DiskUsage,
		NetSent:     m.NetSent,
		NetRecv:     m.NetRecv,
	})
	return nil
}

func (u *ScenarioGeoShapeUsecase) collectElasticMetrics() error {
	usage, err := u.geoShapeElasticsearch.GetSystemResourceUsage(context.Background())
	if err != nil {
		return err
	}

	for node, data := range usage.Nodes {
		for _, threadPoolFs := range data.Fs.Data {
			elasticMonitorings = append(elasticMonitorings, domain.ElasticMonitoring{
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

func (u *ScenarioGeoShapeUsecase) monitoring(name string, ctx context.Context, wg *sync.WaitGroup, collectMetrics func() error) {
	defer wg.Done()

	ticker := time.NewTicker(1 * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if err := collectMetrics(); err != nil {
				log.Printf("Error during %sMetrics collection: %v\n", name, err)
				continue
			}
		case <-ctx.Done():
			return
		}
	}
}

func (u *ScenarioGeoShapeUsecase) elasticLoadTest(ctx context.Context, scenario domain.ScenarioLoadTest, cancel context.CancelFunc, wg *sync.WaitGroup) {
	defer wg.Done()
	defer cancel()

	var (
		ticker    = time.NewTicker(time.Second * time.Duration(scenario.Interval))
		endTicker = time.NewTicker(time.Second * time.Duration(scenario.WholeTime))
	)
	defer ticker.Stop()
	defer endTicker.Stop()
	count := 0

	for {
		select {
		case <-ticker.C:
			count = u.startLoadTest(ctx, scenario, count)
		case <-endTicker.C:
			return
		}
	}
}

func (u *ScenarioGeoShapeUsecase) startLoadTest(ctx context.Context, scenario domain.ScenarioLoadTest, count int) int {
	var wg sync.WaitGroup
	for i := 0; i < scenario.RequestCount; i++ {
		count++
		wg.Add(1)
		go func(j int) {
			err := u.clearElasticsearchCache("geo-shaps-v1")
			if err != nil {
				log.Println(err)
				wg.Done()
				return
			}

			startReq := time.Now()
			res, err := u.geoShapeElasticsearch.FindByQuery(ctx, utilsElasticsearch.Search{
				Query: utilsElasticsearch.SearchQuery{
					GeoShape: utilsElasticsearch.SearchQueryGeoShape{
						Location: utilsElasticsearch.SearchQueryGeoShapeLocation{
							Shape: utilsElasticsearch.SearchQueryGeoShapeLocationShape{
								Type:        "point",
								Coordinates: fake_polygon.GenerateRandomLatLng(),
							},
							Relation: "intersects",
						},
					},
				},
			})

			if err != nil {
				log.Println(err)
			}
			endReq := time.Now()
			elasticLoadTests = append(elasticLoadTests, domain.ElasticLoadTest{
				Start:         startReq,
				End:           endReq,
				RequestNumber: j,
				Status:        err == nil,
			})

			if endReq.Sub(startReq).Milliseconds() == 0 {
				log.Println(j, res, endReq.Sub(startReq).Milliseconds())
			}

			wg.Done()
		}(count)
		time.Sleep(time.Millisecond)
	}
	wg.Wait()
	return count
}
func (u *ScenarioGeoShapeUsecase) clearElasticsearchCache(indexName string) error {
	url := fmt.Sprintf(
		"http://&s:%d/%s/_cache/clear",
		viper.C().Database.ElasticSearch.Address,
		viper.C().Database.ElasticSearch.Port,
		indexName,
	)

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(nil))
	if err != nil {
		return fmt.Errorf("error creating request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending request: %v", err)
	}
	defer func() {
		_ = resp.Body.Close()
	}()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return fmt.Errorf("error clearing cache: %s", body)
	}

	fmt.Println("Cache cleared successfully!")
	return nil
}
