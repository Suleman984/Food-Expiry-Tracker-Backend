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

	c.AddFunc("*/1 * * * *", func() {
		fmt.Println("Running test notifier every minute...")
		jobs.NotifyUsers(cfg)
	})

	c.Start()

	select {}
}
