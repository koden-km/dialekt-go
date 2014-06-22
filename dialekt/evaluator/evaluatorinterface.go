package dialekt

// Interface for expression evaluators.
// An expression evaluator checks whether a set of tags match against a certain expression.
type EvaluatorInterface interface {
	// Evaluate an expression against a set of tags.
	// The expression to evaluate, and the set of tags to evaluate against.
	// Returns the result of the evaluation.
	Evaluate(expression ExpressionInterface, tags []string) EvaluationResult
}
