package actions

import "github.com/gobuffalo/buffalo"

// APIExampleHandler is a default handler to serve up
// a home page.
func APIExampleHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("api-example.html"))
}
