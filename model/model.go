package model

import (
	"fmt"
	"strings"
)

// CloningTargetType represents the type of a cloning target, which can be a
// GitHub user or a GitHub organization.
type CloningTargetType string

const (
	OrganizationTarget = "org"
	UserTarget         = "user"
)

// CloningTarget represents the target from which repositories are cloned.
type CloningTarget struct {
	Name string            // Name of the Target.
	Type CloningTargetType // Type of the cloning target ("org" or "user").
}

// ParseCloningTarget parses a string into a CloningTarget.
//
// Parameters:
//   - target: A string with the format `user:user-name` or
//     `org:organization-name`.
//
// Returns:
//   - cloningTarget: The parsed cloning target.
//   - err: An error, if any, that occurred during parsing.
func ParseCloningTarget(target string) (
	cloningTarget CloningTarget,
	err error,
) {
	errWrongTargetFormat := fmt.Errorf(
		"target '%s' is wrong format, it must be either '%s' or '%s'",
		target,
		"'user:user-name'",
		"'org:organization-name'",
	)
	splits := strings.Split(target, ":")

	if len(splits) != 2 {
		return CloningTarget{}, errWrongTargetFormat
	}

	switch splits[0] {
	case OrganizationTarget:
		return CloningTarget{
			Name: splits[1],
			Type: OrganizationTarget,
		}, nil
	case UserTarget:
		return CloningTarget{
			Name: splits[1],
			Type: UserTarget,
		}, nil
	default:
		return CloningTarget{}, errWrongTargetFormat
	}
}

// GetRepositoriesPageUrl builds the URL containing the repositories for a
// target based on its name and type.
//
// Returns:
//    - url: The URL to the repositories page.
func (t *CloningTarget) GetRepositoriesPageUrl() (url string) {
	switch t.Type {
	case UserTarget:
		return "https://github.com/" + t.Name + "?tab=repositories"
	case OrganizationTarget:
		return "https://github.com/orgs/" + t.Name + "/repositories"
	default:
		return ""
	}
}

// Repository represents a GitHub repository.
type Repository struct {
	// Name of the repository.
	Name string

	// Absolute Url to the repository's GitHub page.
	Url string
}
