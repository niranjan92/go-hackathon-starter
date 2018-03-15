package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/niranjan92/go-hackathon-starter/render"
)

// LoginHandler is a default handler to serve up
// a home page.
func LoginHandler(c buffalo.Context) error {
	return c.Render(200, render.R.HTML("login.html"))
}
