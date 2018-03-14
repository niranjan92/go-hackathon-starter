package actions

import (
	"github.com/gobuffalo/buffalo"

	"github.com/niranjan92/go-hackathon-starter/actions/render"
)

// HomeHandler is a default handler to serve up
// a home page.
func HomeHandler(c buffalo.Context) error {
	return c.Render(200, render.R.HTML("index.html"))
}
