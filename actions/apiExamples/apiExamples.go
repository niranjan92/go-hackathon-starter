package apiExamples

import (
	"github.com/gobuffalo/buffalo"
	"github.com/niranjan92/go-hackathon-starter/actions/render"
)

// APIExampleHandler is a default handler to serve up
// a home page.
func APIExampleHandler(c buffalo.Context) error {
	return c.Render(200, render.R.HTML("api-example.html"))
}
