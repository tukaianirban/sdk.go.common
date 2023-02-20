package logdefault

import (
	"fmt"
)

type providerdefault struct {
}

func New() *providerdefault {

	return &providerdefault{}
}

func (def *providerdefault) GetLastError() error {
	return nil
}

func (def *providerdefault) Printf(message string, args ...interface{}) {
	fmt.Printf(message, args...)
}
