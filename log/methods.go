package log

import (
	"fmt"
	"os"
	"runtime"
	"strings"
	"time"
)

const format_datetime string = "02-01-2006 15:04:05"
const format_datetime_debug string = "02-01-2006 15:04:05.000"

func getDateTime(debug bool) string {

	currentTime := time.Now()

	if debug {
		return currentTime.Format(format_datetime_debug)
	} else {
		return currentTime.Format(format_datetime)
	}
}

// the instance through which exposed functions are triggered
var loggerInstance LoggerInterface

/************ Exposed functions **************/
/**
Exposed Functions from the package
These functions create the entire log message as it should be, including all its contents.
The different logging providers only write the message to the appropriate output.
**/

// prints the log message and adds a newline at the end
func Print(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_DEFAULT, fmt.Sprintf(message, args...)))
}

func Info(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_INFO, fmt.Sprintf(message, args...)))
}

func Notice(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_NOTICE, fmt.Sprintf(message, args...)))
}

func Warning(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_WARNING, fmt.Sprintf(message, args...)))
}

func Error(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_ERROR, fmt.Sprintf(message, args...)))
}

func Critical(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_CRITICAL, fmt.Sprintf(message, args...)))
}

func Alert(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_ALERT, fmt.Sprintf(message, args...)))
}

func Emergency(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_EMERGENCY, fmt.Sprintf(message, args...)))
}

func Fatal(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(false), SEVERITY_FATAL, fmt.Sprintf(message, args...)))
	os.Exit(1)
}

func Debug(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	loggerInstance.Printf(fmt.Sprintf("%s [%s] %v\n", getDateTime(true), SEVERITY_DEBUG, fmt.Sprintf(message, args...)))
}

func Trace(message string, args ...interface{}) {

	if loggerInstance == nil {
		pc, filename, linenumber, _ := runtime.Caller(1)

		// extract only the filename from the entire filepath
		parts := strings.Split(filename, "/")
		filename = parts[len(parts)-1]

		// extract only the function name from the whole name
		parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
		callerfunc := parts[len(parts)-1]

		fmt.Printf("Fatal: logger module called from %s : %s : %d not initialised yet", filename, callerfunc, linenumber)
		return
	}
	pc, filename, linenumber, _ := runtime.Caller(1)

	// extract only the filename from the entire filepath
	parts := strings.Split(filename, "/")
	filename = parts[len(parts)-1]

	// extract only the function name from the whole name
	parts = strings.Split(runtime.FuncForPC(pc).Name(), ".")
	callerfunc := parts[len(parts)-1]

	loggerInstance.Printf(fmt.Sprintf("%s [%s] %s:%s:%d %v\n", getDateTime(true), SEVERITY_TRACER, filename, callerfunc, linenumber, fmt.Sprintf(message, args...)))
}

func TraceVerbose(message string, args ...interface{}) {

	if loggerInstance == nil {
		fmt.Println("Fatal: logger module not initialised yet")
		return
	}
	pc, filepath, linenumber, _ := runtime.Caller(1)

	// extract only the function name from the whole name
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	callerfunc := parts[len(parts)-1]

	loggerInstance.Printf(fmt.Sprintf("%s [%s] %s:%s:%d %v\n", getDateTime(true), SEVERITY_TRACER, filepath, callerfunc, linenumber, fmt.Sprintf(message, args...)))
}

// fetch the last error encountered, if any
func GetLastError() error {
	return loggerInstance.GetLastError()
}
