package apiExamples

import (
	"os"
	"time"

	"github.com/ChimeraCoder/anaconda"
	"github.com/gobuffalo/buffalo"
	"github.com/niranjan92/go-hackathon-starter/render"
	"github.com/pkg/errors"
)

var twitterAPI *anaconda.TwitterApi

func init() {
	anaconda.SetConsumerKey(os.Getenv("TWITTER_CONSUMER_KEY"))
	anaconda.SetConsumerSecret(os.Getenv("TWITTER_CONSUMER_SECRET"))
	twitterAPI = anaconda.NewTwitterApi(os.Getenv("TWITTER_API_TOKEN"), os.Getenv("TWITTER_API_SECRET"))
}

const twitterTimeout = 5 * time.Second

// TwitterHandler is a default handler to serve up samples for Github api
func TwitterHandler(c buffalo.Context) error {
	query := "#tesla"
	searchResult, err := twitterAPI.GetSearch(query, nil)
	if err != nil {
		c.Set("errors", err)
		return errors.WithStack(err)
	}
	tweets := []string{}
	for _, tweet := range searchResult.Statuses {
		if tweet.RetweetedStatus == nil {
			tweets = append(tweets, tweet.Text)
		}
	}
	c.Set("query", query)
	c.Set("tweets", tweets)
	return c.Render(200, render.R.HTML("api-examples/twitter.html"))
}
