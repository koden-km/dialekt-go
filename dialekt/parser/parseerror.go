package dialekt

type ParseError struct {
	msg string
}

func NewParseError(msg string) *ParseError {
	if !len(msg) {
		msg = "A parse error has occured."
	}

	return &ParseError{msg}
}

// Not needed because its "inherited"?
// func (e *ParseError) Error() string {
// 	return e.msg
// }
