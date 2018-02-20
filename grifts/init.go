package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/niranjan92/go-hackathon-starter/actions/app"
)

func init() {
	buffalo.Grifts(app.App())
}
