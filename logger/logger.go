/**
The logger module exposes a set of functions to log messages of different types/levels.
It exposes functions of different severity levels which are backed by a logger store depending on how the logger module is init'ed.
**/
package logger

import "log"

//
// the different modes in which logger can function
//
const (
	MODE_DEFAULT = iota
	MODE_LOCALFILE
	MODE_GCP
	MODE_AWS
)

// todo : move this package to sdk.go.common with support for logging locally or with a cloud provider

// type Logger interface {
// 	Log(name string, sev Sevlevel)
// }

// type Sevlevel uint8

// const (
// 	Default Sevlevel = iota
// 	Debug
// 	Info
// 	Notice
// 	Warning
// 	Error
// 	Critical
// 	Alert
// 	Emergency
// )

// var SevLevelStrings [9]string = [9]string{"DEF", "DEBUG", "INFO",
// 	"NOTIC", "WARN", "ERR", "CRIT", "ALERT", "EMERG"}

// const FORMAT_DATETIME_UNIVERSAL string = "02-01-2006 15:04:05"

// type LogMessage struct {
// 	Name     string
// 	Severity Sevlevel
// }

// var loggerObj Logger

func Init(mode int, args ...string) {

	switch mode {
	case MODE_DEFAULT:
		// default logger implemented as wrapper on "log" package

	case MODE_LOCALFILE:
		// logs into a local file, filename and path should come in as args
		if len(args) < 1 {
			log.Fatalln("logger init failed for localfile mode, reason=local file not provided")
			break
		}

	case MODE_GCP:
		// logs into a GCP logging instance
		// arg should be the path of a configuration file containing gcp project details
		if len(args) < 1 {
			log.Fatalln("logger init failed for GCP mode, reason=GCP config file not provided")
			break
		}

	case MODE_AWS:
		// logs into a AWS service instance for logging
		// arg should be the path of a configuration file containing AWS service details
		if len(args) < 1 {
			log.Fatalln("logger init failed for AWS mode, reason=AWS config file not provided")
			break
		}

	default:
		log.Fatalln("logger init failed, unrecognised mode specified")
	}
}
