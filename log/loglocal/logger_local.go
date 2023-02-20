package loglocal

import (
	"fmt"
	"os"
)

type providerlocal struct {
	filename   string
	filewriter *os.File
	lastError  error
}

func New(filepath string) *providerlocal {

	f, err := os.OpenFile(filepath, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("failed to open file")
		return nil
	}

	return &providerlocal{
		filename:   filepath,
		filewriter: f,
	}
}

func (def *providerlocal) GetLastError() error {
	return def.lastError
}

func (def *providerlocal) Printf(message string, args ...interface{}) {

	if _, err := fmt.Fprintf(def.filewriter, message, args...); err != nil {
		def.lastError = err
	} else {
		def.filewriter.Sync()
	}
}
