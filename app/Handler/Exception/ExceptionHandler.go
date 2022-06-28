package exception

import (
	exception "github.com/titrxw/go-framework/src/Core/Exception"
)

type ExceptionHandler struct {
	exception.ExceptionHandler
}

func (this *ExceptionHandler) Handle(err error, trace string) {
	this.ExceptionHandler.Handle(err, trace)
}
