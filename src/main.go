package main

import (
	"log"
	"net/http"
	"os"

	"github.com/johanavril/logbookbroker/src/bot"
	"github.com/johanavril/logbookbroker/src/service"
)

func main() {
	credential, err := service.GetAppCredential()
	if err != nil {
		log.Fatal(err)
	}
	app, err := bot.New(
		credential.ChannelSecret,
		credential.ChannelToken,
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
