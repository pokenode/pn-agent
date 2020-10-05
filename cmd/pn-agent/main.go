package main

import (
	sys "github.com/pokenode/pn-agent/pkg/system"
)

func main() {
	cpu := sys.CPU(0)
	PPrint(cpu)
	disk := sys.Disk()
	PPrint(disk)
	host := sys.Host()
	PPrint(host)
	mem := sys.Mem()
	PPrint(mem)
	net := sys.Net()
	PPrint(net)
}
