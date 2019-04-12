package splathon

// Error is error object of Splathon API.
type Error struct {
	// HTTP status code.
	Code int
	// Error message.
	Message string
}

func (e *Error) Error() string {
	return e.Message
}
