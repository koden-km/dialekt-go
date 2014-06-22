package dialekt

// Parse an expression.
type ParserInterface interface {
	// The expression to parse.
	// Returns the parsed expression or an error if the expression is invalid.
	Parse(expression string) (expression ExpressionInterface, error ParseError)
}
