package elasticsearch

type ClusterNodeStats struct {
	Node        ClusterNodeStatsName             `json:"_nodes"`
	ClusterName string                           `json:"cluster_name"`
	Nodes       map[string]ClusterNodeStatsNodes `json:"nodes"`
}

type ClusterNodeStatsName struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type ClusterNodeStatsNode struct {
	Total      int `json:"total"`
	Successful int `json:"successful"`
	Failed     int `json:"failed"`
}

type ClusterNodeStatsNodes struct {
	Timestamp        int64                             `json:"timestamp"`
	Name             string                            `json:"name"`
	TransportAddress string                            `json:"transport_address"`
	Host             string                            `json:"host"`
	Ip               string                            `json:"ip"`
	Roles            []string                          `json:"roles"`
	Attributes       ClusterNodeStatsNodesAttributes   `json:"attributes"`
	Os               ClusterNodeStatsNodesOS           `json:"os"`
	Process          ClusterNodeStatsNodesProcess      `json:"process"`
	Jvm              ClusterNodeStatsNodesJvm          `json:"jvm"`
	ThreadPool       ClusterNodeStatsNodesThreadPool   `json:"thread_pool"`
	Fs               ClusterNodeStatsNodesThreadPoolFs `json:"fs"`
	Transport        ClusterNodeStatsNodesTransport    `json:"transport"`
}

type ClusterNodeStatsNodesAttributes struct {
	MlMachineMemory string `json:"ml.machine_memory"`
	XpackInstalled  string `json:"xpack.installed"`
	MlMaxOpenJobs   string `json:"ml.max_open_jobs"`
	MlEnabled       string `json:"ml.enabled"`
}

type ClusterNodeStatsNodesOS struct {
	Timestamp int64                       `json:"timestamp"`
	Cpu       ClusterNodeStatsNodesOSCpu  `json:"cpu"`
	Mem       ClusterNodeStatsNodesOSMem  `json:"mem"`
	Swap      ClusterNodeStatsNodesOSSwap `json:"swap"`
}

type ClusterNodeStatsNodesOSCpu struct {
	Percent     int                                   `json:"percent"`
	LoadAverage ClusterNodeStatsNodesOSCpuLoadAverage `json:"load_average"`
}

type ClusterNodeStatsNodesOSCpuLoadAverage struct {
	M  float64 `json:"1m"`
	M1 float64 `json:"5m"`
	M2 float64 `json:"15m"`
}

type ClusterNodeStatsNodesOSMem struct {
	TotalInBytes int64 `json:"total_in_bytes"`
	FreeInBytes  int   `json:"free_in_bytes"`
	UsedInBytes  int64 `json:"used_in_bytes"`
	FreePercent  int   `json:"free_percent"`
	UsedPercent  int   `json:"used_percent"`
}

type ClusterNodeStatsNodesOSSwap struct {
	TotalInBytes int `json:"total_in_bytes"`
	FreeInBytes  int `json:"free_in_bytes"`
	UsedInBytes  int `json:"used_in_bytes"`
}

type ClusterNodeStatsNodesProcess struct {
	Timestamp           int64                           `json:"timestamp"`
	OpenFileDescriptors int                             `json:"open_file_descriptors"`
	MaxFileDescriptors  int                             `json:"max_file_descriptors"`
	Cpu                 ClusterNodeStatsNodesProcessCpu `json:"cpu"`
	Mem                 ClusterNodeStatsNodesProcessMem `json:"mem"`
}

type ClusterNodeStatsNodesProcessCpu struct {
	Percent       int `json:"percent"`
	TotalInMillis int `json:"total_in_millis"`
}

type ClusterNodeStatsNodesProcessMem struct {
	TotalVirtualInBytes int64 `json:"total_virtual_in_bytes"`
}

type ClusterNodeStatsNodesJvm struct {
	Timestamp      int64                               `json:"timestamp"`
	UptimeInMillis int                                 `json:"uptime_in_millis"`
	Mem            ClusterNodeStatsNodesJvmMem         `json:"mem"`
	Threads        ClusterNodeStatsNodesJvmThreads     `json:"threads"`
	Gc             ClusterNodeStatsNodesJvmGc          `json:"gc"`
	BufferPools    ClusterNodeStatsNodesJvmBufferPools `json:"buffer_pools"`
	Classes        ClusterNodeStatsNodesJvmClasses     `json:"classes"`
}

