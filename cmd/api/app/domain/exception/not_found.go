package exception

type NotFound interface {
	Error() string
	NotFound() bool
}

type DataNotFound struct {
	ErrMessage string
}

func (notFound DataNotFound) Error() string {
	return notFound.ErrMessage
}

func (notFound DataNotFound) NotFound() bool {
	return true
}
