package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/disk"
)

type DiskStats struct {
	Device      string  `json:"device"`
	MountPoint  string  `json:"mount_point"`
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
	//PPrint(pStats)

	for _, p := range pStats {
		// Filter out bind mount
		if isBindMount(p) {
			continue
		}
		// Add to stats list
		dStats, err := disk.Usage(p.Mountpoint)
		if err != nil {
			fmt.Println(err)
		}
		stats := DiskStats{
			Device:      p.Device,
			MountPoint:  p.Mountpoint,
			FSType:      p.Fstype,
			Total:       dStats.Total,
			Free:        dStats.Free,
			Used:        dStats.Used,
			UsedPercent: dStats.UsedPercent,
		}
		sList = append(sList, stats)
	}

	//PPrint(sList)
	return sList
}

func isBindMount(p disk.PartitionStat) bool {
	for _, o := range p.Opts {
		if o == "bind" {
			return true
		}
	}
	return false
}
