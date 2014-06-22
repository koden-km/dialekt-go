package dialekt

// An AST node that is an expression.
// Not all nodes in the tree represent an entire (sub-)expression.
type ExpressionInterface interface {
	NodeInterface
}
