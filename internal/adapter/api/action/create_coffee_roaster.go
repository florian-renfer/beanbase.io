package action

import (
	"encoding/json"
	"net/http"

	"github.com/florian-renfer/beanbase.io/internal/adapter/api/response"
	"github.com/florian-renfer/beanbase.io/internal/usecase"
)

type CreateCoffeeRoasterAction struct {
	uc usecase.CreateCoffeeRoasterUseCase
}

func NewCreateCoffeeRoasterAction(uc usecase.CreateCoffeeRoasterUseCase) CreateCoffeeRoasterAction {
	return CreateCoffeeRoasterAction{uc: uc}
}

// Execute creates a coffee roaster via HTTP request.
// Returns 400 for decode errors, 500 for use case errors, 201 on success.
func (a CreateCoffeeRoasterAction) Execute(w http.ResponseWriter, r *http.Request) {
	var input usecase.CreateCoffeeRoasterInput
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		response.NewError(err, http.StatusBadRequest).Send(w)
		return
	}
	defer r.Body.Close()

	output, err := a.uc.Execute(r.Context(), input)
	if err != nil {
		response.NewError(err, http.StatusInternalServerError).Send(w)
		return
	}

	response.NewSuccess(output, http.StatusCreated).Send(w)
}
