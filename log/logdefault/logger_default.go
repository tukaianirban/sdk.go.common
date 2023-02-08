package logdefault

import (
	"log"
	"os"
)

type providerdefault struct {
	inst      *log.Logger
	instDebug *log.Logger
}

func New() *providerdefault {

	return &providerdefault{
		inst:      log.New(nil, "", log.Ldate|log.Ltime|log.Lshortfile),
		instDebug: log.New(nil, "", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile),
	}
}

func (def *providerdefault) Printf(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Infof(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Noticef(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Warningf(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Errorf(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Criticalf(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Alertf(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Emergencyf(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
}

func (def *providerdefault) Fatalf(message string, args ...interface{}) {
	def.inst.Printf(message, args...)
	os.Exit(1)
}

func (def *providerdefault) Debugf(message string, args ...interface{}) {
	def.instDebug.Printf(message, args...)
}

func (def *providerdefault) Tracef(message string, args ...interface{}) {
	def.instDebug.Printf(message, args...)
}

func (def *providerdefault) TracefVerbose(message string, args ...interface{}) {
	def.instDebug.Printf(message, args...)
}
