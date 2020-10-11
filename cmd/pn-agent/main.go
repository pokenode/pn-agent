package main

import (
	"github.com/robfig/cron/v3"
)

func init() {
	init_env()
}

func main() {
	if MODE == "DEV" {
		SendNodeStats()
	}

	// init
	c := cron.New(cron.WithChain(
		cron.DelayIfStillRunning(cron.DefaultLogger),
	))

	// define jobs
	if MODE == "DEV" {
		c.AddFunc("@every 6s", SendNodeStats)
	} else {
		c.AddFunc("*/5 * * * *", SendNodeStats)
	}

	// run
	c.Run()
}
