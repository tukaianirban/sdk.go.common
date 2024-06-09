package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const format_datetime string = "02-01-2006 15:04:05"

func getDateTime() string {

	return time.Now().Format(format_datetime)
}

func getLogLine(message string, severity string, args ...interface{}) string {
	return fmt.Sprintf("%s [%s] %v \n", getDateTime(), severity, fmt.Sprintf(message, args...))
}

func getLogLineTrace(message string, severity string, args ...interface{}) string {

	pc, filename, linenumber, _ := runtime.Caller(2)

	// extract only the filename from the entire filepath
	parts := strings.Split(filename, "/")
	filename = parts[len(parts)-1]

	// extract only the function name from the whole name
	parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
	callerfunc := parts[len(parts)-1]

	// loggerInstance.Printf(fmt.Sprintf("%s [%s] %s:%s:%d %v\n", getDateTime(true), SEVERITY_TRACER, filename, callerfunc, linenumber, fmt.Sprintf(message, args...)))

	return fmt.Sprintf("%s [%s] %s:%s:%d %v\n", getDateTime(), severity, filename, callerfunc, linenumber, fmt.Sprintf(message, args...))
}

/**
Exposed Functions from the package
These functions create the entire log message as it should be, including all its contents.
The different logging providers only write the message to the appropriate output.
**/

func Info(message string, args ...interface{}) {

	if loggerInstance.IsTraceEnabled() {
		loggerInstance.Printf(getLogLineTrace(message, LEVEL_INFO, args...))
	} else {
		loggerInstance.Printf(getLogLine(message, LEVEL_INFO, args...))
	}
}

func Warning(message string, args ...interface{}) {

	if loggerInstance.IsTraceEnabled() {
		loggerInstance.Printf(getLogLineTrace(message, LEVEL_WARNING, args...))
	} else {
		loggerInstance.Printf(getLogLine(message, LEVEL_WARNING, args...))
	}
}

func Error(message string, args ...interface{}) {

	if loggerInstance.IsTraceEnabled() {
		loggerInstance.Printf(getLogLineTrace(message, LEVEL_ERROR, args...))
	} else {
		loggerInstance.Printf(getLogLine(message, LEVEL_ERROR, args...))
	}
}

func Alert(message string, args ...interface{}) {

	if loggerInstance.IsTraceEnabled() {
		loggerInstance.Printf(getLogLineTrace(message, LEVEL_ALERT, args...))
	} else {
		loggerInstance.Printf(getLogLine(message, LEVEL_ALERT, args...))
	}
}

func Fatal(message string, args ...interface{}) {

	if loggerInstance.IsTraceEnabled() {
		loggerInstance.Printf(getLogLineTrace(message, LEVEL_FATAL, args...))
	} else {
		loggerInstance.Printf(getLogLine(message, LEVEL_FATAL, args...))
	}
	os.Exit(1)
}

func Debug(message string, args ...interface{}) {

	if loggerInstance.IsTraceEnabled() {
		loggerInstance.Printf(getLogLineTrace(message, LEVEL_DEBUG, args...))
	} else {
		loggerInstance.Printf(getLogLine(message, LEVEL_DEBUG, args...))
	}
}

// fetch the last error encountered, if any
func GetLastError() error {
	return loggerInstance.GetLastError()
}
