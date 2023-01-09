package errs

import "errors"

var (
	ErrBruceNotInit                      = errors.New("bruce not init")
	ErrBruceTypeNotSupported             = errors.New("type not supported")
	ErrGCPAccountError                   = errors.New("GCP account connection error")
	ErrPSNContentNotFound                = errors.New("content not found")
	ErrInvalidPSNContent                 = errors.New("PSN content is invalid")
	ErrPushMessageContentStoreNotDefined = errors.New("push Message content store type not defined")
	ErrPushMessageContentStoreNotInit    = errors.New("push message content store is not init")
	ErrDeviceStoreNotInit                = errors.New("device store not init")
	ErrDeviceStoreNotDefined             = errors.New("device store not defined")
	ErrFCMClientNotInit                  = errors.New("fcm client is not init")
	ErrFCMClientNotDefined               = errors.New("fcm client is not defined")
)
