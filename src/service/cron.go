package service

import (
	"bytes"
	"encoding/json"
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
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, constant.URL.Broadcast, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+channelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
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
		return nil
	}

	req, err := http.NewRequest(http.MethodPost, constant.URL.Broadcast, bytes.NewBuffer(body))
	if err != nil {
		return nil
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+channelToken)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil
	}

	defer resp.Body.Close()

	return nil
}
