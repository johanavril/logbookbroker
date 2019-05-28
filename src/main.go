package main

import (
	"log"
	"net/http"
	"os"

	"github.com/johanavril/logbookbroker/src/bot"
)

func main() {
	app, err := bot.New(
		os.Getenv("CHANNEL_SECRET"),
		os.Getenv("CHANNEL_TOKEN"),
	)
	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/app", app.Callback)

	go app.RegisterCron()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatal(err)
	}
}
