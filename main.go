package main

import (
	"backend/config"
	"backend/jobs"
	"fmt"

	"github.com/robfig/cron/v3"
)

func main() {
	cfg := config.LoadConfig()

	c := cron.New()
	// Run every day at 9 AM

	c.AddFunc("*/5 * * * *", func() {
		fmt.Println("Running test notifier every 5 minutes...")
		jobs.NotifyUsers(cfg)
	})

	c.Start()

	// Keep server alive
	select {}
}
