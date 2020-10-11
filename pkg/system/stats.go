package system

type Stats struct {
	CPU  CPUStats
	Mem  MemStats
	Disk DiskStats
	Net  NetStats
	Host string
	Proc string
}

func GetStats() Stats {
	stats := Stats{
		CPU:  CPU(),
		Mem:  Mem(),
		Disk: Disk(),
		Net:  Net(),
		Host: Host(),
		Proc: Proc(),
	}
	return stats
}
