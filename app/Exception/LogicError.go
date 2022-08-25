package exception

type LogicError struct {
	RuntimeError
}

func NewLogicError(msg string, args ...interface{}) *LogicError {
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

	return &LogicError{
		RuntimeError{
			ErrorAbstract{
				Msg:           msg,
				ErrorCode:     code,
				ErrorPrevious: previous,
			},
		},
	}
}
