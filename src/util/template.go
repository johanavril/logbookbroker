package util

import "strings"

func Minus(a, b int) int {
	return a - b
}

func ExtractEditToken(url string) string {
	u := strings.Split(url, "/")
	return u[6]
}
