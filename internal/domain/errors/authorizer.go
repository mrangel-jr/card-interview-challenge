package errors

import "errors"

var (
	ErrInvalidTimestamp = errors.New("timestamp not valid")
	ErrInTheFuture      = errors.New("timestamp on future")
	ErrInvalidPayload   = errors.New("invalid payload")
)
