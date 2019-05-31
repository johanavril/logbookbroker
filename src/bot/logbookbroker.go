package bot

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jasonlvhit/gocron"
	"github.com/johanavril/logbookbroker/src/service"
	"github.com/line/line-bot-sdk-go/linebot"
)

type logbookBroker struct {
	bot          *linebot.Client
	channelToken string
}

func New(channelSecret, channelToken string) (*logbookBroker, error) {
	bot, err := linebot.New(channelSecret, channelToken)
	if err != nil {
		return nil, err
	}

	return &logbookBroker{
		bot:          bot,
		channelToken: channelToken,
	}, nil
}

func (app *logbookBroker) Callback(w http.ResponseWriter, r *http.Request) {
	events, err := app.bot.ParseRequest(r)
	if err != nil {
		w.WriteHeader(500)
		return
	}

	for _, event := range events {
		switch event.Type {
		case linebot.EventTypeMessage:
			switch message := event.Message.(type) {
			case *linebot.TextMessage:
				if err := app.handleText(message, event.ReplyToken, event.Source); err != nil {
					log.Fatal(err)
				}
			default:
				log.Println("Message not supported ", message)
			}
		default:
			log.Println("Event not supported ", event)
		}
	}
}

func (app *logbookBroker) RegisterCron() {
	fmt.Println("Registerin cron")
	gocron.Every(1).Day().At("07:31").Do(service.SubmitReminder, app.channelToken)
	gocron.Every(1).Friday().At("02:00").Do(service.RequestEditReminder, app.channelToken)
	fmt.Println("Cron successfully registered")

	<-gocron.Start()
}
