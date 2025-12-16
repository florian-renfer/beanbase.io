package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/domain"
	"github.com/google/uuid"
)

type (
	// CreateCoffeeRoasterInput represents the input data for creating a coffee roaster.
	CreateCoffeeRoasterInput struct {
		Name          string `json:"name"`
		OnlineShopURL string `json:"online_shop_url"`
	}

	// CreateCoffeeRoasterOutput represents the output data after creating a coffee roaster.
	CreateCoffeeRoasterOutput struct {
		Id            uuid.UUID `json:"id"`
		Name          string    `json:"name"`
		OnlineShopURL string    `json:"online_shop_url"`
		CreatedAt     time.Time `json:"created_at"`
		UpdatedAt     time.Time `json:"update_at"`
	}

	// CreateCoffeeRoasterPresenter formats the output for the create coffee roaster use case.
	CreateCoffeeRoasterPresenter interface {
		Output(domain.CoffeeRoaster) CreateCoffeeRoasterOutput
	}

	// CreateCoffeeRoasterUseCase defines the interface for the create coffee roaster use case.
	CreateCoffeeRoasterUseCase interface {
		Execute(context.Context, CreateCoffeeRoasterInput) (CreateCoffeeRoasterOutput, error)
	}

	// createCoffeeRoasterInteractor implements the CreateCoffeeRoasterUseCase.
	createCoffeeRoasterInteractor struct {
		repo       domain.CoffeeRoasterRepository
		presenter  CreateCoffeeRoasterPresenter
		ctxTimeout time.Duration
	}
)

// NewCreateCoffeeRoasterInteractor creates a new instance of createCoffeeRoasterInteractor.
func NewCreateCoffeeRoasterInteractor(repo domain.CoffeeRoasterRepository, presenter CreateCoffeeRoasterPresenter, ctxTimeout time.Duration) CreateCoffeeRoasterUseCase {
	return createCoffeeRoasterInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

// Execute runs the create coffee roaster use case.
func (uc createCoffeeRoasterInteractor) Execute(ctx context.Context, input CreateCoffeeRoasterInput) (CreateCoffeeRoasterOutput, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.ctxTimeout)
	defer cancel()

	var coffeeRoaster = domain.NewCoffeeRoaster(
		uuid.New(),
		input.Name,
		input.OnlineShopURL,
		time.Now(),
		time.Now(),
	)

	coffeeRoaster, err := uc.repo.Create(ctx, coffeeRoaster)
	if err != nil {
		return uc.presenter.Output(domain.CoffeeRoaster{}), err
	}

	return uc.presenter.Output(coffeeRoaster), nil
}
