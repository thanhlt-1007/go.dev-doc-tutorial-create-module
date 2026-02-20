package errs

type emptyNameError struct{}

func NewEmptyNameError() error {
	return &emptyNameError{}
}

func (err *emptyNameError) Error() string {
	return "name is empty"
}
