package main

import (
	"log"

	"github.com/go-co-op/gocron/v2"
	"github.com/rjhoppe/firekeeper/requests"
)

func main() {
	// Create a new scheduler
	s, err := gocron.NewScheduler()
	if err != nil {
		log.Fatalf("Failed to create scheduler: %v", err)
	}

	// Every Friday at 5:30 PM GET request for Drink of the Day
	_, err = s.NewJob(
		gocron.CronJob("30 17 * * 5", false), // false = use local time
		gocron.NewTask(func() {
			log.Println("Sending scheduled GET request for Drink of the Day...")
			req := requests.GETRequest{Url: "http://localhost:8080/bartender/random"}
			req.Send()
		}),
	)
	if err != nil {
		log.Fatalf("Failed to schedule Friday job: %v", err)
	}

	// Every Saturday at 2:00 PM GET request for dinner recipes
	_, err = s.NewJob(
		gocron.CronJob("0 14 * * 6", false),
		gocron.NewTask(func() {
			log.Println("Sending scheduled GET request for dinner...")
			req := requests.GETRequest{Url: "http://localhost:8080/dinner/random"}
			req.Send()
		}),
	)
	if err != nil {
		log.Fatalf("Failed to schedule Saturday job: %v", err)
	}

	// Start the scheduler
	s.Start()

	log.Println("Scheduler started. Press Ctrl+C to exit.")

	// Keep the program running
	select {}
}
