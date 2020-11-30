package system

type Stats struct {
	CPU    CPUStats
	Mem    MemStats
	Net    NetStats
	Host   HostStats
	Disk   []DiskStats
	Proc   []ProcStats
	Docker []DockerStats
	Lat    LatStats
}

func GetStats(mode string) Stats {
	stats := Stats{
		CPU:    CPU(mode),
		Mem:    Mem(),
		Net:    Net(),
		Host:   Host(),
		Disk:   Disk(),
		Proc:   Proc(),
		Docker: Docker(),
		Lat:    Lat(),
	}
	//PPrint(stats)
	return stats
}
