package binary

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// Coalesce node
type Coalesce struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Left         node.Node
	Right        node.Node
}

// NewCoalesce node constructor
func NewCoalesce(Variable node.Node, Expression node.Node) *Coalesce {
	return &Coalesce{
		FreeFloating: nil,
		Left:         Variable,
		Right:        Expression,
	}
}

// SetPosition sets node position
func (n *Coalesce) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Coalesce) GetPosition() *position.Position {
	return n.Position
}

func (n *Coalesce) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Coalesce) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Coalesce) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Left != nil {
		v.EnterChildNode("Left", n)
		n.Left.Walk(v)
		v.LeaveChildNode("Left", n)
	}

	if n.Right != nil {
		v.EnterChildNode("Right", n)
		n.Right.Walk(v)
		v.LeaveChildNode("Right", n)
	}

	v.LeaveNode(n)
}
