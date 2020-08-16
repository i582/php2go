package variable

import (
	"github.com/i582/php2go/src/types"
)

type Table struct {
	Vars map[string]*Variable
}

func NewTable() Table {
	return Table{Vars: make(map[string]*Variable)}
}

func (t *Table) Add(name string, typ types.Types) bool {
	if t.Contains(name) {
		return false
	}

	t.Vars[name] = NewVariable(name, typ)
	return true
}

func (t *Table) AddManually(v *Variable) bool {
	if t.Contains(v.Name) {
		return false
	}

	t.Vars[v.Name] = v
	return true
}

func (t *Table) Join(t2 Table) bool {
	for _, v := range t2.Vars {
		t.AddManually(v)
	}
	return true
}

func (t Table) Get(name string) (*Variable, bool) {
	v, ok := t.Vars[name]

	return v, ok
}

func (t Table) Contains(name string) bool {
	_, ok := t.Vars[name]
	return ok
}
