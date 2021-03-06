package exceptions

import "errors"

// ErrMissingAuthHeader error missing `Authorization` header on context
var ErrMissingAuthHeader = errors.New("missing auth on context")
