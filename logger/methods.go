package logger

import (
	"fmt"
	"time"
)

const FORMAT_DATETIME_UNIVERSAL string = "02-01-2006 15:04:05"

// the instance through which external modules will trigger the logs
var loggerInstance LoggerInterface

/**
Interface definition that each logger provider (local, gcp, aws, etc) must support
**/
type LoggerInterface interface {

	// the message is output'ed as-is into the active provider
	Printf(message string, args ...interface{})
	Println(message string)

	// includes a DEBUG marker in the log message output'ed
	Debugf(message string, args ...interface{})
	Debug(message string)

	// includes a INFO marker in the log message output'ed
	Infof(message string, args ...interface{})
	Info(message string)

	// includes a NOTICE marker in the log message output'ed
	Noticef(message string, args ...interface{})
	Notice(message string)

	// includes a WARNING marker in the log message output'ed
	Warningf(message string, args ...interface{})
	Warning(message string)

	// includes a ERROR marker in the log message output'ed
	Errorf(message string, args ...interface{})
	Error(message string)

	// includes a CRITICAL marker in the log message output'ed
	Criticalf(message string, args ...interface{})
	Critical(message string)

	// includes a ALERT marker in the log message output'ed
	Alertf(message string, args ...interface{})
	Alert(message string)

	// includes a EMERGENCY marker in the log message output'ed
	Emergencyf(message string, args ...interface{})
	Emergency(message string)

	// includes a FATAL marker in the log message output'ed
	Fatalf(message string, args ...interface{})
	Fatal(message string)

	// include trace information (function name where the log was invoked from )
	Tracef(message string, args ...interface{})
	Trace(message string)

	// include verbose trace information (function, filename where the log was invoked from )
	TracefVerbose(message string, args ...interface{})
	TraceVerbose(message string)
}

//
// modifies the log message to include default params in the log message
// for now: date time is added in the beginning of the message
//
func modLogMessage(message string, args ...interface{}) string {

	return time.Now().Format(FORMAT_DATETIME_UNIVERSAL) + " " + fmt.Sprintf(message, args...)
}

/************ Exposed functions **************/
func Printf(message string, args ...interface{}) {

	loggerInstance.Printf(modLogMessage(message, args))
}

// adds a newline at the end of the log message generated from mod
func Println(message string) {

	loggerInstance.Printf(fmt.Sprintf("%v\n", modLogMessage(message)))
}

func Debugf(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Debug(message string, args ...interface{}) {

}

func Infof(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Info(message string, args ...interface{}) {

}

func Noticef(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Notice(message string, args ...interface{}) {

}

func Warningf(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Warning(message string, args ...interface{}) {

}

func Errorf(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Error(message string, args ...interface{}) {

}

func Criticalf(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Critical(message string, args ...interface{}) {

}

func Alertf(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Alert(message string, args ...interface{}) {

}

func Emergencyf(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Emergency(message string, args ...interface{}) {

}

func Fatalf(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Fatal(message string, args ...interface{}) {

}

func Tracef(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func Trace(message string, args ...interface{}) {

}

func TracefVerbose(message string, args ...interface{}) {

}

// adds a newline at the end of the log message generated from mod
func TraceVerbose(message string, args ...interface{}) {

}
