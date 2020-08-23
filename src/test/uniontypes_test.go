package test

import (
	"testing"

	"github.com/i582/php2go/src/testsuite"
)

func TestUnionTypes(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5; // int
	$a = "Hello"; // string
	// so $a has type int|string

	$b = "string";
	$b = 14;
	$b = 12.56;
	// $b has type int|string|float
}
`))

	s.AddExpected([]byte(`
package test

func Foo() {
	a := NewVar()
	a.Setint64(int64(5))
	a.Setstring("Hello")
	b := NewVar()
	b.Setstring("string")
	b.Setint64(int64(14))
	b.Setfloat64(12.56)
}
`))

	s.RunTest()
}
