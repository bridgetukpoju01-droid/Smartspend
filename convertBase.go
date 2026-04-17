package main

import (
	"fmt"
	"strings"
)

func convertBase(s string) string {
	words := strings.Fields(s)
	for i := 1; i < len(words); i++ {
		switch strings.ToLower(words[i]) {
		case "(up)":
			words[i-1] = strings.ToUpper(words[i-1])
			words = append(words[:i], words[i+1:]...)

		case "(low)":
			words[i-1] = strings.ToLower(words[i-1])
			words = append(words[:i], words[i+1:]...)

		case "(cap)":
			words[i-1] = strings.Title(strings.ToLower(words[i-1]))
			words = append(words[:i], words[i+1:]...)

		}

	}
	return strings.Join(words, " ")
}
func main() {
	fmt.Println(convertBase("welcome home (up)"))
	fmt.Println(convertBase("home SWEET (LOW) "))
	fmt.Println(convertBase("we are very happy (cap) "))
}
