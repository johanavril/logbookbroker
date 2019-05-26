package service

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/johanavril/logbookbroker/src/constant"
)

func GetName() (string, error) {
	resp, err := GetLogbookClient().Get(constant.URL.Profile)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", err
	}

	name := doc.Find(".profile > .row:nth-child(2)").Text()

	return strings.TrimSpace(name), nil
}
