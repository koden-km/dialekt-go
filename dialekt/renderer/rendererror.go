package dialekt

type RenderError struct {
	errorString
}

func NewRenderError(msg string) *RenderError {
	if !len(msg) {
		msg = "A render error has occured."
	}

	return &RenderError{msg}
}

// Not needed because its "inherited"?
// func (e *RenderError) Error() string {
// 	return e.msg
// }
