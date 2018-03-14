package apiExamples

import (
	"context"
	"time"

	"github.com/gobuffalo/buffalo"
	"github.com/google/go-github/github"
	"github.com/pkg/errors"

	"github.com/niranjan92/go-hackathon-starter/actions/render"
)

const githubTimeout = 5 * time.Second

// GithubHandler is a default handler to serve up samples for Github api
func GithubHandler(c buffalo.Context) error {
	owner, repoName := "niranjan92", "go-hackathon-starter" //TODO: replace this with mine
	client := github.NewClient(nil)
	ctx, cancel := context.WithTimeout(context.Background(), githubTimeout)
	defer cancel()

	repo, _, err := client.Repositories.Get(ctx, owner, repoName)
	if err != nil {
		return errors.WithStack(err)
	}
	c.Set("subscribers", repo.GetSubscribersCount())
	c.Set("forks", repo.GetForksCount())
	c.Set("owner", owner)
	c.Set("repoName", repoName)
	c.Set("stars", repo.GetStargazersCount())

	return c.Render(200, render.R.HTML("api-examples/github.html"))
}
