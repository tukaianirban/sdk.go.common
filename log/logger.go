/**
The logger module exposes a set of functions to log messages of different types/levels.
It exposes functions of different severity levels which are backed by a logger store depending on how the logger module is init'ed.
**/
package log

import (
	"github.com/tukaianirban/sdk.go.common/log/defs"
	"github.com/tukaianirban/sdk.go.common/log/logdefault"
)

//
// the global instance for calls to underlying struct
//
var loggerInstance logger

func Init(flagTrace bool, flags ...int) {

	logLevel := defs.LEVEL_INFO
	if len(flags) > 0 {
		logLevel = defs.LOG_LEVEL(flags[0])
	}

	loggerInstance = logdefault.New(flagTrace, logLevel)
}

/**
Interface definition that each log provider must support
**/
type logger interface {
	IsTraceEnabled() bool
	GetLogLevel() defs.LOG_LEVEL
	Printf(message string, args ...interface{})
	GetLastError() error
}
