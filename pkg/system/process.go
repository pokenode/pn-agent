package system

import (
	"encoding/json"
	"fmt"
	"sort"

	"github.com/shirou/gopsutil/process"
)

type ProcStats struct {
	Name    string  `json:"name"`
	Percent float64 `json:"percent"`
}

func Proc() string {
	var str string
	var pList []ProcStats

	procs, err := process.Processes()
	if err != nil {
		fmt.Println(err)
	}
	for _, p := range procs {
		name, _ := p.Name()
		percent, _ := p.CPUPercent()
		stats := ProcStats{
			Name:    name,
			Percent: percent,
		}
		pList = append(pList, stats)
	}
	// Sort proc list
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].Percent > pList[j].Percent
	})
	// Keep top 10
	pList = pList[:10]

	b, err := json.Marshal(pList)
	if err != nil {
		fmt.Println(err)
	}
	str = string(b)
	return str
}
