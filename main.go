package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Error")
		return
	}
	data, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("wrong output")
		return
	}
	result := convertN(string(data))
	result = convertbase(result)
	result = convertcase(result)
	result = fixquotes(result)
	result = punc(result)
	result = vowels(result)

	os.WriteFile(os.Args[2], []byte(result), 0644)
	fmt.Println("omor, why tech.. Go output.txt jor")

}
