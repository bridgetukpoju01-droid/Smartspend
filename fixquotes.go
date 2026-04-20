package main

import (
	"regexp"
)

func fixquotes(text string) string {
	re := regexp.MustCompile(`'\s*(.*?)\s*'`)
	text = re.ReplaceAllString(text, "'$1'")

	return text
}
