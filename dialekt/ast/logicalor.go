package dialekt

// An AST node that represents the logical NOT operator.
type LogicalOr struct {
	*AbstractPolyadicExpression
}

// The child expression being inverted by the NOT operator.
func NewLogicalOr() *LogicalOr {
	return &LogicalOr{}
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (lor *LogicalOr) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitLogicalOr(lor)
}
