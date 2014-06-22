package dialekt

// An AST node that represents the logical AND operator.
type LogicalAnd struct {
	AbstractPolyadicOperator
}

func NewLogicalAnd() *LogicalAnd {
	return &LogicalAnd{}
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (land *LogicalAnd) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitLogicalAnd(land)
}
