package dialekt

// Render an expression.
type RenderInterface interface {
	// Render the given expression to a string.
	// Return the rendered expression.
	Render(expression ExpressionInterface) string
}
