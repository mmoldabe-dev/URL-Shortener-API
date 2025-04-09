package main

import (
	"URL-Shortener-API/config"
	"URL-Shortener-API/db"
	"log"

	"github.com/robfig/cron/v3"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error read config")
	}

	DB := db.InitDB(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)
	defer DB.Close()

	c := cron.New()

	_, err = c.AddFunc("@hourly", func() {
		db.DeleteExpiredRecords(DB)
	})

	if err != nil {
		log.Fatal("Error adding cron job:", err)
	}

	c.Start()

	select {}

}
