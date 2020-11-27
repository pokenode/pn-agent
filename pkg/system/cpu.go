package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUStats struct {
	Cores     int     `json:"cores"`
	MHz       float64 `json:"mhz"`
	ModelName string  `json:"modelName"`
	Percent   float64 `json:"percent"`
}

func CPU(mode string) CPUStats {
	c, err := cpu.Counts(true)
	// Calculate percent over 5min
	var interval int64 = 300
	if mode == "DEV" {
		interval = 0
	}
	p, err := cpu.Percent(time.Duration(interval)*time.Second, false)
	cpuInfo, err := cpu.Info()
	if err != nil {
		fmt.Println(err)
	}
	stats := CPUStats{
		Cores:     c,
		MHz:       cpuInfo[0].Mhz,
		ModelName: cpuInfo[0].ModelName,
		Percent:   p[0],
	}

	return stats
}
