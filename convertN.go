package main

import (
	"fmt"
	"strconv"
	"strings"
)

func convertN(s string) string {
	words := strings.Fields(s)
	for i := 1; i < len(words); i++ {
		if strings.HasPrefix(words[i], "(") && strings.HasSuffix(words[i], ")") {
			parts := strings.Split(strings.Trim(words[i], "()"), ",")
			action := strings.ToLower(strings.TrimSpace(parts[0]))
			count := 1
			if len(parts) == 2 {
				if n, err := strconv.Atoi(strings.TrimSpace(parts[1])); err == nil {
					count = n
				}
			}
			for j := i - count; j < i; j++ {
				if j >= 0 {
					switch action {
					case "up":
						words[j] = strings.ToUpper(words[j])
					case "low":
						words[j] = strings.ToLower(words[j])
					case "cap":
						words[j] = strings.Title(strings.ToLower(words[j]))
					}
				}
			}
			words = append(words[:i], words[i+1:]...)
			i--
		}
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Println(convertN("This is so exciting (up,2)"))
	fmt.Println(convertN("I SHOULD stop SHOUTING (low,2)"))
	fmt.Println(convertN("welcome to the brooklyn bridge (cap,3)"))
}
