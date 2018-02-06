package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/niranjan92/go_hackathon_starter/models"
	"github.com/pkg/errors"
)

// GetUploadHandler is a default handler to serve up
// a /api/uploads page.
func GetUploadHandler(c buffalo.Context) error {
	// get current user details to show
	c.Set("widget", &models.Widget{})
	return c.Render(200, r.HTML("api-examples/upload.html"))
}

// PostUploadHandler is used to handle a file upload
func PostUploadHandler(c buffalo.Context) error {
	f, err := c.File("MyFile") // change name - SomeFile
	if err != nil {
		c.Set("errors", err)
		return errors.WithStack(err)
	}
	widget := &models.Widget{Name: "test_widget", Uploadfile: f}
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(widget)
	if err != nil {
		c.Set("errors", err)
		return errors.WithStack(err)
	}
	if verrs.HasAny() {
		// Make contact available inside the html template
		c.Set("widget", widget)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("/"))
		//return c.Render(422, r.HTML("api-examples/upload.html"))
	}
	// If there are no errors set a success message
	c.Flash().Add("success", "File was uploaded successfully")

	// and redirect to the home index page
	return c.Redirect(302, "/")
}
