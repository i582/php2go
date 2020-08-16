package stmt

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// Echo node
type Echo struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Exprs        []node.Node
}

// NewEcho node constructor
func NewEcho(Exprs []node.Node) *Echo {
	return &Echo{
		FreeFloating: nil,
		Exprs:        Exprs,
	}
}

// SetPosition sets node position
func (n *Echo) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Echo) GetPosition() *position.Position {
	return n.Position
}

func (n *Echo) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Echo) Attributes() map[string]interface{} {
	return nil
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Echo) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.Exprs != nil {
		v.EnterChildList("Exprs", n)
		for _, nn := range n.Exprs {
			if nn != nil {
				nn.Walk(v)
			}
		}
		v.LeaveChildList("Exprs", n)
	}

	v.LeaveNode(n)
}
