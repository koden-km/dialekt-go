package dialekt

// An AST node that represents an empty expression.
type EmptyExpression struct {
	*AbstractExpression
}

func NewEmptyExpression() *EmptyExpression {
	return &EmptyExpression{}
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (exp *EmptyExpression) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitEmptyExpression(exp)
}
