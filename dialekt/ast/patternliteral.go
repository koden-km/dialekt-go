package dialekt

// Represents a literal (exact-match) portion of a pattern expression.
type PatternLiteral struct {
	PatternChildInterface

	// The string to match.
	matchString string
}

// The string to match.
func NewPatternLiteral(matchString string) *PatternLiteral {
	return &PatternLiteral{matchString}
}

// Fetch the string to be matched.
func (plit *PatternLiteral) String() []PatternChildInterface {
	return plit.matchString
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (plit *PatternLiteral) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitPatternLiteral(plit)
}
