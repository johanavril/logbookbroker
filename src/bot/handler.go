package bot

import (
	"bytes"
	"net/http"
	"net/url"
	"strings"
	"text/template"
	"time"

	"github.com/johanavril/logbookbroker/src/constant"
	"github.com/johanavril/logbookbroker/src/model"
	"github.com/johanavril/logbookbroker/src/service"
	"github.com/johanavril/logbookbroker/src/util"
	"github.com/line/line-bot-sdk-go/linebot"
)

func getUserInput(message string) (string, []string) {
	messageSlice := strings.Split(message, "\n")
	command := messageSlice[0]
	input := messageSlice[1:]

	return command, input
}

func (app *logbookBroker) help(replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(constant.Message.Help),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *logbookBroker) helpList(replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(constant.Message.HelpList),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *logbookBroker) helpLogbook(replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(constant.Message.HelpLogbook),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *logbookBroker) helpLogin(replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(constant.Message.HelpLogin),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *logbookBroker) helpStatus(replyToken string) error {
	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewTextMessage(constant.Message.HelpStatus),
	).Do(); err != nil {
		return err
	}
	return nil
}

func (app *logbookBroker) getUserData(replyToken, userId string) (*model.User, error) {
	user := &model.User{}
	if err := user.GetByLineUserId(userId); err != nil {
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(constant.Message.NotLoggedIn),
		).Do(); err != nil {
			return nil, err
		}
		return nil, nil
	}

	return user, nil
}

func (app *logbookBroker) login(input []string) (bool, error) {
	cookies, csrfToken, err := service.GetCSRF(constant.URL.Login)
	if err != nil {
		return false, err
	}

	body := url.Values{}
	body.Set("_token", csrfToken)
	body.Set("username", input[0])
	body.Set("password", input[1])
	req, err := http.NewRequest(http.MethodPost, constant.URL.Login, strings.NewReader(body.Encode()))
	if err != nil {
		return false, err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("cache-control", "no-cache")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}

	resp, err := service.GetLogbookClient().Do(req)
	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	if title, err := service.GetPageTitle(resp.Body); err != nil {
		return false, err
	} else if strings.HasSuffix(title, "Login") {
		return false, nil
	}

	return true, nil
}

func (app *logbookBroker) loginAndRecord(input []string, replyToken, userId string) error {
	if len(input) != 2 {
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(constant.Message.LoginFormat),
		).Do(); err != nil {
			return err
		}
		return nil
	}

	success, err := app.login(input)

	if success {
		name, err := service.GetName()
		if err != nil {
			return err
		}

		user := &model.User{
			UserId:   userId,
			Username: input[0],
			Password: input[1],
			Name:     name,
		}
		if err := user.Upsert(); err != nil {
			return err
		}

		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(constant.Message.LoginSuccess+" "+name+"."),
		).Do(); err != nil {
			return err
		}
	} else {
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(constant.Message.LoginFailed),
		).Do(); err != nil {
			return err
		}
	}

	return err
}

func (app *logbookBroker) logbook(input []string, replyToken, userId string) error {
	if len(input) != 4 {
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(constant.Message.LogbookFormat),
		).Do(); err != nil {
			return err
		}
		return nil
	}

	user, err := app.getUserData(replyToken, userId)
	if err != nil {
		return err
	} else if user == nil {
		return nil
	}

	decryptedPassword, err := util.Decrypt(user.Password)
	if err != nil {
		return err
	}

	userCredential := []string{user.Username, decryptedPassword}
	if _, err := app.login(userCredential); err != nil {
		return err
	}

	cookies, csrfToken, err := service.GetCSRF(constant.URL.LogbookInsert)
	if err != nil {
		return err
	}

	body := url.Values{}
	body.Set("_token", csrfToken)
	body.Set("clock-in", input[0])
	body.Set("clock-out", input[1])
	body.Set("activity", input[2])
	body.Set("description", input[3])

	req, err := http.NewRequest(http.MethodPost, constant.URL.LogbookInsert, strings.NewReader(body.Encode()))
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("cache-control", "no-cache")
	for _, cookie := range cookies {
		req.AddCookie(cookie)
	}
	resp, err := service.GetLogbookClient().Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()

	if resp.StatusCode == http.StatusOK {

		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(constant.Message.LogbookSuccess),
		).Do(); err != nil {
			return err
		}
	} else {
		if _, err := app.bot.ReplyMessage(
			replyToken,
			linebot.NewTextMessage(constant.Message.LogbookFailed),
		).Do(); err != nil {
			return err
		}
	}

	return nil
}

func (app *logbookBroker) status(replyToken, userId string) error {
	user, err := app.getUserData(replyToken, userId)
	if err != nil {
		return err
	} else if user == nil {
		return nil
	}

	decryptedPassword, err := util.Decrypt(user.Password)
	if err != nil {
		return err
	}

	userCredential := []string{user.Username, decryptedPassword}
	if _, err := app.login(userCredential); err != nil {
		return err
	}

	loc, _ := time.LoadLocation("Asia/Jakarta")
	now := time.Now().In(loc)
	logbook, err := service.GetLogbook(now)
	if err != nil {
		return err
	}

	tmpl := template.Must(template.ParseFiles("../template/logbook-status.tmpl"))
	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, logbook); err != nil {
		return err
	}

	content, err := linebot.UnmarshalFlexMessageJSON([]byte(buf.String()))
	if err != nil {
		return err
	}

	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewFlexMessage("Logbook Status", content),
	).Do(); err != nil {
		return err
	}

	return nil
}

func (app *logbookBroker) list(replyToken, userId string) error {
	user, err := app.getUserData(replyToken, userId)
	if err != nil {
		return err
	} else if user == nil {
		return nil
	}

	decryptedPassword, err := util.Decrypt(user.Password)
	if err != nil {
		return err
	}

	userCredential := []string{user.Username, decryptedPassword}
	if _, err := app.login(userCredential); err != nil {
		return err
	}

	logbooks, err := service.GetWeekLogbook()
	if err != nil {
		return err
	}

	funcMap := template.FuncMap{
		"minus": util.Minus,
	}

	tmpl := template.Must(template.New("logbook-list.tmpl").Funcs(funcMap).ParseFiles("../template/logbook-list.tmpl"))

	buf := new(bytes.Buffer)
	if err := tmpl.Execute(buf, logbooks); err != nil {
		return err
	}

	content, err := linebot.UnmarshalFlexMessageJSON([]byte(buf.String()))
	if err != nil {
		return err
	}

	if _, err := app.bot.ReplyMessage(
		replyToken,
		linebot.NewFlexMessage("Logbook List", content),
	).Do(); err != nil {
		return err
	}

	return nil
}

func (app *logbookBroker) handleText(message *linebot.TextMessage, replyToken string, source *linebot.EventSource) error {
	command, input := getUserInput(message.Text)

	switch command {
	case constant.Command.Help:
		return app.help(replyToken)
	case constant.Command.HelpList:
		return app.helpList(replyToken)
	case constant.Command.HelpLogbook:
		return app.helpLogbook(replyToken)
	case constant.Command.HelpLogin:
		return app.helpLogin(replyToken)
	case constant.Command.HelpStatus:
		return app.helpStatus(replyToken)
	case constant.Command.List:
		return app.list(replyToken, source.UserID)
	case constant.Command.Logbook:
		return app.logbook(input, replyToken, source.UserID)
	case constant.Command.Login:
		return app.loginAndRecord(input, replyToken, source.UserID)
	case constant.Command.Status:
		return app.status(replyToken, source.UserID)
	}

	return nil
}
