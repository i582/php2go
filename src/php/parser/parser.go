package parser

import (
	"github.com/i582/php2go/src/php/errors"
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/php5"
	"github.com/i582/php2go/src/php/php7"
	"github.com/i582/php2go/src/php/version"
)

// Parser interface
type Parser interface {
	Parse() int
	GetRootNode() node.Node
	GetErrors() []*errors.Error
	WithFreeFloating()
}

func NewParser(src []byte, v string) (Parser, error) {
	var parser Parser

	r, err := version.Compare(v, "7.0")
	if err != nil {
		return nil, err
	}

	if r == -1 {
		parser = php5.NewParser(src, v)
	} else {
		parser = php7.NewParser(src, v)
	}

	return parser, nil
}
