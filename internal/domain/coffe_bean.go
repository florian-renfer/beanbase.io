package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	RoastBlend string

	Bean struct {
		id        uuid.UUID
		baristaID uuid.UUID
		roasterID uuid.UUID
		countryID uuid.UUID
		label     string
		blend     RoastBlend
		createdAt time.Time
		updatedAt time.Time
	}

	BeanRepository interface {
		FindByBaristaId(context.Context, uuid.UUID) ([]Bean, error)
	}
)

const (
	RoastBlendLight  RoastBlend = "light"
	RoastBlendMedium RoastBlend = "medium"
	RoastBlendDark   RoastBlend = "dark"
)

var (
	ErrBeanNotFound = errors.New("bean not found")
)

// ID returns the Bean's id.
func (b *Bean) ID() uuid.UUID {
	return b.id
}

// BaristaID returns the Bean's barista id.
func (b *Bean) BaristaID() uuid.UUID {
	return b.baristaID
}

// RoasterID returns the Bean's roaster id.
func (b *Bean) RoasterID() uuid.UUID {
	return b.roasterID
}

// CountryID returns the Bean's country id.
func (b *Bean) CountryID() uuid.UUID {
	return b.countryID
}

// Label returns the Bean's label.
func (b *Bean) Label() string {
	return b.label
}

// RoastBlend returns the Bean's roast blend.
func (b *Bean) RoastBlend() RoastBlend {
	return b.blend
}

// CreatedAt returns the Bean's creation time.
func (b *Bean) CreatedAt() time.Time {
	return b.createdAt
}

// UpdatedAt returns the Bean's last update time.
func (b *Bean) UpdatedAt() time.Time {
	return b.updatedAt
}
