package dialekt

// An AST node that represents a pattern-match expression.
type Pattern struct {
	ExpressionInterface

	*AbstractExpression

	// Pattern literals or placeholders.
	children []PatternChildInterface
}

// Zero or more pattern literals or placeholders.
func NewPattern(children ...PatternChildInterface) *Pattern {
	return &Pattern{children}
}

// Add a child to this node.
func (pat *Pattern) Add(child PatternChildInterface) {
	pat.children = append(pat.children, child)
}

// Fetch a slice of this node's children.
func (pat *Pattern) Children() []PatternChildInterface {
	return pat.children
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (pat *Pattern) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitPattern(pat)
}
