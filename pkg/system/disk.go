package system

import (
	"fmt"

	"github.com/shirou/gopsutil/disk"
)

type DiskStats struct {
	FSType      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func Disk() DiskStats {
	diskStats, err := disk.Usage("/")
	if err != nil {
		fmt.Println(err)
	}
	stats := DiskStats{
		FSType:      diskStats.Fstype,
		Total:       diskStats.Total,
		Free:        diskStats.Free,
		Used:        diskStats.Used,
		UsedPercent: diskStats.UsedPercent,
	}
	return stats
}
