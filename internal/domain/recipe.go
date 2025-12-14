package domain

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

type (
	RecipeType string

	BrewingMethod string

	Recipe struct {
		id               uuid.UUID
		baristaID        uuid.UUID
		beanID           uuid.UUID
		description      *string
		type_            RecipeType
		method           BrewingMethod
		dose             float64
		yield            float64
		waterAmount      *float64
		waterTemperature *float64
		createdAt        time.Time
		updatedAt        time.Time
	}

	RecipeRepository interface {
		FindById(context.Context, uuid.UUID) (Recipe, error)
		FindByBaristaId(context.Context, uuid.UUID) ([]Recipe, error)
	}
)

const (
	RecipeTypePortafilter RecipeType = "portafilter"

	BrewingMethodEspresso    BrewingMethod = "espresso"
	BrewingMethodV60         BrewingMethod = "v60"
	BrewingMethodAeropress   BrewingMethod = "aeropress"
	BrewingMethodFrenchPress BrewingMethod = "french_press"
)

var (
	ErrRecipeNotFound = errors.New("recipe not found")
)

// ID returns the Recipe's id.
func (r *Recipe) ID() uuid.UUID {
	return r.id
}

// BaristaID returns the Recipe's barista id.
func (r *Recipe) BaristaID() uuid.UUID {
	return r.baristaID
}

// BeanID returns the Recipe's bean id.
func (r *Recipe) BeanID() uuid.UUID {
	return r.beanID
}

// Description returns the Recipe's description.
func (r *Recipe) Description() *string {
	return r.description
}

// Type returns the Recipe's type.
func (r *Recipe) Type() RecipeType {
	return r.type_
}

// Method returns the Recipe's brewing method.
func (r *Recipe) Method() BrewingMethod {
	return r.method
}

// Dose returns the Recipe's dose.
func (r *Recipe) Dose() float64 {
	return r.dose
}

// Yield returns the Recipe's yield.
func (r *Recipe) Yield() float64 {
	return r.yield
}

// WaterAmount returns the Recipe's water amount.
func (r *Recipe) WaterAmount() *float64 {
	return r.waterAmount
}

// WaterTemperature returns the Recipe's water temperature.
func (r *Recipe) WaterTemperature() *float64 {
	return r.waterTemperature
}

// CreatedAt returns the Recipe's creation time.
func (r *Recipe) CreatedAt() time.Time {
	return r.createdAt
}

// UpdatedAt returns the Recipe's last update time.
func (r *Recipe) UpdatedAt() time.Time {
	return r.updatedAt
}
