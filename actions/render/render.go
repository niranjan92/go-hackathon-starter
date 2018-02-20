package render

import (
	"github.com/gobuffalo/buffalo/render"
	"github.com/gobuffalo/packr"
)

// ReactR acts as a renderer for react templates
var ReactR *render.Engine

// R is a renderer
var R *render.Engine

// AssetsBox ...
var AssetsBox = packr.NewBox("../../public/assets")

func init() {
	R = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../../templates"),
		AssetsBox:    AssetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{},
	})

	ReactR = render.New(render.Options{
		// HTML layout to be used for all HTML requests:
		HTMLLayout: "application.html",

		// Box containing all of the templates:
		TemplatesBox: packr.NewBox("../../react_templates"),
		AssetsBox:    AssetsBox,

		// Add template helpers here:
		Helpers: render.Helpers{},
	})
}
