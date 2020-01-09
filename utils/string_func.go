package utils

import "strings"

func ReplaceString(src string, s []string) (newString string) {
	newString = src
	for _, each := range s {
		newString = strings.ReplaceAll(newString, each, "")
	}
	return newString
}
