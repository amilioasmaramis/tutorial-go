package main

import (
	"fmt"
	"regexp"
)

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
