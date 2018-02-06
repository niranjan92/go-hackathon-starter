package actions

import (
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gobuffalo/buffalo"
	"github.com/pkg/errors"
)

var twitter_api *anaconda.TwitterApi

func init() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	twitter_api = anaconda.NewTwitterApi(os.Getenv("TWITTER_API_TOKEN"), os.Getenv("TWITTER_API_SECRET"))
}

const twitterTimeout = 5 * time.Second

// GithubHandler is a default handler to serve up samples for Github api
func TwitterHandler(c buffalo.Context) error {
	query := "#golang"
	searchResult, err := twitter_api.GetSearch(query, nil)
	if err != nil {
		c.Set("errors", err)
		return errors.WithStack(err)
	}
	tweets := []string{}
	for _, tweet := range searchResult.Statuses {
		tweets = append(tweets, tweet.Text)
	}
	c.Set("query", query)
	c.Set("tweets", tweets)
	return c.Render(200, r.HTML("api-examples/twitter.html"))
}
