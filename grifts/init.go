package grifts

import (
	"github.com/gobuffalo/buffalo"
	app "github.com/niranjan92/go-hackathon-starter/actions"
)

func init() {
	buffalo.Grifts(app.App())
}
