package domain

import "time"

type LoadTest struct {
	Counter int
}

type FakePolygon struct {
	Counter int
}

type GeoShapeV1Index struct {
	Location   GeoShapeV1IndexLocation `json:"location"`
	ShopID     uint                    `json:"shop_id"`
	PolygonID  uint                    `json:"polygon_id"`
	RadiusBase bool                    `json:"radius_base"`
}

type GeoShapeV1IndexLocation struct {
	Type        string        `json:"type"`
	Coordinates [][][]float64 `json:"coordinates"`
}

type LocalMonitoring struct {
	Timestamp   string
	CpuUsage    float64
	MemoryUsage float64
	DiskUsage   float64
	NetSent     float64
	NetRecv     float64
}

type ElasticMonitoring struct {
	Timestamp    string
	Node         string
	Memory       ElasticMonitoringMemory
	Disk         ElasticMonitoringDisk
	ThreadPoolFs ElasticMonitoringThreadPoolFs
	CPU          ElasticMonitoringCPU
	JVM          ElasticMonitoringJVM
}

type ElasticMonitoringMemory struct {
	TotalInBytes int64
	FreeInBytes  int
	UsedInBytes  int64
	FreePercent  int
	UsedPercent  int
}

type ElasticMonitoringDisk struct {
	LeastUsedDiskPercent  float64
	LeastTotalInBytes     int64
	LeastAvailableInBytes int64
	MostUsedDiskPercent   float64
	MostTotalInBytes      int64
	MostAvailableInBytes  int64
}

type ElasticMonitoringThreadPoolFs struct {
	TotalInBytes     int64
	FreeInBytes      int64
	AvailableInBytes int64
}

type ElasticMonitoringCPU struct {
	Percent int
}

type ElasticMonitoringJVM struct {
	HeapUsedInBytes         int
	HeapUsedPercent         int
	HeapCommittedInBytes    int
	HeapMaxInBytes          int
	NonHeapUsedInBytes      int
	NonHeapCommittedInBytes int
}

type ElasticLoadTest struct {
	RequestNumber int
	Start         time.Time
	End           time.Time
	Status        bool
}
