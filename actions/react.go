package actions

import "github.com/gobuffalo/buffalo"

// ReactHandler is a default handler to serve up
// a home page.
// TODO: fix jsx part of this
func ReactHandler(c buffalo.Context) error {
	return c.Render(200, reactR.HTML("index.html"))
}
