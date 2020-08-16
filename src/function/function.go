package function

import (
	"fmt"

	"github.com/i582/php2go/src/types"
	"github.com/i582/php2go/src/variable"
)

type Param struct {
	Name string
	Type types.Types
}

func (v Param) String() string {
	return fmt.Sprintf("%s: %v", v.Name, v.Type)
}

type Function struct {
	Name       string
	ReturnType types.Types
	Params     []Param
	Variables  variable.Table
}

func NewFunction(name string, returnType types.Types, params []Param) *Function {
	return &Function{Name: name, ReturnType: returnType, Params: params}
}

func (v Function) String() string {
	var params string

	for _, param := range v.Params {
		params += param.String()
	}

	return fmt.Sprintf("%s(%s): %v", v.Name, params, v.ReturnType)
}
