package dialekt

// An AST node that represents a literal tag expression.
type Tag struct {
	AbstractExpression
	ExpressionInterface

	// The tag name.
	name string
}

// The tag name.
func NewTag(name string) *Tag {
	return &Tag{name}
}

// Fetch the tag name.
func (tag *Tag) Name() string {
	return tag.name
}

// Pass this node to the appropriate method on the given visitor.
// The visitation result will be returned.
func (tag *Tag) Accept(visitor VisitorInterface) (result interface{}) {
	return visitor.VisitTag(tag)
}
