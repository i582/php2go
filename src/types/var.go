package types

import (
	"fmt"

	"github.com/i582/php2go/src/utils"
)

type VarInfo struct {
	Fields map[string]struct{}
}

func NewVarInfo() VarInfo {
	return VarInfo{Fields: make(map[string]struct{})}
}

func (v *VarInfo) Add(f string) {
	v.Fields[f] = struct{}{}
}

func (v *VarInfo) AddTypes(types Types) {
	for _, t := range types.Types {
		v.Fields[t.String()] = struct{}{}
	}
}

func (v *VarInfo) Generate() string {
	var res string

	res += "\n"
	res += "type ValueType uint8\n"

	var constants string
	for f := range v.Fields {
		constants += "\tConstant" + utils.TransformType(f) + " ValueType = iota\n"
	}
	res += "\n"

	res += "const (\n"
	res += constants
	res += ")\n"

	res += `
type Var struct {
	Val  interface{}
	Type ValueType
}

func NewVar() Var {
	return Var{}
}

`

	getterTemplate := `func (v *Var) Get%s() %s {
	return v.Val.(%s)
}

`

	setterTemplate := `func (v *Var) Set%s(val %s)  {
	v.Val = val
	v.Type = Constant%s
}

`

	caseTemplate := `	case Constant%s:
		%s
`

	res += `func (v *Var) Bool() bool {
	switch v.Type {
`

	for f := range v.Fields {
		var code string
		switch f {
		case "int64":
			code = "return v.Val.(int64) != 0"
		case "float64":
			code = "return v.Val.(float64) != 0"
		case "string":
			code = "return v.Val.(string) != \"\""
		case "bool":
			code = "return v.Val.(bool)"
		default:
			code = "return false"
		}

		res += fmt.Sprintf(caseTemplate, utils.TransformType(f), code)
	}

	res += `	}

	return false
}

`

	res += `func (v *Var) String() string {
	switch v.Type {
`

	for f := range v.Fields {
		var code string
		switch f {
		case "int64":
			code = "return fmt.Sprint(v.Val.(int64))"
		case "float64":
			code = "return fmt.Sprint(v.Val.(float64))"
		case "string":
			code = "return v.Val.(string)"
		case "bool":
			code = "return fmt.Sprint(v.Val.(bool))"
		default:
			code = "return \"\""
		}

		res += fmt.Sprintf(caseTemplate, utils.TransformType(f), code)
	}

	res += `	}

	return ""
}

`

	res += "\n"
	res += "type CompareType uint8\n"

	var compareConst string

	compareConst += "\tEqual CompareType = iota\n"
	compareConst += "\tNotEqual\n"
	compareConst += "\tGreater\n"
	compareConst += "\tGreaterEqual\n"
	compareConst += "\tSmaller\n"
	compareConst += "\tSmallerEqual\n"

	res += "\n"

	res += "const (\n"
	res += compareConst
	res += ")\n\n"

	compareTemplate := `func (v *Var) CompareWith%s(val %s, compare CompareType) bool {
	switch v.Type {
`

	compareSwitchTemplate := `switch compare {
		case Equal:
			return v.Val.(%[1]s) == val
		case NotEqual:
			return v.Val.(%[1]s) != val
		case Greater:
			return v.Val.(%[1]s) > val
		case GreaterEqual:
			return v.Val.(%[1]s) >= val
		case Smaller:
			return v.Val.(%[1]s) < val
		case SmallerEqual:
			return v.Val.(%[1]s) <= val
		}`

	for fieldFor := range v.Fields {

		res += fmt.Sprintf(compareTemplate, utils.TransformType(fieldFor), fieldFor)

		for f := range v.Fields {
			var code string
			switch f {
			case "int64":
				if fieldFor == "int64" {
					code = fmt.Sprintf(compareSwitchTemplate, "int64")
				} else {
					code = "return false"
				}
			case "float64":
				if fieldFor == "float64" {
					code = fmt.Sprintf(compareSwitchTemplate, "float64")
				} else {
					code = "return false"
				}
			case "string":
				if fieldFor == "string" {
					code = fmt.Sprintf(compareSwitchTemplate, "string")
				} else {
					code = "return false"
				}
			case "bool":
				if fieldFor == "bool" {
					code = `switch compare {
		case Equal:
			return v.Val.(bool) == val
		case NotEqual:
			return v.Val.(bool) != val
		case Greater:
			return false
		case GreaterEqual:
			return false
		case Smaller:
			return false
		case SmallerEqual:
			return false
		}`
				} else {
					code = "return false"
				}

			default:
				code = "return false"
			}

			res += fmt.Sprintf(caseTemplate, utils.TransformType(f), code)
		}

		res += `	}

	return false
}

`
	}

	for f := range v.Fields {
		res += fmt.Sprintf(getterTemplate, utils.TransformType(f), f, f)
	}

	for f := range v.Fields {
		res += fmt.Sprintf(setterTemplate, utils.TransformType(f), f, utils.TransformType(f))
	}

	return res
}
