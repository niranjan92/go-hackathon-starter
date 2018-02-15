package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/markbates/pop"
	"github.com/niranjan92/go-hackathon-starter/models"
	"github.com/pkg/errors"
)

// This file is generated by Buffalo. It offers a basic structure for
// adding, editing and deleting a page. If your model is more
// complex or you need more than the basic implementation you need to
// edit this file.

// Following naming logic is implemented in Buffalo:
// Model: Singular (Contact)
// DB Table: Plural (contacts)
// Resource: Plural (Contacts)
// Path: Plural (/contacts)
// View Template Folder: Plural (/templates/contacts/)

// ContactsResource is the resource for the Contact model
type ContactsResource struct {
	buffalo.Resource
}

// List gets all Contacts. This function is mapped to the path
// GET /contacts
func (v ContactsResource) List(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	contacts := &models.Contacts{}

	// Paginate results. Params "page" and "per_page" control pagination.
	// Default values are "page=1" and "per_page=20".
	q := tx.PaginateFromParams(c.Params())

	// Retrieve all Contacts from the DB
	if err := q.All(contacts); err != nil {
		return errors.WithStack(err)
	}

	// Make Contacts available inside the html template
	c.Set("contacts", contacts)

	// Add the paginator to the context so it can be used in the template.
	c.Set("pagination", q.Paginator)

	return c.Render(200, r.HTML("contacts/index.html"))
}

// Show gets the data for one Contact. This function is mapped to
// the path GET /contacts/{contact_id}
func (v ContactsResource) Show(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Contact
	contact := &models.Contact{}

	// To find the Contact the parameter contact_id is used.
	if err := tx.Find(contact, c.Param("contact_id")); err != nil {
		return c.Error(404, err)
	}

	// Make contact available inside the html template
	c.Set("contact", contact)

	return c.Render(200, r.HTML("contacts/show.html"))
}

// New renders the form for creating a new Contact.
// This function is mapped to the path GET /contacts/new
func (v ContactsResource) New(c buffalo.Context) error {
	// Make contact available inside the html template
	c.Set("contact", &models.Contact{})

	return c.Render(200, r.HTML("contacts/new.html"))
}

// Create adds a Contact to the DB. This function is mapped to the
// path POST /contacts
func (v ContactsResource) Create(c buffalo.Context) error {
	// Allocate an empty Contact
	contact := &models.Contact{}

	// Bind contact to the html form elements
	if err := c.Bind(contact); err != nil {
		return errors.WithStack(err)
	}

	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Validate the data from the html form
	verrs, err := tx.ValidateAndCreate(contact)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make contact available inside the html template
		c.Set("contact", contact)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the new.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("contacts/new.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Contact was created successfully")

	// and redirect to the contacts index page
	return c.Redirect(302, "/")
}

// Edit renders a edit form for a Contact. This function is
// mapped to the path GET /contacts/{contact_id}/edit
func (v ContactsResource) Edit(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Contact
	contact := &models.Contact{}

	if err := tx.Find(contact, c.Param("contact_id")); err != nil {
		return c.Error(404, err)
	}

	// Make contact available inside the html template
	c.Set("contact", contact)
	return c.Render(200, r.HTML("contacts/edit.html"))
}

// Update changes a Contact in the DB. This function is mapped to
// the path PUT /contacts/{contact_id}
func (v ContactsResource) Update(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Contact
	contact := &models.Contact{}

	if err := tx.Find(contact, c.Param("contact_id")); err != nil {
		return c.Error(404, err)
	}

	// Bind Contact to the html form elements
	if err := c.Bind(contact); err != nil {
		return errors.WithStack(err)
	}

	verrs, err := tx.ValidateAndUpdate(contact)
	if err != nil {
		return errors.WithStack(err)
	}

	if verrs.HasAny() {
		// Make contact available inside the html template
		c.Set("contact", contact)

		// Make the errors available inside the html template
		c.Set("errors", verrs)

		// Render again the edit.html template that the user can
		// correct the input.
		return c.Render(422, r.HTML("contacts/edit.html"))
	}

	// If there are no errors set a success message
	c.Flash().Add("success", "Contact was updated successfully")

	// and redirect to the contacts index page
	return c.Redirect(302, "/contacts/%s", contact.ID)
}

// Destroy deletes a Contact from the DB. This function is mapped
// to the path DELETE /contacts/{contact_id}
func (v ContactsResource) Destroy(c buffalo.Context) error {
	// Get the DB connection from the context
	tx := c.Value("tx").(*pop.Connection)

	// Allocate an empty Contact
	contact := &models.Contact{}

	// To find the Contact the parameter contact_id is used.
	if err := tx.Find(contact, c.Param("contact_id")); err != nil {
		return c.Error(404, err)
	}

	if err := tx.Destroy(contact); err != nil {
		return errors.WithStack(err)
	}

	// If there are no errors set a flash message
	c.Flash().Add("success", "Contact was destroyed successfully")

	// Redirect to the contacts index page
	return c.Redirect(302, "/contacts")
}
