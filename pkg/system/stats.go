package system

type Stats struct {
	CPU    CPUStats
	Mem    MemStats
	Disk   []DiskStats
	Net    NetStats
	Host   HostStats
	Proc   []ProcStats
	Docker []DockerStats
}

func GetStats(mode string) Stats {
	stats := Stats{
		CPU:    CPU(mode),
		Mem:    Mem(),
		Disk:   Disk(),
		Net:    Net(),
		Host:   Host(),
		Proc:   Proc(),
		Docker: Docker(),
	}
	//PPrint(stats)
	return stats
}
