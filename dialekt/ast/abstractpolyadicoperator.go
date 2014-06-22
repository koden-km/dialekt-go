package dialekt

// A base class providing common functionality for polyadic operators.
type AbstractPolyadicOperator struct {
	ExpressionInterface
	children []ExpressionInterface
}

// Zero or more children to add to this operator.
func NewAbstractPolyadicOperator(children ...ExpressionInterface) *AbstractPolyadicOperator {
	return &AbstractPolyadicOperator{children}
}

// Add a child expression to this operator.
func (po *AbstractPolyadicOperator) Add(expression ExpressionInterface) {
	append(po.children, expression)
}

// Fetch a slice of this operator's children.
func (po *AbstractPolyadicOperator) Children() []ExpressionInterface {
	return po.children
}
