package main

import (
	"log"
	"net/http"

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

	app.RegisterCron()

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
