package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	Barista struct {
		id        uuid.UUID
		firstname string
		lastname  string
		createdAt time.Time
		updatedAt time.Time
	}

	BaristaRepository interface {
		FindById(context.Context, uuid.UUID) (Barista, error)
	}
)

var (
	ErrBaristaNotFound = errors.New("barista not found")
)

// ID returns the Barista's id.
func (b *Barista) ID() uuid.UUID {
	return b.id
}

// Firstname returns the Barista's firstname.
func (b *Barista) Firstname() string {
	return b.firstname
}

// Lastname returns the Barista's lastname.
func (b *Barista) Lastname() string {
	return b.lastname
}

// CreatedAt returns the Barista's creation time.
func (b *Barista) CreatedAt() time.Time {
	return b.createdAt
}

// UpdatedAt returns the Barista's last update time.
func (b *Barista) UpdatedAt() time.Time {
	return b.updatedAt
}
