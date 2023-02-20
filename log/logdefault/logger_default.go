package logdefault

import (
	"fmt"
	"log"
	"os"
)

type providerdefault struct {
	inst *log.Logger
	// instDebug *log.Logger
}

func New() *providerdefault {

	// TODO: shortfile and logfile both refers to this file ( not of the using application file )
	return &providerdefault{
		inst: log.New(os.Stdout, "", log.Default().Flags()),
		// instDebug: log.New(os.Stdout, "", log.Ldate|log.Ltime|log.Lmicroseconds),
	}
}

func (def *providerdefault) Printf(message string, args ...interface{}) {
	// def.inst.Printf(message, args...)
	fmt.Printf(message, args...)
}

// func (def *providerdefault) Infof(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// }

// func (def *providerdefault) Noticef(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// }

// func (def *providerdefault) Warningf(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// }

// func (def *providerdefault) Errorf(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// }

// func (def *providerdefault) Criticalf(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// }

// func (def *providerdefault) Alertf(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// }

// func (def *providerdefault) Emergencyf(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// }

// func (def *providerdefault) Fatalf(message string, args ...interface{}) {
// 	def.inst.Printf(message, args...)
// 	os.Exit(1)
// }

// func (def *providerdefault) Debugf(message string, args ...interface{}) {
// 	def.instDebug.Printf(message, args...)
// }

// func (def *providerdefault) Tracef(message string, args ...interface{}) {
// 	pc, filename, linenumber, _ := runtime.Caller(2)

// 	// extract only the filename from the entire filepath
// 	parts := strings.Split(filename, "/")
// 	filename = parts[len(parts)-1]
// 	callerfunc := runtime.FuncForPC(pc).Name()

// 	def.instDebug.Printf("%s:%s:%d:%v", filename, callerfunc, linenumber, fmt.Sprintf(message, args...))
// }

// func (def *providerdefault) TracefVerbose(message string, args ...interface{}) {
// 	def.Tracef(message, args...)
// }
