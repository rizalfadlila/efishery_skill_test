package exceptions

import "errors"

// ErrMissingAuthHeader error missing `Authorization` header on context
var ErrMissingAuthHeader = errors.New("missing auth on context")

// ErrMissingName error missing name header on context
var ErrMissingName = errors.New("missing name on JWT token")

// ErrMissingEmail error missing email header on context
var ErrMissingEmail = errors.New("missing email on JWT token")

// ErrMissingPhone error missing phone header on context
var ErrMissingPhone = errors.New("missing phone on JWT token")

// ErrMissingRole error missing role header on context
var ErrMissingRole = errors.New("missing role on JWT token")

// ErrMissingTimestamp error missing timestamp header on context
var ErrMissingTimestamp = errors.New("missing timestamp on JWT token")
