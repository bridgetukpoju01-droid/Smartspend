package main

import (
	"fmt"
	"strconv"
	"strings"
)

func caseconvert(s string) string {
	word := strings.Fields(s)

	for i := 1; i < len(word); i++ {
		if word[i] == "(up)" {
			word[i-1] = strings.ToUpper(word[i-1])
			word = append(word[:i], word[i+1:]...)
			i--
		} else if word[i] == "(low)" {
			word[i-1] = strings.ToLower(word[i-1])
			word = append(word[:i], word[i+1:]...)

		} else if word[i] == "(cap)" {
			word[i-1] = strings.Title(word[i-1])
			word = append(word[:i], word[i+1:]...)

		}
	}

	return strings.Join(word, " ")
}

func convertN(s string) string {
	word := strings.Fields(s)
	for i := 1; i < len(word); i++ {
		if strings.HasPrefix(word[i], "(") && !strings.HasSuffix(word[i], ")") {
			mod := word[i] + word[i+1]
			mod = strings.Trim(mod, "()")
			modifer := strings.Split(mod, ",")
			val, _ := strconv.Atoi(modifer[1])
			for j := i - val; j < i; j++ {
				switch modifer[0] {
				case "up":
					word[j] = strings.ToUpper(word[j])

				case "low":
					word[j] = strings.ToLower(word[j])
				case "cap":
					word[j] = strings.ToLower(word[j])
					word[j] = strings.Title(word[j])
				}

			}
			word = append(word[:i], word[i+2:]...)

		}

	}
	return strings.Join(word, " ")
}
func main() {
	fmt.Println(convertN("hello world (up) HUNDRED CARS (cap, 2)"))
	fmt.Println(caseconvert("hello, world (cap)"))
}
