package infrastructure

import (
	"os"
	"strconv"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/adapter/repository"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/database"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/router"
)

type application struct {
	ctxTimeout    time.Duration
	dbSQL         repository.SQL
	webServerPort router.Port
	webServer     router.Server
}

func NewApp() *application {
	return &application{}
}

func (app *application) Timeout(ctxTimeout time.Duration) *application {
	app.ctxTimeout = ctxTimeout
	return app
}

func (app *application) Persistence(instance int) *application {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		os.Exit(1)
	}

	app.dbSQL = db
	return app
}

func (app *application) WebServerPort(port string) *application {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		os.Exit(1)
	}

	app.webServerPort = router.Port(p)
	return app
}

func (app *application) WebServer(instance int) *application {
	s, err := router.NewHttpRouterFactory(router.InstanceNet, app.webServerPort, app.dbSQL, app.ctxTimeout)
	if err != nil {
		os.Exit(1)
	}

	app.webServer = s
	return app
}

func (app *application) Start() {
	app.webServer.Listen()
}
