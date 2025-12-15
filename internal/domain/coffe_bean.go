package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	RoastBlend string

	CoffeeBean struct {
		id            uuid.UUID
		baristaID     uuid.UUID
		roasterID     uuid.UUID
		countryID     uuid.UUID
		label         string
		amountArabica uint8
		amountRobusta uint8
		blend         RoastBlend
		createdAt     time.Time
		updatedAt     time.Time
	}

	BeanRepository interface {
		FindByBaristaId(context.Context, uuid.UUID) ([]CoffeeBean, error)
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
func (b CoffeeBean) ID() uuid.UUID {
	return b.id
}

// BaristaID returns the Bean's barista id.
func (b CoffeeBean) BaristaID() uuid.UUID {
	return b.baristaID
}

// RoasterID returns the Bean's roaster id.
func (b CoffeeBean) RoasterID() uuid.UUID {
	return b.roasterID
}

// CountryID returns the Bean's country id.
func (b CoffeeBean) CountryID() uuid.UUID {
	return b.countryID
}

// Label returns the Bean's label.
func (b CoffeeBean) Label() string {
	return b.label
}

// AmountArabica returns thee Bean's amount of arabica.
func (b CoffeeBean) AmountArabica() uint8 {
	return b.amountArabica
}

// AmountRobusta returns thee Bean's amount of robusta.
func (b CoffeeBean) AmountRobusta() uint8 {
	return b.amountRobusta
}

// RoastBlend returns the Bean's roast blend.
func (b CoffeeBean) RoastBlend() RoastBlend {
	return b.blend
}

// CreatedAt returns the Bean's creation time.
func (b CoffeeBean) CreatedAt() time.Time {
	return b.createdAt
}

// UpdatedAt returns the Bean's last update time.
func (b CoffeeBean) UpdatedAt() time.Time {
	return b.updatedAt
}
