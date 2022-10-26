package custom_error

import "errors"

var ErrRegisterFail = errors.New("REGISTER_FAILED")
var ErrInternalServer = errors.New("INTERNAL_SERVER")
var ErrNotFound = errors.New("NOT_FOUND")
var ErrUnauthorized = errors.New("UNAUTHORIZED")
