package stmt

import (
	"github.com/i582/php2go/src/php/freefloating"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/position"
	"github.com/i582/php2go/src/php/walker"
)

// Constant node
type Constant struct {
	FreeFloating  freefloating.Collection
	Position      *position.Position
	PhpDocComment string
	ConstantName  node.Node
	Expr          node.Node
}

// NewConstant node constructor
func NewConstant(ConstantName node.Node, Expr node.Node, PhpDocComment string) *Constant {
	return &Constant{
		FreeFloating:  nil,
		PhpDocComment: PhpDocComment,
		ConstantName:  ConstantName,
		Expr:          Expr,
	}
}

// SetPosition sets node position
func (n *Constant) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Constant) GetPosition() *position.Position {
	return n.Position
}

func (n *Constant) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Attributes returns node attributes as map
func (n *Constant) Attributes() map[string]interface{} {
	return map[string]interface{}{
		"PhpDocComment": n.PhpDocComment,
	}
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Constant) Walk(v walker.Visitor) {
	if v.EnterNode(n) == false {
		return
	}

	if n.ConstantName != nil {
		v.EnterChildNode("ConstantName", n)
		n.ConstantName.Walk(v)
		v.LeaveChildNode("ConstantName", n)
	}

	if n.Expr != nil {
		v.EnterChildNode("Expr", n)
		n.Expr.Walk(v)
		v.LeaveChildNode("Expr", n)
	}

	v.LeaveNode(n)
}
