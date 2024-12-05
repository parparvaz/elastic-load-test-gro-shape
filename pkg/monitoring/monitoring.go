package monitoring

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

type Monitoring struct {
	CpuUsage    float64
	MemoryUsage float64
	DiskUsage   float64
	NetSent     float64
	NetRecv     float64
}

func monitorCPUUsage() (float64, error) {
	cpuUsage, err := cpu.Percent(0, true)
	if err != nil {
		return 0, err
	}
	return cpuUsage[0], nil
}

func monitorMemoryUsage() (float64, error) {
	vmStat, err := mem.VirtualMemory()
	if err != nil {
		return 0, err
	}
	return vmStat.UsedPercent, nil
}

func monitorDiskUsage() (float64, error) {
	diskUsage, err := disk.Usage("/")
	if err != nil {
		return 0, err
	}
	return diskUsage.UsedPercent, nil
}

func monitorNetworkUsage() (float64, float64, error) {
	netIO, err := net.IOCounters(false)
	if err != nil {
		return 0, 0, err
	}
	if len(netIO) > 0 {
		return float64(netIO[0].BytesSent), float64(netIO[0].BytesRecv), nil
	}
	return 0, 0, nil
}

func Run() (Monitoring, error) {
	var (
		m   Monitoring
		err error
	)

	m.CpuUsage, err = monitorCPUUsage()
	if err != nil {
		return Monitoring{}, err
	}

	m.MemoryUsage, err = monitorMemoryUsage()
	if err != nil {
		return Monitoring{}, err
	}

	m.DiskUsage, err = monitorDiskUsage()
	if err != nil {
		return Monitoring{}, err
	}

	m.NetSent, m.NetRecv, err = monitorNetworkUsage()
	if err != nil {
		return Monitoring{}, err
	}

	return m, nil
}
