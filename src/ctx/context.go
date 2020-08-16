package ctx

import (
	"github.com/i582/php2go/src/function"
	"github.com/i582/php2go/src/variable"
)

type Context struct {
	Parent          *Context
	Variables       variable.Table
	CurrentFunction *function.Function

	InAssign            bool
	InPrintFunctionCall bool
	InCondition         bool
	InCompare           bool
	InBoolean           bool
}

func (c Context) GetVariable(name string) (*variable.Variable, bool) {
	v, ok := c.Variables.Get(name)
	if !ok {
		if c.Parent == nil {
			return nil, false
		}

		v, ok := c.Parent.GetVariable(name)
		if !ok {
			return nil, false
		}

		return v, true
	}

	return v, true
}
