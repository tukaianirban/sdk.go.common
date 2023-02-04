package errs

import "errors"

var (
	ErrBruceNotInit          = errors.New("bruce not init")
	ErrBruceTypeNotSupported = errors.New("type not supported")
	ErrGCPAccountError       = errors.New("GCP account connection error")
	ErrDeviceStoreNotInit    = errors.New("device store not init")
	ErrDeviceStoreNotDefined = errors.New("device store not defined")
	ErrFCMClientNotInit      = errors.New("fcm client is not init")
	ErrFCMClientNotDefined   = errors.New("fcm client is not defined")
)
