package dialekt

// An AST node that represents the logical NOT operator.
type LogicalOr struct {
	*AbstractPolyadicExpression
}

// One or more children to add to this operator.
func NewLogicalOr(children ...ExpressionInterface) *LogicalOr {
	return &LogicalOr{NewAbstractPolyadicOperator(children...)}
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (lor *LogicalOr) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitLogicalOr(lor)
}
