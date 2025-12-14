package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	Country struct {
		id              uuid.UUID
		label           string
		isoAbbreviation string
		createdAt       time.Time
		updatedAt       time.Time
	}

	CountryRepository interface {
		Create(context.Context, Country) error
		FindAll(context.Context) ([]Country, error)
	}
)

var (
	ErrCountryNotFound = errors.New("country not found")
)

// ID returns the Country's id.
func (c *Country) ID() uuid.UUID {
	return c.id
}

// Label returns the Country's label.
func (c *Country) Label() string {
	return c.label
}

// ISOAbbreviation returns the Country's ISO abbreviation.
func (c *Country) ISOAbbreviation() string {
	return c.isoAbbreviation
}

// CreatedAt returns the Country's creation time.
func (c *Country) CreatedAt() time.Time {
	return c.createdAt
}

// UpdatedAt returns the Country's last update time.
func (c *Country) UpdatedAt() time.Time {
	return c.updatedAt
}
