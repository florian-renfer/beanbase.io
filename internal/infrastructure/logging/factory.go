package logging

import (
	"errors"

	"github.com/florian-renfer/beanbase.io/internal/adapter/logger"
)

const (
	InstanceSlog = iota
)

var (
	errInvalidLoggingInstance = errors.New("invalid logging instance")
)

// NewLoggingFactory createas a new logger based on the given instance type.
func NewLoggingFactory(instance int) (logger.Logger, error) {
	switch instance {
	case InstanceSlog:
		return newSlogLogger()
	default:
		return nil, errInvalidLoggingInstance
	}
}
