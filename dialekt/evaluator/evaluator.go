package dialekt

import (
	"regexp"
	"strings"
)

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
	ev.expressionResults = make([]ExpressionResult)

	result = NewEvaluationResult(expression.Accept(ev).IsMatch(), ev.expressionResults)

	ev.tags = nil
	ev.expressionResults = nil

	return result
}

// Visit a LogicalAnd node.
// The result will be returned.
func (ev *Evaluator) VisitLogicalAnd(node *LogicalAnd) interface{} {
	matchedTags := make([]string)
	isMatch := false

	for n := range node.Children() {
		result := n.Accept(ev)

		if !result.IsMatch() {
			isMatch = false
		}

		matchedTags = append(matchedTags, result.MatchedTags()...)
	}

	unmatchedTags := stringsDiff(ev.tags, matchedTags)

	ev.expressionResults = append(ev.expressionResults, NewExpressionResult(node, isMatch, matchedTags, unmatchedTags))

	return ev.expressionResults
}

// Visit a LogicalOr node.
// The result will be returned.
func (ev *Evaluator) VisitLogicalOr(node LogicalOr) (result interface{}) {
	matchedTags := make([]string)
	isMatch := false

	for n := range node.Children() {
		result := n.Accept(ev)

		if !result.IsMatch() {
			isMatch = false
		}

		matchedTags = append(matchedTags, result.MatchedTags()...)
	}

	unmatchedTags := stringsDiff(ev.tags, matchedTags)

	newExpResult := NewExpressionResult(node, isMatch, matchedTags, unmatchedTags)

	ev.expressionResults = append(ev.expressionResults, newExpResult)

	return ev.expressionResults
}

// Visit a LogicalNot node.
// The result will be returned.
func (ev *Evaluator) VisitLogicalNot(node LogicalNot) (result interface{}) {
	childResult := node.Child().Accept(ev)

	newExpResult := NewExpressionResult(node, !childResult.IsMatch(), childResult.UnmatchedTags(), childResult.MatchedTags())

	ev.expressionResults = append(ev.expressionResults, newExpResult)

	return ev.expressionResults
}

// Visit a Tag node.
// The result will be returned.
func (ev *Evaluator) VisitTag(node Tag) (result interface{}) {
	if ev.caseSensitive {
		predicate := func(tag string) bool {
			return node.name() == tag
		}

		return ev.matchTags(node, predicate)
	} else {
		predicate := func(tag string) bool {
			return strings.EqualFold(node.name(), tag)
		}

		return ev.matchTags(node, predicate)
	}
}

// Visit a pattern node.
// The result will be returned.
func (ev *Evaluator) VisitPattern(node Pattern) (result interface{}) {
	pattern := "/^"

	for n := range node.Children() {
		pattern += n.Accept(ev)
	}

	pattern += "$/"

	if !ev.caseSensitive {
		pattern += "i"
	}

	predicate := func(tag string) bool {
		matched, err := regex.MatchString(pattern, tag)
		return matched
	}

	return ev.matchTags(node, predicate)
}

// Visit a PatternLiteral node.
// The result will be returned.
func (ev *Evaluator) VisitPatternLiteral(node PatternLiteral) (result interface{}) {
	return regexp.QuoteMeta(node.String)
}

// Visit a PatternWildcard node.
// The result will be returned.
func (ev *Evaluator) VisitPatternWildcard(node PatternWildcard) (result interface{}) {
	return ".*"
}

// Visit a EmptyExpression node.
// The result will be returned.
func (ev *Evaluator) VisitEmptyExpression(node EmptyExpression) (result interface{}) {
	otherTags := make([]string)

	if ev.emptyIsWildcard {
		newExpResult := NewExpressionResult(node, ev.emptyIsWildcard, ev.tags, otherTags)
	} else {
		newExpResult := NewExpressionResult(node, ev.emptyIsWildcard, otherTags, ev.tags)
	}

	ev.expressionResults = append(ev.expressionResults, newExpResult)

	return ev.expressionResults
}

func (ev *Evaluator) matchTags(expression ExpressionInterface, predicate func(string) bool) []ExpressionResult {
	matchedTags := make([]string)
	unmatchedTags := make([]string)

	for tag := range ev.tags {
		if predicate(tag) {
			matchedTags = append(matchedTags, tag)
		} else {
			unmatchedTags = append(unmatchedTags, tag)
		}
	}

	newExpResult := NewExpressionResult(expression, len(matchedTags) > 0, matchedTags, unmatchedTags)

	ev.expressionResults = append(ev.expressionResults, newExpResult)

	return ev.expressionResults
}

// Helper function that works like PHP array_diff()
// Remove this and use the golang-set third party lib?
// "github.com/deckarep/golang-set"
func stringsDiff(left, right []string) []string {
	result := make([]string)

	for l := range left {
		addToResult := true

		for r := range right {
			if l == r {
				addToResult = false
				break
			}
		}

		if addToResult {
			result = append(result, l)
		}
	}

	return result
}
