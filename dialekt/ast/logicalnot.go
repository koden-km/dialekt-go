package dialekt

// An AST node that represents the logical NOT operator.
type LogicalNot struct {
	// The expression being inverted by the NOT operator.
	child ExpressionInterface
}

// The child expression being inverted by the NOT operator.
func NewLogicalNot(child ExpressionInterface) *LogicalNot {
	return &LogicalNot{child}
}

// Fetch the expression being inverted by the NOT operator.
func (lnot *LogicalNot) Child() ExpressionInterface {
	return lnot.child
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (lnot *LogicalNot) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitLogicalAnd(lnot)
}
