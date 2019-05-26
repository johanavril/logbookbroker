package service

import (
	"errors"
	"io"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/johanavril/logbookbroker/src/constant"
)

var logbookClient *http.Client
var once sync.Once

func initClient() {
	jar, _ := cookiejar.New(nil)
	var cookies []*http.Cookie

	cookie := &http.Cookie{
		Name:  "XSRF-TOKEN",
		Value: "",
	}
	cookies = append(cookies, cookie)

	cookie = &http.Cookie{
		Name:  "laravel_session",
		Value: "",
	}
	cookies = append(cookies, cookie)

	u, _ := url.Parse(constant.URL.Login)
	jar.SetCookies(u, cookies)
	client := &http.Client{
		Jar: jar,
	}

	logbookClient = client
}

func GetLogbookClient() *http.Client {
	once.Do(func() {
		initClient()
	})

	return logbookClient
}

func GetCSRF(url string) ([]*http.Cookie, string, error) {
	resp, err := GetLogbookClient().Get(url)
	if err != nil {
		return nil, "", err
	}

	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, "", err
	}

	csrf, exists := doc.Find("input[type=hidden]").Attr("value")
	if !exists {
		return nil, "", errors.New("CSRF Token doesn't exists")
	}

	return resp.Cookies(), csrf, nil
}

func GetPageTitle(body io.Reader) (string, error) {
	doc, err := goquery.NewDocumentFromReader(body)
	if err != nil {
		return "", err
	}

	title := doc.Find("title").Text()

	return strings.TrimSpace(title), nil
}
