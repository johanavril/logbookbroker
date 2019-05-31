package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/johanavril/logbookbroker/src/constant"
)

type message struct {
	Type string `json:"type"`
	Text string `json:"text"`
}

func RequestEditReminder(channelToken string) error {
	m := message{
		Type: "text",
		Text: constant.Message.RequestEditReminder,
	}
	data := map[string][]message{"messages": []message{m}}
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, constant.URL.Broadcast, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+channelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	return nil
}

func SubmitReminder(channelToken string) error {
	m := message{
		Type: "text",
		Text: constant.Message.SubmitReminder,
	}
	data := map[string][]message{"messages": []message{m}}
	body, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return err
	}

	req, err := http.NewRequest(http.MethodPost, constant.URL.Broadcast, bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+channelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return err
	}

	defer resp.Body.Close()

	return nil
}