type ClusterNodeStatsNodesJvmMem struct {
	HeapUsedInBytes         int                              `json:"heap_used_in_bytes"`
	HeapUsedPercent         int                              `json:"heap_used_percent"`
	HeapCommittedInBytes    int                              `json:"heap_committed_in_bytes"`
	HeapMaxInBytes          int                              `json:"heap_max_in_bytes"`
	NonHeapUsedInBytes      int                              `json:"non_heap_used_in_bytes"`
	NonHeapCommittedInBytes int                              `json:"non_heap_committed_in_bytes"`
	Pools                   ClusterNodeStatsNodesJvmMemPools `json:"pools"`
}

type ClusterNodeStatsNodesJvmMemPools struct {
	Young    ClusterNodeStatsNodesJvmMemPoolsStats `json:"young"`
	Survivor ClusterNodeStatsNodesJvmMemPoolsStats `json:"survivor"`
	Old      ClusterNodeStatsNodesJvmMemPoolsStats `json:"old"`
}

type ClusterNodeStatsNodesJvmMemPoolsStats struct {
	UsedInBytes     int `json:"used_in_bytes"`
	MaxInBytes      int `json:"max_in_bytes"`
	PeakUsedInBytes int `json:"peak_used_in_bytes"`
	PeakMaxInBytes  int `json:"peak_max_in_bytes"`
}

type ClusterNodeStatsNodesJvmThreads struct {
	Count     int `json:"count"`
	PeakCount int `json:"peak_count"`
}

type ClusterNodeStatsNodesJvmGc struct {
	Collectors ClusterNodeStatsNodesJvmGcCollectors `json:"collectors"`
}

type ClusterNodeStatsNodesJvmGcCollectors struct {
	Young ClusterNodeStatsNodesJvmGcCollectorsStats `json:"young"`
	Old   ClusterNodeStatsNodesJvmGcCollectorsStats `json:"old"`
}

type ClusterNodeStatsNodesJvmGcCollectorsStats struct {
	CollectionCount        int `json:"collection_count"`
	CollectionTimeInMillis int `json:"collection_time_in_millis"`
}

type ClusterNodeStatsNodesJvmBufferPools struct {
	Mapped ClusterNodeStatsNodesJvmBufferPoolsMemoryUsage `json:"mapped"`
	Direct ClusterNodeStatsNodesJvmBufferPoolsMemoryUsage `json:"direct"`
}

type ClusterNodeStatsNodesJvmBufferPoolsMemoryUsage struct {
	Count                int `json:"count"`
	UsedInBytes          int `json:"used_in_bytes"`
	TotalCapacityInBytes int `json:"total_capacity_in_bytes"`
}

type ClusterNodeStatsNodesJvmClasses struct {
	CurrentLoadedCount int `json:"current_loaded_count"`
	TotalLoadedCount   int `json:"total_loaded_count"`
	TotalUnloadedCount int `json:"total_unloaded_count"`
}

type ClusterNodeStatsNodesThreadPool struct {
	Analyze           ClusterNodeStatsNodesThreadPoolQueueStatus `json:"analyze"`
	Ccr               ClusterNodeStatsNodesThreadPoolQueueStatus `json:"ccr"`
	FetchShardStarted ClusterNodeStatsNodesThreadPoolQueueStatus `json:"fetch_shard_started"`
	FetchShardStore   ClusterNodeStatsNodesThreadPoolQueueStatus `json:"fetch_shard_store"`
	Flush             ClusterNodeStatsNodesThreadPoolQueueStatus `json:"flush"`
	ForceMerge        ClusterNodeStatsNodesThreadPoolQueueStatus `json:"force_merge"`
	Generic           ClusterNodeStatsNodesThreadPoolQueueStatus `json:"generic"`
	Get               ClusterNodeStatsNodesThreadPoolQueueStatus `json:"get"`
	Index             ClusterNodeStatsNodesThreadPoolQueueStatus `json:"index"`
	Listener          ClusterNodeStatsNodesThreadPoolQueueStatus `json:"listener"`
	Management        ClusterNodeStatsNodesThreadPoolQueueStatus `json:"management"`
	MlAutodetect      ClusterNodeStatsNodesThreadPoolQueueStatus `json:"ml_autodetect"`
	MlDatafeed        ClusterNodeStatsNodesThreadPoolQueueStatus `json:"ml_datafeed"`
	MlUtility         ClusterNodeStatsNodesThreadPoolQueueStatus `json:"ml_utility"`
	Refresh           ClusterNodeStatsNodesThreadPoolQueueStatus `json:"refresh"`
	RollupIndexing    ClusterNodeStatsNodesThreadPoolQueueStatus `json:"rollup_indexing"`
	Search            ClusterNodeStatsNodesThreadPoolQueueStatus `json:"search"`
	SearchThrottled   ClusterNodeStatsNodesThreadPoolQueueStatus `json:"search_throttled"`
	SecurityTokenKey  ClusterNodeStatsNodesThreadPoolQueueStatus `json:"security-token-key"`
	Snapshot          ClusterNodeStatsNodesThreadPoolQueueStatus `json:"snapshot"`
	Warmer            ClusterNodeStatsNodesThreadPoolQueueStatus `json:"warmer"`
	Watcher           ClusterNodeStatsNodesThreadPoolQueueStatus `json:"watcher"`
	Write             ClusterNodeStatsNodesThreadPoolQueueStatus `json:"write"`
}

