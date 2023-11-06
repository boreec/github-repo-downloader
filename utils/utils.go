package utils

import (
	"fmt"
	"strings"
)

// remove all whitespaces and newline characters from a given string
func SanitizeString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", "")
}

func DetermineTargetUrl(target string) (string, error) {
	errWrongFromat := fmt.Errorf(
		"target '%s' is wrong format, it must be either '%s' or '%s'",
		target,
		"'user:user-name'",
		"'org:organization-name'",
	)

	splits := strings.Split(target, ":")
	if len(splits) != 2 {
		return "", errWrongFromat
	}

	switch splits[0] {
	case "org":
		return "https://github.com/orgs/" + splits[1] + "/repositories", nil
	case "user":
		return "https://github.com/" + splits[1] + "?tab=repositories", nil
	default:
		return "", errWrongFromat
	}
}
