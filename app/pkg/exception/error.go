package exception

import (
	"errors"
)

type RuntimeError struct {
	Previous error
	Msg      string
}

func NewRuntimeError(msg string) RuntimeError {
	return RuntimeError{
		Msg: msg,
	}
}

func (e RuntimeError) Error() string {
	return e.Msg
}

func (e RuntimeError) Unwrap() error {
	return e.Previous
}

func NewResponseError(msg string) error {
	return errors.New(msg)
}
