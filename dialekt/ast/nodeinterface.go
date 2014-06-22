package dialekt

// An AST node.
type NodeInterface interface {
	// Pass this node to the appropriate method on the given visitor.
	// The visitation result will be returned.
	Accept(visitor VisitorInterface) (result interface{})
}
