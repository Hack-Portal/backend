package hperror

import "errors"

var (
	ErrFieldRequired = errors.New("field required")
	ErrFieldInvalid  = errors.New("field invalid")
	ErrNotFound      = errors.New("not found")
)
