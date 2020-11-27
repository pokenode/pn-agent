package system

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/net"
)

type NetStats struct {
	Name      string `json:"name"`
	BytesSent uint64 `json:"bytes_sent"`
	BytesRecv uint64 `json:"bytes_recv"`
}

func Net() NetStats {
	counters, err := net.IOCounters(true)
	if err != nil {
		fmt.Println(err)
	}
	var counter net.IOCountersStat
	for _, c := range counters {
		if c.Name[0] == 'e' {
			counter = c
			break
		}
	}
	stats := NetStats{
		Name:      counter.Name,
		BytesSent: counter.BytesSent,
		BytesRecv: counter.BytesRecv,
	}
	return stats
}
