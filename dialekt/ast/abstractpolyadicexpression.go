package dialekt

// A base class providing common functionality for polyadic operators.
type AbstractPolyadicExpression struct {
	AbstractExpression

	children []ExpressionInterface
}

// Zero or more children to add to this operator.
func NewAbstractPolyadicOperator(children ...ExpressionInterface) *AbstractPolyadicExpression {
	return &AbstractPolyadicExpression{children}
}

// Add a child expression to this operator.
func (po *AbstractPolyadicExpression) Add(expression ExpressionInterface) {
	append(po.children, expression)
}

// Fetch a slice of this operator's children.
func (po *AbstractPolyadicExpression) Children() []ExpressionInterface {
	return po.children
}
