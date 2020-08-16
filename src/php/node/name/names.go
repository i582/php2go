package name

import (
	"github.com/i582/php2go/src/php/node"
)

// Names is generalizing the Name types
type Names interface {
	node.Node
	GetParts() []node.Node
}
