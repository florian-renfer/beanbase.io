package infrastructure

import (
	"os"
	"strconv"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/adapter/logger"
	"github.com/florian-renfer/beanbase.io/internal/adapter/repository"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/database"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/logging"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/router"
)

type application struct {
	ctxTimeout    time.Duration
	logger        logger.Logger
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

func (app *application) Logger(instance int) *application {
	l, err := logging.NewLoggingFactory(instance)
	if err != nil {
		os.Exit(1)
	}

	app.logger = l
	return app
}

// TODO: graceful shutdown
func (app *application) Persistence(instance int) *application {
	db, err := database.NewDatabaseSQLFactory(instance)
	if err != nil {
		app.logger.Errorf("error connecting to persistence provider: %+v", err)
		os.Exit(1)
	}

	app.dbSQL = db
	return app
}

func (app *application) WebServerPort(port string) *application {
	p, err := strconv.ParseInt(port, 10, 64)
	if err != nil {
		app.logger.Errorf("error parsing web server port: %+v", err)
		os.Exit(1)
	}

	app.webServerPort = router.Port(p)
	return app
}

func (app *application) WebServer(instance int) *application {
	s, err := router.NewHttpRouterFactory(instance, app.webServerPort, app.dbSQL, app.ctxTimeout)
	if err != nil {
		app.logger.Errorf("error creating web server instance: %+v", err)
		os.Exit(1)
	}

	app.webServer = s
	return app
}

func (app *application) Start() {
	app.webServer.Listen()
}
