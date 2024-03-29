/**
The logger module exposes a set of functions to log messages of different types/levels.
It exposes functions of different severity levels which are backed by a logger store depending on how the logger module is init'ed.
**/
package log

import (
	"fmt"
	"os"
	"reflect"

	"github.com/tukaianirban/sdk.go.common/log/logdefault"
	"github.com/tukaianirban/sdk.go.common/log/loglocal"
)

//
// the different modes in which logger can function
//
const (
	MODE_DEFAULT = iota
	MODE_LOCALFILE
	MODE_GCP
	MODE_AWS
)

//
// severity-based prefixes definition
//
const (
	SEVERITY_DEFAULT   = "DEF"
	SEVERITY_DEBUG     = "DEB"
	SEVERITY_INFO      = "INF"
	SEVERITY_NOTICE    = "NOT"
	SEVERITY_WARNING   = "WAR"
	SEVERITY_ERROR     = "ERR"
	SEVERITY_CRITICAL  = "CRI"
	SEVERITY_ALERT     = "ALE"
	SEVERITY_EMERGENCY = "EME"
	SEVERITY_FATAL     = "FAT"
	SEVERITY_TRACER    = "TRA"
)

const PREFIX_SEVERITY_SEPARATOR = "/"

func Init(mode int, args ...string) {

	switch mode {
	case MODE_DEFAULT:
		loggerInstance = logdefault.New()

	case MODE_LOCALFILE:
		// logs into a local file, filename and path should come in as args
		if len(args) < 1 {
			fmt.Println("logger init failed for localfile mode, reason=local file not provided")
			os.Exit(1)
			break
		}

		loggerInstance = loglocal.New(args[0])

	case MODE_GCP:
		// logs into a GCP logging instance
		// arg should be the path of a configuration file containing gcp project details
		if len(args) < 1 {
			fmt.Println("logger init failed for GCP mode, reason=GCP config file not provided")
			os.Exit(1)
			break
		}

	case MODE_AWS:
		// logs into a AWS service instance for logging
		// arg should be the path of a configuration file containing AWS service details
		if len(args) < 1 {
			fmt.Println("logger init failed for AWS mode, reason=AWS config file not provided")
			os.Exit(1)
			break
		}

	default:
		fmt.Println("logger init failed, unrecognised mode specified")
		os.Exit(1)
	}

	if reflect.ValueOf(loggerInstance).IsNil() {
		fmt.Println("failed to initialise logger module")
		os.Exit(1)
	}
}

/**
Interface definition that each logger provider (local, gcp, aws, etc) must support
**/
type LoggerInterface interface {
	Printf(message string, args ...interface{})
	GetLastError() error
}
