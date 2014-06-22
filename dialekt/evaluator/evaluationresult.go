package dialekt

// The overall result of the evaluation of an expression.
type EvaluationResult struct {
	// True if the expression matched the tag set; otherwise, false.
	isMatch bool

	// The individual sub-expression results. A map of expression to result.
	expressionResults [ExpressionInterface]ExpressionResult
}

// Is match is true if the expression matched the tag set; otherwise, false.
// Expression results are the individual sub-expression results.
func NewExpressionResult(isMatch bool, expressionResults []ExpressionResult) *EvaluationResult {
	return &EvaluationResult{isMatch, expressionResults}
}

// Indicates whether or not the expression matched the tag set.
// Returns true if the expression matched the tag set; otherwise, false.
func (er *EvaluationResult) IsMatch() bool {
	return er.isMatch
}

// Fetch the result for an individual expression node from the AST.
// The expression for which the result is fetched.
// Returns the result for the given expression or an error if there is no result for the given expression.
func (er *EvaluationResult) ResultOf(expression ExpressionInterface) (result ExpressionResult, error UnexpectedValueError) {
	return er.expressionResults[expression]
}