type ClusterNodeStatsNodesThreadPoolQueueStatus struct {
	Threads   int `json:"threads"`
	Queue     int `json:"queue"`
	Active    int `json:"active"`
	Rejected  int `json:"rejected"`
	Largest   int `json:"largest"`
	Completed int `json:"completed"`
}

type ClusterNodeStatsNodesThreadPoolFs struct {
	Timestamp          int64                                               `json:"timestamp"`
	Total              ClusterNodeStatsNodesThreadPoolFsLeastTotal         `json:"total"`
	LeastUsageEstimate ClusterNodeStatsNodesThreadPoolFsLeastUsageEstimate `json:"least_usage_estimate"`
	MostUsageEstimate  ClusterNodeStatsNodesThreadPoolFsMostUsageEstimate  `json:"most_usage_estimate"`
	Data               []ClusterNodeStatsNodesThreadPoolFsData             `json:"data"`
	IoStats            ClusterNodeStatsNodesThreadPoolFsIoStats            `json:"io_stats"`
}

type ClusterNodeStatsNodesThreadPoolFsLeastTotal struct {
	TotalInBytes     int64 `json:"total_in_bytes"`
	FreeInBytes      int64 `json:"free_in_bytes"`
	AvailableInBytes int64 `json:"available_in_bytes"`
}

type ClusterNodeStatsNodesThreadPoolFsLeastUsageEstimate struct {
	Path             string  `json:"path"`
	TotalInBytes     int64   `json:"total_in_bytes"`
	AvailableInBytes int64   `json:"available_in_bytes"`
	UsedDiskPercent  float64 `json:"used_disk_percent"`
}

type ClusterNodeStatsNodesThreadPoolFsMostUsageEstimate struct {
	Path             string  `json:"path"`
	TotalInBytes     int64   `json:"total_in_bytes"`
	AvailableInBytes int64   `json:"available_in_bytes"`
	UsedDiskPercent  float64 `json:"used_disk_percent"`
}

type ClusterNodeStatsNodesThreadPoolFsData struct {
	Path             string `json:"path"`
	Mount            string `json:"mount"`
	Type             string `json:"type"`
	TotalInBytes     int64  `json:"total_in_bytes"`
	FreeInBytes      int64  `json:"free_in_bytes"`
	AvailableInBytes int64  `json:"available_in_bytes"`
}

type ClusterNodeStatsNodesThreadPoolFsIoStats struct {
	Devices []ClusterNodeStatsNodesThreadPoolFsIoStatsDevices `json:"devices"`
	Total   ClusterNodeStatsNodesThreadPoolFsIoStatsTotal     `json:"total"`
}

type ClusterNodeStatsNodesThreadPoolFsIoStatsDevices struct {
	DeviceName      string `json:"device_name"`
	Operations      int    `json:"operations"`
	ReadOperations  int    `json:"read_operations"`
	WriteOperations int    `json:"write_operations"`
	ReadKilobytes   int    `json:"read_kilobytes"`
	WriteKilobytes  int    `json:"write_kilobytes"`
}

type ClusterNodeStatsNodesThreadPoolFsIoStatsTotal struct {
	Operations      int `json:"operations"`
	ReadOperations  int `json:"read_operations"`
	WriteOperations int `json:"write_operations"`
	ReadKilobytes   int `json:"read_kilobytes"`
	WriteKilobytes  int `json:"write_kilobytes"`
}

type ClusterNodeStatsNodesTransport struct {
	ServerOpen    int `json:"server_open"`
	RxCount       int `json:"rx_count"`
	RxSizeInBytes int `json:"rx_size_in_bytes"`
	TxCount       int `json:"tx_count"`
	TxSizeInBytes int `json:"tx_size_in_bytes"`
}
