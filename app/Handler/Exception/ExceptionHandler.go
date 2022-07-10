package exception

import (
	exception "github.com/titrxw/go-framework/src/Core/Exception"
)

type ExceptionHandler struct {
	exception.ExceptionHandler
}

func (handler *ExceptionHandler) Handle(err error, trace string) {
	handler.ExceptionHandler.Handle(err, trace)
}
