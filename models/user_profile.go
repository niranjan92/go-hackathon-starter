package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/pop/nulls"
	"github.com/gobuffalo/uuid"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
)

// UserProfile is orm for database user profile
// saves profile per login provider like facebook, github etc
type UserProfile struct {
	ID         uuid.UUID    `json:"id" db:"id"`
	Name       string       `json:"name" db:"name"`
	CreatedAt  time.Time    `json:"created_at" db:"created_at"`
	UpdatedAt  time.Time    `json:"updated_at" db:"updated_at"`
	Provider   string       `json:"provider" db:"provider"`
	ProviderID string       `json:"provider_id" db:"provider_id"`
	Email      nulls.String `json:"email" db:"email"`
	Data       string       `json:"data" db:"data"`
	UserID     uuid.UUID    `json:"user_id" db:"user_id"`
}

// String is not required by pop and may be deleted
func (u UserProfile) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// UserProfiles is not required by pop and may be deleted
type UserProfiles []UserProfile

// String is not required by pop and may be deleted
func (u UserProfiles) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *UserProfile) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Provider, Name: "Provider"},
		&validators.StringIsPresent{Field: u.ProviderID, Name: "ProviderID"},
		&validators.EmailIsPresent{Field: u.Email.String, Name: "Email"},
		&validators.StringIsPresent{Field: u.Data, Name: "Data"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *UserProfile) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *UserProfile) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
