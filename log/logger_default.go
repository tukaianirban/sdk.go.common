package log

import (
	"fmt"
)

type providerdefault struct {
	isTrace  bool
	logLevel LOG_LEVEL
}

func NewDefaultLogger(isTracingEnabled bool, logLevel LOG_LEVEL) *providerdefault {

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

func (def *providerdefault) GetLogLevel() LOG_LEVEL {
	return def.logLevel
}
