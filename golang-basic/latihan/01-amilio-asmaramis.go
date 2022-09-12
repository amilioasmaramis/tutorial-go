package main

import "strings"

func Palindrome(str string) bool {
	var tmp = len(str)
	tmpStr := []string{}
	for i := tmp - 1; i >= 0; i-- {
		tmpStr = append(tmpStr, string(str[i]))
	}
	resultJoin := strings.Join(tmpStr, "")
	return resultJoin == str
}
