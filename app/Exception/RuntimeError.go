package exception

type RuntimeError struct {
	ErrorAbstract
}

func NewRuntimeError(msg string, args ...interface{}) *RuntimeError {
	code := 500
	var previous error

	for _, arg := range args {
		switch value := arg.(type) {
		case int:
			code = value
		case error:
			previous = value
		default:
			panic("Unknown argument")
		}
	}
	return &RuntimeError{
		ErrorAbstract{
			Msg:           msg,
			ErrorCode:     code,
			ErrorPrevious: previous,
		},
	}
}
