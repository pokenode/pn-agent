package system

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

type HostStats struct {
	Hostname   string `json:"hostname"`
	Uptime     uint64 `json:"uptime"`
	Procs      uint64 `json:"procs"`
	OS         string `json:"os"`
	Platform   string `json:"platform"`
	KernelArch string `json:"kernel_arch"`
	VirtSys    string `json:"virt_sys"`
}

func Host() HostStats {
	hostStats, err := host.Info()
	if err != nil {
		fmt.Println(err)
	}
	stats := HostStats{
		Hostname:   hostStats.Hostname,
		Uptime:     hostStats.Uptime,
		Procs:      hostStats.Procs,
		OS:         hostStats.OS,
		Platform:   hostStats.Platform,
		KernelArch: hostStats.KernelArch,
		VirtSys:    hostStats.VirtualizationSystem,
	}
	return stats
}
