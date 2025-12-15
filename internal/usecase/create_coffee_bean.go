package usecase

import (
	"context"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/domain"
)

type (
	CreateCoffeeBeanInput struct{}

	CreateCoffeeBeanOutput struct{}

	CreateCoffeeBeanPresenter interface {
		Output(domain.CoffeeBean) CreateCoffeeBeanOutput
	}

	CreateCoffeeBeanUseCase interface {
		Execute(context.Context, CreateCoffeeBeanInput) (CreateCoffeeBeanOutput, error)
	}

	createCoffeeBeanInteractor struct {
		repo       domain.BeanRepository
		presenter  CreateCoffeeBeanPresenter
		ctxTimeout time.Duration
	}
)

func NewCreateCoffeeBeanInteractor(repo domain.BeanRepository, presenter CreateCoffeeBeanPresenter, ctxTimeout time.Duration) CreateCoffeeBeanUseCase {
	return createCoffeeBeanInteractor{
		repo:       repo,
		presenter:  presenter,
		ctxTimeout: ctxTimeout,
	}
}

func (uc createCoffeeBeanInteractor) Execute(context.Context, CreateCoffeeBeanInput) (CreateCoffeeBeanOutput, error) {
	return CreateCoffeeBeanOutput{}, nil
}
