package system

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

type CPUStats struct {
	Cores     int     `json:"cores"`
	MHz       float64 `json:"mhz"`
	ModelName string  `json:"modelName"`
	Percent   float64 `json:"percent"`
}

func CPU(interval time.Duration) CPUStats {
	c, err := cpu.Counts(true)
	p, err := cpu.Percent(interval, false)
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
