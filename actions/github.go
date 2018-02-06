package actions

import (
	"context"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"
)

const githubTimeout = 5 * time.Second

// GithubHandler is a default handler to serve up samples for Github api
func GithubHandler(c buffalo.Context) error {
	owner, repoName := "sahat", "hackathon-starter" //TODO: replace this with mine
	client := github.NewClient(nil)
	ctx, _ := context.WithTimeout(context.Background(), githubTimeout)
	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("subscribers", repo.GetSubscribersCount())
	c.Set("forks", repo.GetForksCount())
	c.Set("owner", owner)
	c.Set("repoName", repoName)
	c.Set("stars", repo.GetStargazersCount())

	return c.Render(200, r.HTML("api-examples/github.html"))
}
