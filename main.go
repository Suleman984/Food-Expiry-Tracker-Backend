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

	c.AddFunc("0 18 * * *", func() {
		fmt.Println("Running test notifier at 6:00 PM...")
		jobs.NotifyUsers(cfg)
	})

	c.Start()

	// Keep server alive
	select {}
}
