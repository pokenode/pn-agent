package main

import (
	"github.com/robfig/cron/v3"
)

func init() {
	init_env()
}

func main() {
	SendNodeStats()

	// init
	c := cron.New(cron.WithChain(
		cron.DelayIfStillRunning(cron.DefaultLogger),
	))

	// define jobs
	c.AddFunc("*/5 * * * *", SendNodeStats)

	// run
	c.Run()
}
