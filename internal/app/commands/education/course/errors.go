package course

type BadRequestError struct {
	message string
}

func NewBadRequestError(msg string) *BadRequestError {
	return &BadRequestError{
		message: msg,
	}
}

func (b *BadRequestError) Error() string {
	return b.message
}
