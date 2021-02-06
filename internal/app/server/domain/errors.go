package domain

import "errors"

var (
	// ErrInternalServerError will throw if any the Internal Server Error happen
	ErrInternalServerError = errors.New("Internal Server Error")
	// ErrNotFound will throw if the requested item is not exists
	ErrNotFound        = errors.New("No item with the specified ID found")
	ErrWrongType       = errors.New("WRONGTYPE Operation against a key holding the wrong kind of value")
	ErrIndexOutOfRange = errors.New("ERR index out of range")
	ErrNoSuchKey       = errors.New("ERR no such key")
)
