package types

import (
	"fmt"
	"sort"
	"strings"
)

type Types struct {
	Types []Type
}

func NewTypes(tp ...Type) Types {
	types := Types{}
	for _, tp := range tp {
		types.Add(tp)
	}
	return types
}

func NewBaseTypes(tp ...Base) Types {
	types := Types{}
	for _, tp := range tp {
		types.Add(NewType(tp))
	}
	return types
}

func (ts *Types) Resolved() bool {
	for _, t := range ts.Types {
		if t.IsLazy() {
			return false
		}
	}

	return true
}

func (ts *Types) ElementType() Types {
	var res Types

	for _, t := range ts.Types {
		res.Merge(t.ElemTypes)
	}

	return res
}

func (ts *Types) KeyType() Types {
	var res Types

	for _, t := range ts.Types {
		res.Merge(t.KeysTypes)
	}

	return res
}

func (ts *Types) SingleType() bool {
	return ts.Len() == 1
}

func (ts *Types) Add(t Type) {
	if ts.Contains(t) {
		return
	}

	ts.Types = append(ts.Types, t)
}

func (ts *Types) AddArrayType(keyTypes Types, elemTypes Types, dim uint8) {
	ts.Add(NewAssociativeArrayType(keyTypes, elemTypes, dim))
}

func (ts *Types) Contains(t Type) bool {
	for _, tp := range ts.Types {
		if tp.Is(t.BaseType) {
			return true
		}
	}

	return false
}

func (ts *Types) ContainsMap(ts2 Types) bool {
	for _, tp := range ts2.Types {
		if !ts.Contains(tp) {
			return false
		}
	}

	return true
}

func (ts *Types) Equal(ts2 Types) bool {
	if ts.Len() != ts2.Len() {
		return false
	}

	for _, t := range ts.Types {
		if !ts2.Contains(t) {
			return false
		}
	}

	return true
}

func (ts *Types) Merge(ts2 Types) {
	for _, t := range ts2.Types {
		ts.Add(t)
	}
}

func (ts Types) Is(t Base) bool {
	if ts.Len() == 0 {
		return false
	}

	if ts.Len() == 1 {
		return ts.Types[0].Is(t)
	}

	return false
}

func (ts *Types) Len() int {
	return len(ts.Types)
}

func (ts Types) String() string {
	var typesString []string

	for _, tp := range ts.Types {
		typesString = append(typesString, tp.String())
	}

	sort.Slice(typesString, func(i, j int) bool {
		return typesString[i] < typesString[j]
	})

	if len(ts.Types) == 0 {
		return "empty"
	}

	return strings.Join(typesString, "|")
}

func (ts Types) GenerateName() string {
	if ts.Len() == 0 {
		return ""
	}
	if ts.SingleType() {
		return ts.String()
	}
	return "Var"
}

func (ts Types) GenerateCreation(ts2 Types) (string, bool) {
	if ts.Equal(ts2) {
		return "", false
	}
	if ts.SingleType() {
		return "", false
	}

	return fmt.Sprintf("Var"), true
}
