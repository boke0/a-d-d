package exception

import (
	"errors"
)

var (
	Unauthorized = errors.New("unauthorized")
	InvalidToken = errors.New("invalid token")
	SomeReason = errors.New("some reason")
)
