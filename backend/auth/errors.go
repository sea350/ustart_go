package auth

import "errors"

var (
	errDuplicateUname = errors.New("User cannot be registered, because the provided username was a duplicate")
)
