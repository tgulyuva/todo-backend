package helpers

import "strings"

func CheckDoubleSpace(text string) bool {
	return strings.Contains(text, "  ")
}

func RemoveDoubleSpace(text string) string {
	return strings.Replace(text, "  ", " ", -1)
}
