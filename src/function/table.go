package function

type Table struct {
	Functions map[string]*Function
}

func NewTable() Table {
	return Table{Functions: make(map[string]*Function)}
}

func (t *Table) Add(f *Function) bool {
	if t.Contains(f.Name) {
		return false
	}

	t.Functions[f.Name] = f
	return true
}

func (t Table) Contains(name string) bool {
	_, ok := t.Functions[name]
	return ok
}

func (t Table) Get(name string) (*Function, bool) {
	if !t.Contains(name) {
		return nil, false
	}

	fn, ok := t.Functions[name]
	return fn, ok
}
