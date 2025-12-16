package usecase

import (
	"context"
	"reflect"
	"testing"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/domain"
	"github.com/google/uuid"
)

type (
	mockCreateCoffeeRoasterRepository struct {
		result domain.CoffeeRoaster
		err    error
	}

	mockCreateCoffeeRoasterPresenter struct {
		output CreateCoffeeRoasterOutput
	}
)

const (
	coffeeRoasterId            = "813612d2-7b62-4af9-b45d-b648e11165e6"
	coffeeRoasterName          = "Example Coffee Roaster"
	coffeeRoasterOnlineShopUrl = "https://example.com"
)

func (r mockCreateCoffeeRoasterRepository) Create(_ context.Context, _ domain.CoffeeRoaster) (domain.CoffeeRoaster, error) {
	return r.result, r.err
}

func (p mockCreateCoffeeRoasterPresenter) Output(_ domain.CoffeeRoaster) CreateCoffeeRoasterOutput {
	return p.output
}

func TestCreateCoffeeRoaster(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name          string
		args          CreateCoffeeRoasterInput
		repository    domain.CoffeeRoasterRepository
		presenter     CreateCoffeeRoasterPresenter
		expected      CreateCoffeeRoasterOutput
		expectedError any
	}{
		{
			name: "Should create coffee roaster when input is valid",
			args: CreateCoffeeRoasterInput{
				Name:          coffeeRoasterName,
				OnlineShopURL: coffeeRoasterOnlineShopUrl,
			},
			repository: mockCreateCoffeeRoasterRepository{
				result: domain.NewCoffeeRoaster(
					uuid.MustParse(coffeeRoasterId),
					coffeeRoasterName,
					coffeeRoasterOnlineShopUrl,
					time.Date(2025, 12, 16, 19, 26, 43, 0, time.UTC),
					time.Date(2025, 12, 16, 19, 26, 43, 0, time.UTC),
				),
				err: nil,
			},
			presenter: mockCreateCoffeeRoasterPresenter{
				output: CreateCoffeeRoasterOutput{
					Id:            uuid.MustParse(coffeeRoasterId),
					Name:          coffeeRoasterName,
					OnlineShopURL: coffeeRoasterOnlineShopUrl,
					CreatedAt:     time.Date(2025, 12, 16, 19, 26, 43, 0, time.UTC),
					UpdatedAt:     time.Date(2025, 12, 16, 19, 26, 43, 0, time.UTC),
				},
			},
			expected: CreateCoffeeRoasterOutput{
				Id:            uuid.MustParse(coffeeRoasterId),
				Name:          coffeeRoasterName,
				OnlineShopURL: coffeeRoasterOnlineShopUrl,
				CreatedAt:     time.Date(2025, 12, 16, 19, 26, 43, 0, time.UTC),
				UpdatedAt:     time.Date(2025, 12, 16, 19, 26, 43, 0, time.UTC),
			},
			expectedError: nil,
		},
		{
			name: "Should return error when database returns an error",
			args: CreateCoffeeRoasterInput{
				Name:          coffeeRoasterName,
				OnlineShopURL: coffeeRoasterOnlineShopUrl,
			},
			repository: mockCreateCoffeeRoasterRepository{
				result: domain.CoffeeRoaster{},
				err:    domain.ErrCoffeeRoasterExistsAlready,
			},
			presenter:     mockCreateCoffeeRoasterPresenter{},
			expected:      CreateCoffeeRoasterOutput{},
			expectedError: "coffee roaster exists already",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var uc = NewCreateCoffeeRoasterInteractor(tt.repository, tt.presenter, time.Second)

			result, err := uc.Execute(context.TODO(), tt.args)
			if (err != nil) && (err.Error() != tt.expectedError) {
				t.Errorf("[TestCase '%s'] Result: '%v' | ExpectedError: '%v'", tt.name, err, tt.expectedError)
			}

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("[TestCase '%s'] Result: '%v' | Expected: '%v'", tt.name, result, tt.expected)
			}
		})
	}
}
