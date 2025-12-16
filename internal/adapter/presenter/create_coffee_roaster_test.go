package presenter

import (
	"reflect"
	"testing"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/domain"
	"github.com/florian-renfer/beanbase.io/internal/usecase"
	"github.com/google/uuid"
)

const (
	coffeeRoasterId            = "813612d2-7b62-4af9-b45d-b648e11165e6"
	coffeeRoasterName          = "Example Coffee Roaster"
	coffeeRoasterOnlineShopUrl = "https://example.com"

	defaultTimestamp = "0001-01-01T00:00:00Z"
)

func TestCreateCoffeeRoasterPresenter(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name     string
		input    domain.CoffeeRoaster
		expected usecase.CreateCoffeeRoasterOutput
	}{
		{
			name:  "Should return output when valid input",
			input: domain.NewCoffeeRoaster(uuid.MustParse(coffeeRoasterId), coffeeRoasterName, coffeeRoasterOnlineShopUrl, time.Time{}, time.Time{}),
			expected: usecase.CreateCoffeeRoasterOutput{
				ID:            uuid.MustParse(coffeeRoasterId),
				Name:          coffeeRoasterName,
				OnlineShopURL: coffeeRoasterOnlineShopUrl,
				CreatedAt:     defaultTimestamp,
				UpdatedAt:     defaultTimestamp,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pre := NewCreateCoffeeRoasterPresenter()
			if got := pre.Output(tt.input); !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("[TestCase '%s'] Got: '%+v' | Want: '%+v'", tt.name, got, tt.expected)
			}
		})
	}
}
