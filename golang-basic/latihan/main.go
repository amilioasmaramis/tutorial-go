package main

import (
	"fmt"
	"regexp"
	"strings"
)

func Palindrome(str string) bool {
	var tmp = len(str)
	tmpStr := []string{}
	for i := tmp - 1; i >= 0; i-- {
		tmpStr = append(tmpStr, string(str[i]))
	}
	resultJoin := strings.Join(tmpStr, "")
	return resultJoin == str
}

func VocalInput(str string) string {
	switch true {
	case len(str) > 1:
		return "Str length should be 1 character"
	case str == "":
		return "Str should not be empty"
	}

	match, err := regexp.MatchString(`^[aiueoAIUEO]*$`, str)
	if err != nil {
		fmt.Println("Error", err)
	}
	fmt.Println(match, "match")
	if match == true {
		return "Vocal"
	}
	return "Konsonan"
}

func main() {
	fmt.Println(Palindrome("katak"))
	fmt.Println(VocalInput("c"))
}
