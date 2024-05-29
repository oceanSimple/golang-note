package myError

type Error struct {
	error string
}

func (e *Error) Error() string {
	return e.error
}

func New(text string) error {
	return &Error{text}
}
