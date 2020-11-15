package main

import (
	"github.com/robfig/cron/v3"
)

func init() {
	init_env()
	init_log()
}

func main() {
	// init
	c := cron.New(cron.WithChain(
		cron.DelayIfStillRunning(cron.DefaultLogger),
	))

	// define jobs
	if MODE == "DEV" {
		c.AddFunc("@every 1s", SendNodeStats)
	} else {
		c.AddFunc("@every 5m", SendNodeStats)
	}

	// run first job
	SendNodeStats()

	// run jobs
	c.Run()
}
