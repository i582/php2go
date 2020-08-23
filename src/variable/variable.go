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

func (v *Variable) AddType(ts types.Types, inBranching bool) {
	v.Type.Merge(ts)
	if !inBranching {
		v.CurrentType = ts
	} else {
		v.CurrentType = types.Types{}
	}
}

func (v *Variable) GenerateDefinition() string {
	name := v.Type.GenerateName()
	return fmt.Sprintf("%s := New%s()\n", v.Name, name)
}

func (v *Variable) GenerateAccess(inAssignLvalue, inAssignRvalue, inPrint, inCompare, inBoolean, inIsT bool) string {
	var field string
	varHasUnionType := !v.Type.SingleType()
	currentTypeIsSingle := v.CurrentType.SingleType()

	if varHasUnionType && !currentTypeIsSingle {
		if inPrint {
			return fmt.Sprintf("%s.String()", v.Name)
		}

		if inBoolean {
			return fmt.Sprintf("%s.Bool()", v.Name)
		}
	}

	if varHasUnionType && currentTypeIsSingle {
		if inAssignLvalue {
			field = ".Set" + utils.TransformType(v.CurrentType.String()) + "("
		} else {
			field = ".Get" + utils.TransformType(v.CurrentType.String()) + "()"
		}
	}

	return fmt.Sprintf("%s%s", v.Name, field)
}
