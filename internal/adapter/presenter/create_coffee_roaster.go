package presenter

import (
	"time"

	"github.com/florian-renfer/beanbase.io/internal/domain"
	"github.com/florian-renfer/beanbase.io/internal/usecase"
)

// createCoffeeRoasterPresenter implements the CreateCoffeeRoasterPresenter interface.
type createCoffeeRoasterPresenter struct{}

// NewCreateCoffeeRoasterPresenter returns a new CreateCoffeeRoasterPresenter.
func NewCreateCoffeeRoasterPresenter() usecase.CreateCoffeeRoasterPresenter {
	return createCoffeeRoasterPresenter{}
}

// Output formats a CoffeeRoaster domain model into a CreateCoffeeRoasterOutput DTO.
func (p createCoffeeRoasterPresenter) Output(coffeeRoaster domain.CoffeeRoaster) usecase.CreateCoffeeRoasterOutput {
	return usecase.CreateCoffeeRoasterOutput{
		ID:            coffeeRoaster.ID(),
		Name:          coffeeRoaster.Name(),
		OnlineShopURL: coffeeRoaster.OnlineShopURL(),
		CreatedAt:     coffeeRoaster.CreatedAt().Format(time.RFC3339),
		UpdatedAt:     coffeeRoaster.UpdatedAt().Format(time.RFC3339),
	}
}
