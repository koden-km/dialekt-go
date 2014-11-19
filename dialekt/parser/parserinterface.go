package dialekt

// Parse an expression.
type ParserInterface interface {
	// The expression to parse.
	// Returns the parsed expression or an error if the expression is invalid.
	Parse(expression string, lexer LexerInterface) (parsedExpression ExpressionInterface, error ParseError)

	// Parse an expression that has already beed tokenized.
	// Returns the parsed expression or an error if the expression is invalid.
	ParseTokens(tokens []Token) (parsedExpression ExpressionInterface, error ParseError)
}
