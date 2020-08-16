package meta

import (
	"github.com/i582/php2go/src/function"
	"github.com/i582/php2go/src/variable"
)

var (
	AllVariables = variable.NewTable()
	AllFunctions = function.NewTable()
)

func AddVariable(v variable.Variable) {
	AllVariables.Add(v.Name, v.Type)
}

func AddFunction(f *function.Function) {
	AllFunctions.Add(f)
}

func GetFunction(name string) (*function.Function, bool) {
	return AllFunctions.Get(name)
}
