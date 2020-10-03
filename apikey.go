package apikeys

import (
	"github.com/gdbu/dbl"
	"github.com/hatchify/errors"
)

func newAPIKey(userID, name, key string) (a APIKey) {
	a.UserID = userID
	a.Name = name
	a.Key = key
	return
}

// APIKey represents an api key reference
type APIKey struct {
	dbl.Entry

	UserID string `json:"userID"`

	Name string `json:"name"`
	Key  string `json:"key"`

	UpdatedAt int64 `json:"updatedAt"`
}

// Validate will validate an API key
func (a *APIKey) Validate() (err error) {
	var errs errors.ErrorList
	if len(a.UserID) == 0 {
		errs.Push(ErrInvalidUserID)
	}

	if len(a.Name) == 0 {
		errs.Push(ErrInvalidName)
	}

	return errs.Err()
}

// core.Value interface methods below

// GetID will get the message ID
func (a *APIKey) GetID() (id string) { return a.ID }

// GetCreatedAt will get the created at timestamp
func (a *APIKey) GetCreatedAt() (createdAt int64) { return a.CreatedAt }

// GetUpdatedAt will get the updated at timestamp
func (a *APIKey) GetUpdatedAt() (updatedAt int64) { return a.UpdatedAt }

// GetRelationships will get the associated relationship IDs
func (a *APIKey) GetRelationships() (r dbl.Relationships) {
	r.Append(a.Key)
	r.Append(a.UserID)
	return
}

// SetID will get the message ID
func (a *APIKey) SetID(id string) { a.ID = id }

// SetCreatedAt will get the created at timestamp
func (a *APIKey) SetCreatedAt(createdAt int64) { a.CreatedAt = createdAt }

// SetUpdatedAt will get the updated at timestamp
func (a *APIKey) SetUpdatedAt(updatedAt int64) { a.UpdatedAt = updatedAt }
