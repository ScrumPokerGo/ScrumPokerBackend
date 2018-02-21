package errors

// New returns an error that formats as the given text.
func New(text string,status int) error {
	return &ErrorCore{status,text}
}

// errorString is a trivial implementation of error.
type ErrorCore struct {
	status int
	desc string
}

func (e *ErrorCore) Error() string {
	return e.desc
}