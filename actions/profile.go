package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"

	"github.com/gobuffalo/pop/nulls"
	"github.com/niranjan92/go-hackathon-starter/actions/render"
	"github.com/niranjan92/go-hackathon-starter/models"
	"github.com/pkg/errors"
)

// ProfileHandler is a default handler to serve up
// a profile page.
func ProfileHandler(c buffalo.Context) error {
	// get current user details to show
	return c.Render(200, render.R.HTML("profile/profile.html"))
}

// UpdateProfileHandler is a default handler to serve up
// a profile page.
func UpdateProfileHandler(c buffalo.Context) error {
	// get current user details set initially by middleware
	u := c.Data()["current_user"].(*models.User)

	u.Name = c.Request().Form.Get("Name")
	u.Email = nulls.NewString(c.Request().Form.Get("Email"))
	u.Gender = nulls.NewString(c.Request().Form.Get("Gender"))
	u.Location = nulls.NewString(c.Request().Form.Get("Location"))
	u.Website = nulls.NewString(c.Request().Form.Get("Website"))

	tx := c.Value("tx").(*pop.Connection)
	if err := tx.Save(u); err != nil {
		return errors.WithStack(err)
	}

	c.Set("current_user", u)

	c.Flash().Add("success", "profile updated!")
	return c.Render(200, render.R.HTML("profile/profile.html"))
}
