package errors

type BadRequest struct {
	Message    string
	StatusCode int
}

func (e *BadRequest) Error() string {
	return e.Message
}
