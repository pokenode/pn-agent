package system

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type MemStats struct {
	Total           uint64  `json:"total"`
	UsedPercent     float64 `json:"used_percent"`
	SwapTotal       uint64  `json:"swap_total"`
	SwapUsedPercent float64 `json:"swap_used_percent"`
}

func Mem() MemStats {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println(err)
	}
	var used_percent float64
	used_percent = (1 - float64(memStats.Available)/float64(memStats.Total)) * 100
	var swap_used_percent float64
	if memStats.SwapTotal > 0 {
		swap_used_percent = (1 - (float64(memStats.SwapFree)+float64(memStats.SwapCached))/float64(memStats.SwapTotal)) * 100
	}
	stats := MemStats{
		Total:           memStats.Total,
		UsedPercent:     used_percent,
		SwapTotal:       memStats.SwapTotal,
		SwapUsedPercent: swap_used_percent,
	}
	return stats
}
