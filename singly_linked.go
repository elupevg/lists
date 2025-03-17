package lists

// NodeSL represents a node in a singly-linked list.
type NodeSL[T any] struct {
	Value T
	Next  *NodeSL[T]
}

// NewSinglyLinked creates a new singly-linked list from its slice representation.
func NewSinglyLinked[T any](s []T) *NodeSL[T] {
	var head *NodeSL[T]
	for _, val := range s {
		head = head.Append(val)
	}
	return head
}

// Append adds a node at the tail of the singly-linked list.
func (n *NodeSL[T]) Append(val T) *NodeSL[T] {
	if n == nil {
		return &NodeSL[T]{Value: val}
	}
	n.Next = n.Next.Append(val)
	return n
}

// Slice returns a slice representation of the singly-linked list.
func (n *NodeSL[T]) Slice() []T {
	if n == nil {
		return []T{}
	}
	return append([]T{n.Value}, n.Next.Slice()...)
}
