package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

type DiskStats struct {
	Path        string  `json:"path"`
	FSType      string  `json:"fstype"`
	Total       uint64  `json:"total"`
	Free        uint64  `json:"free"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"used_percent"`
}

func Disk() []DiskStats {
	var sList []DiskStats

	pStats, err := disk.Partitions(false)
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range pStats {
		dStats, err := disk.Usage(p.Mountpoint)
		if err != nil {
			fmt.Println(err)
		}
		stats := DiskStats{
			Path:        dStats.Path,
			FSType:      dStats.Fstype,
			Total:       dStats.Total,
			Free:        dStats.Free,
			Used:        dStats.Used,
			UsedPercent: dStats.UsedPercent,
		}
		sList = append(sList, stats)
	}
	PPrint(sList)

	return sList
}
