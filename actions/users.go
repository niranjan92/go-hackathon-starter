package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/gobuffalo/pop"
	"github.com/niranjan92/go-hackathon-starter/models"
	"github.com/pkg/errors"
)

// UserDestroy deletes a User from the DB. This function is mapped
// to the path DELETE /users/{user_id}
func UserDestroy(c buffalo.Context) error {
	// clear user session first
	c.Session().Clear()

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty User
	user := &models.User{}

	// To find the User the parameter user_id is used.
	if err := tx.Find(user, c.Param("user_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(user); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Account was destroyed successfully")

	// Redirect to the home page
	return c.Redirect(302, "/")
}
