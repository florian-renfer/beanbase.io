package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	Roaster struct {
		id            uuid.UUID
		name          string
		onlineShopURL string
		createdAt     time.Time
		updatedAt     time.Time
	}

	RoasterRepository interface {
		FindById(context.Context, uuid.UUID) (Roaster, error)
		FindByName(context.Context, string) (Roaster, error)
	}
)

var (
	ErrRoasterNotFound = errors.New("roaster not found")
)

// ID returns the Roaster's id.
func (r Roaster) ID() uuid.UUID {
	return r.id
}

// Name returns the Roaster's name.
func (r Roaster) Name() string {
	return r.name
}

// OnlineShopURL returns the Roaster's online shop URL.
func (r Roaster) OnlineShopURL() string {
	return r.onlineShopURL
}

// CreatedAt returns the Roaster's creation time.
func (r Roaster) CreatedAt() time.Time {
	return r.createdAt
}

// UpdatedAt returns the Roaster's last update time.
func (r Roaster) UpdatedAt() time.Time {
	return r.updatedAt
}
