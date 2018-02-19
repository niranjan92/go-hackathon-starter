package models

import (
	"encoding/json"
	"time"

	"io"
	"os"
	"path/filepath"

	"github.com/gobuffalo/buffalo/binding"
	"github.com/markbates/pop"
	"github.com/markbates/validate"
	"github.com/markbates/validate/validators"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

// Widget is orm for widgets
type Widget struct {
	ID         uuid.UUID    `json:"id" db:"id"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	Name       string       `json:"name" db:"name"`
	Uploadfile binding.File `db:"-" form:"someFile"`
}

// String is not required by pop and may be deleted
func (w Widget) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Widgets is not required by pop and may be deleted
type Widgets []Widget

// String is not required by pop and may be deleted
func (w Widgets) String() string {
	jw, _ := json.Marshal(w)
	return string(jw)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (w *Widget) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: w.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (w *Widget) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (w *Widget) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// AfterCreate gets run every time after you save widget to database
func (w *Widget) AfterCreate(tx *pop.Connection) error {
	if !w.Uploadfile.Valid() {
		return nil
	}
	dir := filepath.Join(".", "uploads")
	if err := os.MkdirAll(dir, 0755); err != nil {
		return errors.WithStack(err)
	}
	f, err := os.Create(filepath.Join(dir, w.Uploadfile.Filename))
	if err != nil {
		return errors.WithStack(err)
	}
	defer f.Close()
	_, err = io.Copy(f, w.Uploadfile)
	return err
}
