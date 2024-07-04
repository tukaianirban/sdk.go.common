package logdefault

import (
	"fmt"

	"github.com/tukaianirban/sdk.go.common/log/defs"
)

type providerdefault struct {
	isTrace  bool
	logLevel defs.LOG_LEVEL
}

func New(isTracingEnabled bool, logLevel defs.LOG_LEVEL) *providerdefault {

	return &providerdefault{
		isTrace:  isTracingEnabled,
		logLevel: logLevel,
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

func (def *providerdefault) GetLogLevel() defs.LOG_LEVEL {
	return def.logLevel
}
