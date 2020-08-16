package expr

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// ClosureUse node
type ClosureUse struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Uses         []node.Node
}

// NewClosureUse node constructor
func NewClosureUse(Uses []node.Node) *ClosureUse {
	return &ClosureUse{
		FreeFloating: nil,
		Uses:         Uses,
	}
}

// SetPosition sets node position
func (n *ClosureUse) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *ClosureUse) GetPosition() *position.Position {
	return n.Position
}

func (n *ClosureUse) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *ClosureUse) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *ClosureUse) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Uses != nil {
		v.EnterChildList("Uses", n)
		for _, nn := range n.Uses {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Uses", n)
	}

	v.LeaveNode(n)
}
