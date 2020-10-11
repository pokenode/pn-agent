package system

type Stats struct {
	CPU    CPUStats
	Mem    MemStats
	Disk   DiskStats
	Net    NetStats
	Host   HostStats
	Proc   []ProcStats
	Docker []DockerStats
}

func GetStats() Stats {
	stats := Stats{
		CPU:    CPU(),
		Mem:    Mem(),
		Disk:   Disk(),
		Net:    Net(),
		Host:   Host(),
		Proc:   Proc(),
		Docker: Docker(),
	}
	return stats
}
