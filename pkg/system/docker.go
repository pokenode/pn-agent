package system

import (
	"fmt"

	"github.com/shirou/gopsutil/docker"
)

type DockerStats struct {
	ContainerID string `json:"container_id"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	Running     bool   `json:"running"`
}

func Docker() []DockerStats {
	var sList []DockerStats

	list, err := docker.GetDockerStat()
	if err != nil {
		fmt.Println(err)
	}
	for _, s := range list {
		stats := DockerStats{
			ContainerID: s.ContainerID[:12],
			Name:        s.Name,
			Status:      s.Status,
			Running:     s.Running,
		}
		sList = append(sList, stats)
	}

	return sList
}