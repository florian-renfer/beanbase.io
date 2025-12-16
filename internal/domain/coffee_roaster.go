package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	CoffeeRoaster struct {
		id            uuid.UUID
		name          string
		onlineShopURL string
		createdAt     time.Time
		updatedAt     time.Time
	}

	CoffeeRoasterRepository interface {
		Create(context.Context, CoffeeRoaster) (CoffeeRoaster, error)
	}
)

var (
	ErrCoffeeRoasterExistsAlready = errors.New("coffee roaster exists already")
	ErrCoffeeRoasterNotFound      = errors.New("coffee roaster not found")
)

func NewCoffeeRoaster(id uuid.UUID, name, onlineShopURL string, createdAt, updatedAt time.Time) CoffeeRoaster {
	return CoffeeRoaster{
		id:            id,
		name:          name,
		onlineShopURL: onlineShopURL,
		createdAt:     createdAt,
		updatedAt:     updatedAt,
	}
}

// ID returns the Roaster's id.
func (r CoffeeRoaster) ID() uuid.UUID {
	return r.id
}

// Name returns the Roaster's name.
func (r CoffeeRoaster) Name() string {
	return r.name
}

// OnlineShopURL returns the Roaster's online shop URL.
func (r CoffeeRoaster) OnlineShopURL() string {
	return r.onlineShopURL
}

// CreatedAt returns the Roaster's creation time.
func (r CoffeeRoaster) CreatedAt() time.Time {
	return r.createdAt
}

// UpdatedAt returns the Roaster's last update time.
func (r CoffeeRoaster) UpdatedAt() time.Time {
	return r.updatedAt
}
