package router

import (
	"fmt"
	"net/http"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/adapter/api/action"
	"github.com/florian-renfer/beanbase.io/internal/adapter/presenter"
	"github.com/florian-renfer/beanbase.io/internal/adapter/repository"
	"github.com/florian-renfer/beanbase.io/internal/usecase"
)

type netHttpRouter struct {
	server     *http.Server
	port       Port
	db         repository.SQL
	ctxTimeout time.Duration
}

func newNetHttpRouter(port Port, db repository.SQL, ctxTimeout time.Duration) Server {
	return &netHttpRouter{
		server:     nil,
		port:       port,
		db:         db,
		ctxTimeout: ctxTimeout,
	}
}

// TODO: implement graceful shutdown
func (r *netHttpRouter) Listen() {
	mux := http.NewServeMux()
	r.registerRoutes(mux)

	r.server = &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 15 * time.Second,
		Addr:         fmt.Sprintf(":%d", r.port),
		Handler:      mux,
	}

	r.server.ListenAndServe()
}

func (r netHttpRouter) registerRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /api/v1/coffee-roasters", r.coffeeRoastersCreate())
}

func (r netHttpRouter) coffeeRoastersCreate() http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		var (
			uc = usecase.NewCreateCoffeeRoasterInteractor(
				repository.NewAccountSQL(r.db),
				presenter.NewCreateCoffeeRoasterPresenter(),
				r.ctxTimeout,
			)

			act = action.NewCreateCoffeeRoasterAction(uc)
		)

		act.Execute(res, req)
	}
}
