package exception

type InternalServerErrorPort interface {
	Error() string
	InternalServerErrorPort() bool
}

type InternalServerError struct {
	ErrMessage string
}

func (internalServerError InternalServerError) Error() string {
	return internalServerError.ErrMessage
}
