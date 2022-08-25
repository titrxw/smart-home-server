package exception

type ErrorAbstract struct {
	error

	Msg           string
	ErrorCode     int
	ErrorPrevious error
}

func (errorAbstract *ErrorAbstract) Error() string {
	return errorAbstract.Msg
}

func (errorAbstract *ErrorAbstract) Code() int {
	return errorAbstract.ErrorCode
}

func (errorAbstract *ErrorAbstract) Previous() error {
	return errorAbstract.ErrorPrevious
}
