package dialekt

// Tokenize an expression.
type LexerInterface interface {
	// The expression to tokenize.
	// Returns the tokens of the expression or an error if the expression is invalid.
	Lex(expression string) (tokens []Token, error ParseError)
}
