package dialekt

// The result for an invidiual expression in the AST.
type ExpressionResult struct {
	// The expression to which this result applies.
	expression ExpressionInterface

	// True if the expression matched the tag set; otherwise, false.
	isMatch bool

	// The set of tags that matched.
	matchedTags []string

	// The set of tags that did not match.
	unmatchedTags []string
}

// Expression is the expression to which this result applies.
// Is match is true if the expression matched the tag set; otherwise, false.
// Matched tags is a set of the tags that matched.
// Unmatched tags is a set of the tags that did not match.
func NewExpressionResult(expression ExpressionInterface, isMatch bool, matchedTags []string, unmatchedTags []string) *ExpressionResult {
	return &ExpressionResult{expression, isMatch, matchedTags, unmatchedTags}
}

// Fetch the expression to which this result applies.
func (exr *ExpressionResult) Expression() ExpressionInterface {
}

// Indicates whether or not the expression matched the tag set.
// Returns true if the expression matched the tag set; otherwise, false.
func (exr *ExpressionResult) IsMatch() bool {
	return exr.isMatch
}

// Fetch the set of tags that matched.
func (exr *ExpressionResult) MatchedTags() []string {
	return exr.matchedTags
}

// Fetch set of tags that did not match.
func (exr *ExpressionResult) UnmatchedTags() []string {
	return exr.unmatchedTags
}
