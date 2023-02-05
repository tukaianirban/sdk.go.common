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

	// includes a INFO marker in the log message output'ed
	Infof(message string, args ...interface{})

	// includes a NOTICE marker in the log message output'ed
	Noticef(message string, args ...interface{})

	// includes a WARNING marker in the log message output'ed
	Warningf(message string, args ...interface{})

	// includes a ERROR marker in the log message output'ed
	Errorf(message string, args ...interface{})

	// includes a CRITICAL marker in the log message output'ed
	Criticalf(message string, args ...interface{})

	// includes a ALERT marker in the log message output'ed
	Alertf(message string, args ...interface{})

	// includes a EMERGENCY marker in the log message output'ed
	Emergencyf(message string, args ...interface{})

	// includes a FATAL marker in the log message output'ed
	Fatalf(message string, args ...interface{})

	// includes verbose info (longfile name, time+usec, etc) in the log message output'ed
	// preferably the provider should use a different type of logging object underneath for debugging
	Debugf(message string, args ...interface{})

	// include trace information (function name where the log was invoked from )
	// preferably the provider should use a different type of logging object underneath for tracer info
	Tracef(message string, args ...interface{})

	// include verbose trace information (function, filename where the log was invoked from )
	// preferably the provider should use a different type of logging object underneath for tracer info
	TracefVerbose(message string, args ...interface{})
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

	loggerInstance.Printf(fmt.Sprintf("%s:%v", SEVERITY_DEFAULT, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Println(message string) {

	loggerInstance.Printf(fmt.Sprintf("%s:%v", SEVERITY_DEFAULT, fmt.Sprintf("%v\n", message)))
}

func Infof(message string, args ...interface{}) {

	loggerInstance.Infof(fmt.Sprintf("%s:%v", SEVERITY_INFO, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Info(message string) {

	loggerInstance.Infof(fmt.Sprintf("%s:%v", SEVERITY_INFO, fmt.Sprintf("%v\n", message)))
}

func Noticef(message string, args ...interface{}) {

	loggerInstance.Noticef(fmt.Sprintf("%s:%v", SEVERITY_NOTICE, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Notice(message string) {

	loggerInstance.Noticef(fmt.Sprintf("%s:%v", SEVERITY_NOTICE, fmt.Sprintf("%v\n", message)))
}

func Warningf(message string, args ...interface{}) {

	loggerInstance.Warningf(fmt.Sprintf("%s:%v", SEVERITY_WARNING, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Warning(message string) {

	loggerInstance.Warningf(fmt.Sprintf("%s:%v", SEVERITY_WARNING, fmt.Sprintf("%v\n", message)))
}

func Errorf(message string, args ...interface{}) {

	loggerInstance.Errorf(fmt.Sprintf("%s:%v", SEVERITY_ERROR, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Error(message string) {

	loggerInstance.Errorf(fmt.Sprintf("%s:%v", SEVERITY_ERROR, fmt.Sprintf("%v\n", message)))
}

func Criticalf(message string, args ...interface{}) {

	loggerInstance.Criticalf(fmt.Sprintf("%s:%v", SEVERITY_CRITICAL, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Critical(message string) {

	loggerInstance.Criticalf(fmt.Sprintf("%s:%v", SEVERITY_CRITICAL, fmt.Sprintf("%v\n", message)))
}

func Alertf(message string, args ...interface{}) {

	loggerInstance.Alertf(fmt.Sprintf("%s:%v", SEVERITY_ALERT, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Alert(message string) {

	loggerInstance.Alertf(fmt.Sprintf("%s:%v", SEVERITY_ALERT, fmt.Sprintf("%v\n", message)))
}

func Emergencyf(message string, args ...interface{}) {

	loggerInstance.Emergencyf(fmt.Sprintf("%s:%v", SEVERITY_EMERGENCY, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Emergency(message string) {

	loggerInstance.Emergencyf(fmt.Sprintf("%s:%v", SEVERITY_EMERGENCY, fmt.Sprintf("%v\n", message)))
}

func Fatalf(message string, args ...interface{}) {

	loggerInstance.Fatalf(fmt.Sprintf("%s:%v", SEVERITY_FATAL, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Fatal(message string) {

	loggerInstance.Fatalf(fmt.Sprintf("%s:%v", SEVERITY_FATAL, fmt.Sprintf("%v\n", message)))
}

func Debugf(message string, args ...interface{}) {

	loggerInstance.Debugf(fmt.Sprintf("%s:%v", SEVERITY_DEBUG, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Debug(message string) {

	loggerInstance.Debugf(fmt.Sprintf("%s:%v", SEVERITY_DEBUG, fmt.Sprintf("%v\n", message)))
}

func Tracef(message string, args ...interface{}) {

	loggerInstance.Tracef(fmt.Sprintf("%s:%v", SEVERITY_TRACER, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func Trace(message string) {

	loggerInstance.Tracef(fmt.Sprintf("%s:%v", SEVERITY_TRACER, fmt.Sprintf("%v\n", message)))
}

func TracefVerbose(message string, args ...interface{}) {

	loggerInstance.TracefVerbose(fmt.Sprintf("%s:%v", SEVERITY_TRACER, fmt.Sprintf(message, args...)))
}

// adds a newline at the end of the log message generated from mod
func TraceVerbose(message string) {

	loggerInstance.TracefVerbose(fmt.Sprintf("%s:%v", SEVERITY_TRACER, fmt.Sprintf("%v\n", message)))
}
