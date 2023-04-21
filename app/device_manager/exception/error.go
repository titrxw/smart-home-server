package exception

import errorhandler "github.com/we7coreteam/w7-rangine-go/src/core/err_handler"

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

func NewResponseError(msg string) errorhandler.ResponseError {
	return errorhandler.ResponseError{
		Msg: msg,
	}
}
