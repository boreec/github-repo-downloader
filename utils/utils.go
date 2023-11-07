package utils

import (
	"fmt"
	"strings"
)

// Remove all whitespaces and newline characters from a given string.
func SanitizeString(s string) string {
	return strings.ReplaceAll(strings.ReplaceAll(s, " ", ""), "\n", "")
}

// Determine the appropriate url to use for a target, depending if it is an
// user or an organization.
//
// Parameters:
//   - target: the profile from which repositories are going to fetched. It
//   it must follow the format `org:organization-name` or `user:user-name`.
//
// Returns:
//   - url: the target's url for fetching repositories.
//   - err: an errory, if any, that occurred during the process.
func DetermineTargetUrl(target string) (url string, err error) {
	errWrongFormat := fmt.Errorf(
		"target '%s' is wrong format, it must be either '%s' or '%s'",
		target,
		"'user:user-name'",
		"'org:organization-name'",
	)

	splits := strings.Split(target, ":")
	if len(splits) != 2 {
		return "", errWrongFormat
	}

	switch splits[0] {
	case "org":
		return "https://github.com/orgs/" + splits[1] + "/repositories", nil
	case "user":
		return "https://github.com/" + splits[1] + "?tab=repositories", nil
	default:
		return "", errWrongFormat
	}
}
