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

func TestEchoUnionType(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5; // int
	$a = "Hello"; // string
	
	echo $a;

	$b = "string";

	if ($a == "") {
		$b = 12;
	}

}
`))

	s.AddExpected([]byte(`
package test

import (
	"fmt"
)

func Foo() {
	a := NewVar()
	a.Setint64(int64(5))
	a.Setstring("Hello")
	fmt.Print(a.Getstring())
	b := NewVar()
	b.Setstring("string")
	if a.Getstring() == "" {
		b.Setint64(int64(12))
	}
}
`))

	s.RunTest()
}

func TestGetCurrentType(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5; // int
	$a = "Hello"; // string
	
	$b = $a;
	echo $b;
}
`))

	s.AddExpected([]byte(`
package test

import (
	"fmt"
)

func Foo() {
	a := NewVar()
	a.Setint64(int64(5))
	a.Setstring("Hello")
	b := a.Getstring()
	fmt.Print(b)
}
`))

	s.RunTest()
}
