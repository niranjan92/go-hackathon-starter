package apiExamples

import (
	"github.com/gobuffalo/buffalo"

	"net/http"

	"github.com/pkg/errors"
	"github.com/yhat/scrape"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"

	"github.com/niranjan92/go-hackathon-starter/render"
)

// add a better crawling example
// https://github.com/PuerkitoBio/gocrawl/blob/master/cmd/example/main.go
// grab all articles and print them
type article struct {
	Text string `json:"text"`
	Href string `json:"href"`
}

const numArticles = 25

// ScrapingHandler is a default handler to serve up samples for scraping
// hackernews topics
func ScrapingHandler(c buffalo.Context) error {
	url := "https://news.ycombinator.com/"
	articles, err := getArticles(url)
	if err != nil {
		c.Set("errors", err)
		return errors.WithStack(err)
	}
	c.Set("articles", articles)
	return c.Render(200, render.R.HTML("api-examples/scraping.html"))
}

func getArticles(url string) ([]article, error) {
	// request and parse the front page
	resp, err := http.Get("https://news.ycombinator.com/")
	if err != nil {
		return nil, err
	}
	root, err := html.Parse(resp.Body)
	if err != nil {
		return nil, err
	}

	res := []article{}
	matcher := getmatcher()
	parsedArticles := scrape.FindAll(root, matcher)
	for _, art := range parsedArticles {
		res = append(res, article{Text: scrape.Text(art), Href: scrape.Attr(art, "href")})
	}
	return res[:numArticles], nil
}

// add your custom matcher here
func getmatcher() scrape.Matcher {
	return func(n *html.Node) bool {
		// must check for nil values
		if n.DataAtom == atom.A && n.Parent != nil && n.Parent.Parent != nil {
			return scrape.Attr(n.Parent.Parent, "class") == "athing"
		}
		return false
	}
}
