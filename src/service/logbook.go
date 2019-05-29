package service

import (
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/johanavril/logbookbroker/src/constant"
	"github.com/johanavril/logbookbroker/src/util"
)

type action struct {
	Label string
	URI   string
}

type Logbook struct {
	Date        string
	ClockIn     string
	ClockOut    string
	Activity    string
	Description string
	Action      action
}

func GetLogbook(date time.Time) (*Logbook, error) {
	resp, err := GetLogbookClient().Get(constant.URL.Logbook)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	selector := fmt.Sprintf(".active.attached.tab.segment tbody > tr:nth-child(%d)", date.Day())
	tr := doc.Find(selector)

	logbook := Logbook{}
	tr.Children().Each(func(i int, s *goquery.Selection) {
		switch i {
		case 0, 1, 2, 3, 4:
			text := strings.TrimSpace(s.Text())
			if text == "" {
				text = " "
			}
			reflect.ValueOf(&logbook).Elem().Field(i).SetString(text)
		}
	})

	return &logbook, nil
}

func GetWeekLogbook() ([]Logbook, error) {
	resp, err := GetLogbookClient().Get(constant.URL.Logbook)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}
	days := util.ThisWeekRange()

	logbooks := []Logbook{}

	for _, day := range days {
		selector := fmt.Sprintf(`.attached.tab.segment[data-tab^="%s"] tbody > tr:nth-child(%d)`, day.Month(), day.Day())
		tr := doc.Find(selector)

		logbook := Logbook{}
		tr.Children().Each(func(j int, td *goquery.Selection) {
			switch j {
			case 0, 1, 2, 3, 4:
				text := strings.TrimSpace(td.Text())
				if text == "" {
					text = " "
				}
				reflect.ValueOf(&logbook).Elem().Field(j).SetString(text)
			case 6:
				if td.Children() == nil {
					break
				}
				logbook.Action.Label = td.Children().Find("button").Text()
				logbook.Action.URI, _ = td.Children().Attr("href")
			}
		})

		logbooks = append(logbooks, logbook)
	}

	return logbooks, nil
}

func GetEditURL(token string) string {
	return fmt.Sprintf("%s/%s/edit", constant.URL.Logbook, token)
}

func ConstructEditMessage(logbook Logbook) string {
	return fmt.Sprintf("edit-tmp\\n%s\\n%s\\n%s\\n%s\\n%s",
		strings.Split(logbook.Date, " ")[1],
		logbook.ClockIn,
		logbook.ClockOut,
		logbook.Activity,
		logbook.Description)
}
