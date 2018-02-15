package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/niranjan92/go-hackathon-starter/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
