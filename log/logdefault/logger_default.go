package logdefault

import (
	"fmt"
)

type providerdefault struct {
}

func New() *providerdefault {

	return &providerdefault{}
}

func (def *providerdefault) Printf(message string, args ...interface{}) {
	fmt.Printf(message, args...)
}
