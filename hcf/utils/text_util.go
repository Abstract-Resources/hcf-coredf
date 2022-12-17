package utils

import "strings"

func Colour(text string) string {
	return strings.Replace(text, "&", "§", -1)
}

func ReplacePlaceHolders(key string, args ...string) string {

}