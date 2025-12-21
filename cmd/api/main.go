package main

import (
	"time"

	"github.com/florian-renfer/beanbase.io/internal/infrastructure"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/database"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/logging"
	"github.com/florian-renfer/beanbase.io/internal/infrastructure/router"
)

func main() {
	infrastructure.NewApp().
		Logger(logging.InstanceSlog).
		Persistence(database.InstancePostgres).
		Timeout(time.Second * 10).
		WebServerPort("4000").
		WebServer(router.InstanceNet).
		Start()
}
