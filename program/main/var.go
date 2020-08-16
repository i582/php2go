package main

type Foo struct {
}

type Var struct {
	Val interface{}
}

func (v *Var) GetInt64() int64 {
	return v.Val.(int64)
}
