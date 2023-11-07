package model

import (
	"fmt"
	"strings"
)

type TargetType string

const (
	OrganizationTarget = "org"
	UserTarget         = "user"
)

// Represents a GitHub user or organization.
type Target struct {
	Name string
	Type TargetType
}

func ParseTarget(target string) (Target, error) {
	errWrongTargetFormat := fmt.Errorf(
		"target '%s' is wrong format, it must be either '%s' or '%s'",
		target,
		"'user:user-name'",
		"'org:organization-name'",
	)
	splits := strings.Split(target, ":")

	if len(splits) != 2 {
		return Target{}, errWrongTargetFormat
	}

	switch splits[0] {
	case OrganizationTarget:
		return Target{
			Name: splits[1],
			Type: OrganizationTarget,
		}, nil
	case UserTarget:
		return Target{
			Name: splits[1],
			Type: UserTarget,
		}, nil
	default:
		return Target{}, errWrongTargetFormat
	}
}

func (t *Target) GetRepositoriesPageUrl() string {
	switch t.Type {
	case UserTarget:
		return "https://github.com/" + t.Name + "?tab=repositories"
	case OrganizationTarget:
		return "https://github.com/orgs/" + t.Name + "/repositories"
	default:
		return ""
	}
}

// A repository hosted on GitHub.
type Repository struct {
	// Name of the repository.
	Name string

	// Absolute Url to the repository's GitHub page.
	Url string
}
