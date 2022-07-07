package errs

import "errors"

var (
	ErrBruceNotInit          = errors.New("bruce not init")
	ErrBruceTypeNotSupported = errors.New("type not supported")
	ErrGCPAccountError       = errors.New("GCP account connection error")
)
