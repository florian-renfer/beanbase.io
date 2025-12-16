package repository

import (
	"context"
	"errors"

	"github.com/florian-renfer/beanbase.io/internal/domain"
)

// coffeeRoasterSQL provides methods to interact with coffee roaster data in a SQL database.
type coffeeRoasterSQL struct {
	db SQL
}

// NewAccountSQL creates a new coffeeRoasterSQL with the given SQL database.
func NewAccountSQL(db SQL) coffeeRoasterSQL {
	return coffeeRoasterSQL{
		db: db,
	}
}

// Create inserts a new CoffeeRoaster into the database.
func (s coffeeRoasterSQL) Create(ctx context.Context, coffeeRoaster domain.CoffeeRoaster) (domain.CoffeeRoaster, error) {
	var query = `
		INSERT INTO 
			coffee_roasters (id, name, online_shop_url, created_at, updated_at)
		VALUES 
			($1, $2, $3, $4, $5)
	`

	if err := s.db.ExecuteContext(
		ctx,
		query,
		coffeeRoaster.ID(),
		coffeeRoaster.Name(),
		coffeeRoaster.OnlineShopURL(),
		coffeeRoaster.CreatedAt(),
		coffeeRoaster.UpdatedAt(),
	); err != nil {
		return domain.CoffeeRoaster{}, errors.Join(err)
	}

	return coffeeRoaster, nil
}
