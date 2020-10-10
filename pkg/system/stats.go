package system

type Stats struct {
	Host HostStats
	CPU  CPUStats
	Mem  MemStats
	Disk DiskStats
	Net  NetStats
}

func GetStats() Stats {
	stats := Stats{
		Host: Host(),
		CPU:  CPU(),
		Mem:  Mem(),
		Disk: Disk(),
		Net:  Net(),
	}
	return stats
}
