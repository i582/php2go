package types

import (
	"fmt"
)

const (
	Undefined Base = iota
	Void
	Integer
	Float
	Bool
	String

	Arr

	Lazy
)

type Base uint8

const (
	None LazyType = iota
	FunctionCall
)

type LazyType uint8

type LazyTypeFields struct {
	FunctionName string
}

type Type struct {
	BaseType Base

	LazyType LazyType
	LazyTypeFields

	Array
}

func NewType(baseType Base) Type {
	return Type{BaseType: baseType}
}

func NewArrayType(baseType Base) Type {
	return NewPlainArrayType(NewTypes(NewType(baseType)), 1)
}

func NewLazyFunctionCallType(fn string) Type {
	return Type{
		BaseType: Lazy,
		LazyType: FunctionCall,

		LazyTypeFields: LazyTypeFields{
			FunctionName: fn,
		},
	}
}

func (t Type) String() string {
	var str string

	switch t.BaseType {
	case Undefined:
		str += "undefined"
	case Integer:
		str += "int64"
	case Float:
		str += "float64"
	case String:
		str += "string"
	case Bool:
		str += "bool"
	case Void:
		str += "void"

	case Arr:
		arr := t.Array
		keys := arr.KeysTypes.String()
		elems := arr.ElemTypes.String()

		if t.IsAssociative {
			str += fmt.Sprintf("map[%s]%s", keys, elems)
		} else {
			str += fmt.Sprintf("[]%s", elems)
		}

		break

	case Lazy:
		str += "lazy"

		switch t.LazyType {
		case FunctionCall:
			str += "<FunctionCall: " + t.FunctionName + ">"
		}
	}

	return str
}

func (t Type) Is(tp Base) bool {
	return t.BaseType == tp
}

func (t Type) IsLazy() bool {
	return t.BaseType == Lazy
}

func (t Type) ElementTypes() (Types, bool) {
	if t.BaseType != Arr {
		return Types{}, false
	}

	return t.ElemTypes, true
}

type Array struct {
	KeysTypes Types
	ElemTypes Types
	ArrayDim  uint8

	IsAssociative bool
}

func NewAssociativeArrayType(keyTypes Types, elemTypes Types, dim uint8) Type {
	return Type{
		BaseType: Arr,

		Array: Array{
			KeysTypes:     keyTypes,
			ElemTypes:     elemTypes,
			ArrayDim:      dim,
			IsAssociative: true,
		},
	}
}

func NewPlainArrayType(elemTypes Types, dim uint8) Type {
	return Type{
		BaseType: Arr,

		Array: Array{
			KeysTypes: Types{Types: []Type{{
				BaseType: Integer,
			}}},
			ElemTypes:     elemTypes,
			ArrayDim:      dim,
			IsAssociative: false,
		},
	}
}
