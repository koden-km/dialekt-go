package dialekt

// Represents a literal (exact-match) portion of a pattern expression.
type PatternLiteral struct {
	PatternChildInterface

	// The string to match.
	string string
}

// Zero or more pattern literals or placeholders.
func NewPatternLiteral(string string) *PatternLiteral {
	return &PatternLiteral{string}
}

// Fetch the string to be matched.
func (lit *PatternLiteral) String() []PatternChildInterface {
	return lit.string
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (lit *PatternLiteral) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitPattern(lit)
}
