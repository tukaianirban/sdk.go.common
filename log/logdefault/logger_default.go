package logdefault

import (
	"fmt"
)

type providerdefault struct {
	isTrace bool
}

func New(isTracingEnabled bool) *providerdefault {

	return &providerdefault{
		isTrace: isTracingEnabled,
	}
}

func (def *providerdefault) GetLastError() error {
	return nil
}

func (def *providerdefault) Printf(message string, args ...interface{}) {
	fmt.Printf(message, args...)
}

func (def *providerdefault) IsTraceEnabled() bool {
	return def.isTrace
}
