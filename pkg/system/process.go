package system

import (
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/process"
)

type ProcStats struct {
	Name     string  `json:"name"`
	Percent  float64 `json:"percent"`
	Memory   uint64  `json:"memory"`
	Username string  `json:"username"`
}

func Proc() []ProcStats {
	var pList []ProcStats

	procs, err := process.Processes()
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range procs {
		name, _ := p.Name()
		percent, _ := p.CPUPercent()
		username, _ := p.Username()
		memInfo, err := p.MemoryInfo()
		var memory uint64
		if err != nil {
			fmt.Println(err)
			fmt.Println(name)
			memory = 0
		} else {
			memory = memInfo.RSS
		}
		stats := ProcStats{
			Name:     name,
			Percent:  percent,
			Username: username,
			Memory:   memory,
		}
		pList = append(pList, stats)
	}
	// Sort proc list
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].Percent > pList[j].Percent
	})
	// Keep top 10
	pList = pList[:10]

	return pList
}
