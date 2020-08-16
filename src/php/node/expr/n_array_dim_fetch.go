package expr

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// ArrayDimFetch node
type ArrayDimFetch struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Variable     node.Node
	Dim          node.Node
}

// NewArrayDimFetch node constructor
func NewArrayDimFetch(Variable node.Node, Dim node.Node) *ArrayDimFetch {
	return &ArrayDimFetch{
		FreeFloating: nil,
		Variable:     Variable,
		Dim:          Dim,
	}
}

// SetPosition sets node position
func (n *ArrayDimFetch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ArrayDimFetch) GetPosition() *position.Position {
	return n.Position
}

func (n *ArrayDimFetch) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *ArrayDimFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ArrayDimFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Variable != nil {
		v.EnterChildNode("Variable", n)
		n.Variable.Walk(v)
		v.LeaveChildNode("Variable", n)
	}

	if n.Dim != nil {
		v.EnterChildNode("Dim", n)
		n.Dim.Walk(v)
		v.LeaveChildNode("Dim", n)
	}

	v.LeaveNode(n)
}
