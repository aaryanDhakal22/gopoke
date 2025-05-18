package utils

import "strings"

func CleanInput(text string) []string {
	ftd := strings.Trim(text, " ")
	ftd = strings.Trim(ftd, "\n")
	ftd = strings.ToLower(ftd)
	ret := strings.Split(ftd, " ")
	return ret
}
