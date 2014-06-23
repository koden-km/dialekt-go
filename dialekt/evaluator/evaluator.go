package dialekt

// Evaluate an expression against a set of tags.
type Evaluator struct {
	// True if tag matching should be case-sensitive; otherwise, false.
	caseSensitive bool

	// True if an empty expression matches all tags; or false to match none.
	emptyIsWildcard bool

	// The set of tags to evaluate against.
	tags []string

	// The expression results.
	expressionResults []ExpressionResult
}

// Case sensitive is true if tag matching should be case-sensitive; otherwise, false.
// Empty is wildcard is true if an empty expression matches all tags; or false to match none.
func NewEvaluator(caseSensitive, emptyIsWildcard bool) *Evaluator {
	return &Evaluator{caseSensitive, emptyIsWildcard, nil, nil}
}

// Evaluate an expression against a set of tags.
// Return the result of the evaluation.
func (ev *Evaluator) Evaluate(expression ExpressionInterface, tags []string) (result *EvaluationResult) {
	ev.tags = tags
	ev.expressionResults = []ExpressionResult(nil)

	result = NewEvaluationResult(expression.Accept(ev).IsMatch(), ev.expressionResults)

	ev.tags = nil
	ev.expressionResults = nil

	return result
}

// Visit a LogicalAnd node.
// The result will be returned.
func (ev *Evaluator) VisitLogicalAnd(node LogicalAnd) (result interface{}) {
	// TODO
}

// Visit a LogicalOr node.
// The result will be returned.
func (ev *Evaluator) VisitLogicalOr(node LogicalOr) (result interface{}) {
	// TODO
}

// Visit a LogicalNot node.
// The result will be returned.
func (ev *Evaluator) VisitLogicalNot(node LogicalNot) (result interface{}) {
	// TODO
}

// Visit a Tag node.
// The result will be returned.
func (ev *Evaluator) VisitTag(node Tag) (result interface{}) {
	// TODO
}

// Visit a pattern node.
// The result will be returned.
func (ev *Evaluator) VisitPattern(node Pattern) (result interface{}) {
	// TODO
}

// Visit a PatternLiteral node.
// The result will be returned.
func (ev *Evaluator) VisitPatternLiteral(node PatternLiteral) (result interface{}) {
	// TODO
}

// Visit a PatternWildcard node.
// The result will be returned.
func (ev *Evaluator) VisitPatternWildcard(node PatternWildcard) (result interface{}) {
	// TODO
}

// Visit a EmptyExpression node.
// The result will be returned.
func (ev *Evaluator) VisitEmptyExpression(node EmptyExpression) (result interface{}) {
	// TODO
}

func (ev *Evaluator) matchTags(expression ExpressionInterface, predicate func(string) bool) []ExpressionResult {
	// TODO
}
