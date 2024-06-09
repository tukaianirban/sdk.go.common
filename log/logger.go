/**
The logger module exposes a set of functions to log messages of different types/levels.
It exposes functions of different severity levels which are backed by a logger store depending on how the logger module is init'ed.
**/
package log

import "github.com/tukaianirban/sdk.go.common/log/logdefault"

//
// severity-based prefixes definition
//
const (
	LEVEL_DEBUG   = "D"
	LEVEL_INFO    = "I"
	LEVEL_WARNING = "W"
	LEVEL_ERROR   = "E"
	LEVEL_ALERT   = "A"
	LEVEL_FATAL   = "F"
)

//
// the global instance for calls to underlying struct
//
var loggerInstance logger

func Init(flagTrace bool) {

	loggerInstance = logdefault.New(flagTrace)
}

/**
Interface definition that each log provider must support
**/
type logger interface {
	IsTraceEnabled() bool
	Printf(message string, args ...interface{})
	GetLastError() error
}
