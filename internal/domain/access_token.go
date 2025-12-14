package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	AccessToken struct {
		id           uuid.UUID
		baristaID    uuid.UUID
		value        string
		refreshToken string
		expiresAt    time.Time
		createdAt    time.Time
		updatedAt    time.Time
	}

	AccessTokenRepository interface {
		FindByToken(context.Context, string) (AccessToken, error)
		FindByIsExpired(context.Context) ([]AccessToken, error)
	}
)

var (
	ErrAccessTokenNotFound = errors.New("access token not found")
)

// ID returns the AccessToken's id.
func (a *AccessToken) ID() uuid.UUID {
	return a.id
}

// BaristaID returns the AccessToken's barista id.
func (a *AccessToken) BaristaID() uuid.UUID {
	return a.baristaID
}

// Value returns the AccessToken's value.
func (a *AccessToken) Value() string {
	return a.value
}

// RefreshToken returns the AccessToken's refresh token.
func (a *AccessToken) RefreshToken() string {
	return a.refreshToken
}

// ExpiresAt returns the AccessToken's expiration time.
func (a *AccessToken) ExpiresAt() time.Time {
	return a.expiresAt
}

// CreatedAt returns the AccessToken's creation time.
func (a *AccessToken) CreatedAt() time.Time {
	return a.createdAt
}

// UpdatedAt returns the AccessToken's last update time.
func (a *AccessToken) UpdatedAt() time.Time {
	return a.updatedAt
}
