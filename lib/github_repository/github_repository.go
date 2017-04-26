package github_repository

import (
	"regexp"

	"fmt"

	"context"

	"github.com/google/go-github/github"
	"github.com/zerkms/awesome-stats/lib"
	"golang.org/x/oauth2"
)

var ownerProjectRegex = regexp.MustCompile(`^https://github.com/([^/]+)/([^/]+)`)

type GithubRepository struct {
	AccessToken string
}

func splitOwnerAndProject(url lib.URL) (owner, project string, err error) {
	match := ownerProjectRegex.FindStringSubmatch(string(url))

	if len(match) == 0 {
		return "", "", fmt.Errorf("Github url in wrong format: %s", url)
	}

	return match[1], match[2], nil
}

func (gh GithubRepository) Fetch(url lib.URL) (lib.Stats, error) {
	result := lib.Stats{}

	owner, project, err := splitOwnerAndProject(url)
	if err != nil {
		return result, err
	}

	ctx := context.TODO()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: gh.AccessToken},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)

	repo, _, err := client.Repositories.Get(ctx, owner, project)
	if err != nil {
		return result, err
	}

	result.Stars = *repo.StargazersCount
	result.ID = fmt.Sprintf("github.com:%s:%s", owner, project)

	return result, nil
}
