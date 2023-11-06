package utils

import "strings"

// remove all whitespaces and newline characters from a given string
func SanitizeString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", "")
}
