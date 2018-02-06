package actions

import "github.com/gobuffalo/buffalo"

// DashboardHandler is a default handler to serve up
// a home page.
func DashboardHandler(c buffalo.Context) error {
	return c.Render(200, r.HTML("dashboard/dashboard.html"))
}
