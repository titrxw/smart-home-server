package exception

type ArgsError struct {
	ErrorAbstract
}

func NewArgsError(msg string, args ...interface{}) *ArgsError {
	code := 403
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

	return &ArgsError{
		ErrorAbstract{
			Msg:           msg,
			ErrorCode:     code,
			ErrorPrevious: previous,
		},
	}
}
