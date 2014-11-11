package dialekt

// An AST node that is an expression.
//
// Not all nodes in the tree represent an entire (sub-)expression.
type ExpressionInterface interface {
	NodeInterface

	// Fetch the first token from the source that is part of this expression.
	FirstToken() *Token

	// Fetch the last token from the source that is part of this expression.
	LastToken() *Token

	// Set the delimiting tokens for this expression.
	SetTokens(firstToken, lastToken *Token)
}
