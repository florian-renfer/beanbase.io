package router

import (
	"errors"
	"time"

	"github.com/florian-renfer/beanbase.io/internal/adapter/repository"
)

type Server interface {
	Listen()
}

type Port int64

var (
	errInvalidWebServerInstance = errors.New("invalid web server instance")
)

const (
	InstanceNet int = iota
)

func NewHttpRouterFactory(instance int, port Port, db repository.SQL, timeout time.Duration) (Server, error) {
	switch instance {
	case InstanceNet:
		return newNetHttpRouter(port, db, timeout), nil
	default:
		return nil, errInvalidWebServerInstance
	}
}
