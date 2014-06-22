package dialekt

// Interface for node visitors.
type VisitorInterface interface {
	// Visit a LogicalAnd node.
	// The result will be returned.
    VisitLogicalAnd(node LogicalAnd) (result interface{})

    // Visit a LogicalOr node.
	// The result will be returned.
    VisitLogicalOr(node LogicalOr) (result interface{})

    // Visit a LogicalNot node.
	// The result will be returned.
    VisitLogicalNot(node LogicalNot) (result interface{})

    // Visit a Tag node.
	// The result will be returned.
    VisitTag(node Tag) (result interface{})

    // Visit a pattern node.
	// The result will be returned.
    VisitPattern(node Pattern) (result interface{})

    // Visit a PatternLiteral node.
	// The result will be returned.
    VisitPatternLiteral(node PatternLiteral) (result interface{})

    // Visit a PatternWildcard node.
	// The result will be returned.
    VisitPatternWildcard(node PatternWildcard) (result interface{})

    // Visit a EmptyExpression node.
	// The result will be returned.
    VisitEmptyExpression(node EmptyExpression) (result interface{})
}
