package expr

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// StaticPropertyFetch node
type StaticPropertyFetch struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Class        node.Node
	Property     node.Node
}

// NewStaticPropertyFetch node constructor
func NewStaticPropertyFetch(Class node.Node, Property node.Node) *StaticPropertyFetch {
	return &StaticPropertyFetch{
		FreeFloating: nil,
		Class:        Class,
		Property:     Property,
	}
}

// SetPosition sets node position
func (n *StaticPropertyFetch) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *StaticPropertyFetch) GetPosition() *position.Position {
	return n.Position
}

func (n *StaticPropertyFetch) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *StaticPropertyFetch) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *StaticPropertyFetch) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Class != nil {
		v.EnterChildNode("Class", n)
		n.Class.Walk(v)
		v.LeaveChildNode("Class", n)
	}

	if n.Property != nil {
		v.EnterChildNode("Property", n)
		n.Property.Walk(v)
		v.LeaveChildNode("Property", n)
	}

	v.LeaveNode(n)
}
