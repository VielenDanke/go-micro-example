package pb

func (e *ErrorNotFound) Error() string {
	return e.Msg
}

func (e *Error) Error() string {
	return e.Msg
}
