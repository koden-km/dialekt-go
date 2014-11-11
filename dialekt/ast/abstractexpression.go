package dialekt

type AbstractExpression struct {
	ExpressionInterface

	firstToken *Token
	lastToken  *Token
}

// Fetch the first token from the source that is part of this expression.
func (exp *AbstractExpression) FirstToken() *Token {
	return exp.firstToken
}

// Fetch the last token from the source that is part of this expression.
func (exp *AbstractExpression) LastToken() *Token {
	return exp.lastToken
}

// Set the delimiting tokens for this expression.
func (exp *AbstractExpression) SetTokens(firstToken, lastToken) {
	exp.firstToken = firstToken
	exp.lastToken = lastToken
}
