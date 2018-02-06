package actions

import "github.com/gobuffalo/buffalo"

// HomeHandler is a default handler to serve up
// a home page.
func ApiExampleHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("api-example.html"))
}
