package main

import (
	"log"
	"net/http"

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
			req := requests.GETRequest{Url: "http://192.168.50.51:8080/bartender/random"}
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
			req := requests.GETRequest{Url: "http://192.168.50.51:8080/dinner/random"}
			req.Send()
		}),
	)
	if err != nil {
		log.Fatalf("Failed to schedule Saturday job: %v", err)
	}

	// Start the scheduler
	s.Start()

	log.Println("Scheduler started. Press Ctrl+C to exit.")

	// Start health check endpoint in a goroutine
	go func() {
		http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(http.StatusOK)
			w.Write([]byte("ok"))
		})
		log.Println("Health check endpoint running on :8081/health")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Health check server failed: %v", err)
		}
	}()

	// Keep the program running
	select {}
}
