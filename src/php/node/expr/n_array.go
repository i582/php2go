package expr

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// Array node
type Array struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Items        []node.Node
}

// NewArray node constructor
func NewArray(Items []node.Node) *Array {
	return &Array{
		FreeFloating: nil,
		Items:        Items,
	}
}

// SetPosition sets node position
func (n *Array) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Array) GetPosition() *position.Position {
	return n.Position
}

func (n *Array) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Array) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Array) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Items != nil {
		v.EnterChildList("Items", n)
		for _, nn := range n.Items {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Items", n)
	}

	v.LeaveNode(n)
}
