package dialekt

// An AST node that represents the logical AND operator.
type LogicalAnd struct {
	*AbstractPolyadicExpression
}

// One or more children to add to this operator.
func NewLogicalAnd(children ...ExpressionInterface) *LogicalAnd {
	return &LogicalAnd{NewAbstractPolyadicOperator(children...)}
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (land *LogicalAnd) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitLogicalAnd(land)
}
