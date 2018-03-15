package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/niranjan92/go-hackathon-starter/render"
)

// ReactHandler is a default handler to serve up
// a home page. TODO
func ReactHandler(c buffalo.Context) error {
	return c.Render(200, render.ReactR.HTML("index.html"))
}
