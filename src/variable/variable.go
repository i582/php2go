package variable

import (
	"fmt"

	"github.com/i582/php2go/src/types"
	"github.com/i582/php2go/src/utils"
)

type Variable struct {
	Name string
	Type types.Types

	WasInitialize bool
	CurrentType   types.Types
	FromIfElse    bool
}

func NewVariable(name string, typ types.Types) *Variable {
	return &Variable{Name: name, Type: typ}
}

func (v Variable) String() string {
	return fmt.Sprintf("$%s: %v", v.Name, v.Type)
}

func (v *Variable) AddType(ts types.Types) {
	v.Type.Merge(ts)
}

func (v *Variable) GenerateDefinition() string {
	name := v.Type.GenerateName()
	return fmt.Sprintf("%s := New%s()\n", v.Name, name)
}

func (v *Variable) GenerateAccess(set, inPrint, inCompare, inBoolean bool) string {
	var field string
	if v.CurrentType.SingleType() && !v.Type.SingleType() {
		if set {
			field = ".Set" + utils.TransformType(v.CurrentType.String()) + "("
		} else {
			field = ".Get" + utils.TransformType(v.CurrentType.String()) + "()"
		}
	} else if inBoolean && !inCompare && !v.Type.SingleType() {
		field = ".Bool()"
	}

	if (v.CurrentType.Types == nil || !v.CurrentType.SingleType()) && !v.Type.SingleType() && inPrint {
		field = ".String()"
	}

	return fmt.Sprintf("%s%s", v.Name, field)
}
