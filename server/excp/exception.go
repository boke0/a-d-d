package exception

import (
	"errors"
)

var (
	Unauthorized = errors.New("unauthorized")
	SomeReason = errors.New("some reason")
)
