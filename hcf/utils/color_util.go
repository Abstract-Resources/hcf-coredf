package utils

import "strings"

func Colour(text string) string {
	return strings.Replace(text, "&", "§", -1)
}
