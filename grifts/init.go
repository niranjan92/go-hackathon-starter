package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/niranjan92/go_hackathon_starter/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
