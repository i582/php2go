package test

import (
	"testing"

	"github.com/i582/php2go/src/testsuite"
)

func TestSimpleTypesIsT(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5;
	if (is_int($a)) {
		echo "integer";
	}
	$b = 5.56;
	if (is_float($b)) {
		echo "float";
	}
	$c = true;
	if (is_bool($c)) {
		echo "bool";
	}
	$d = "string";
	if (is_string($d)) {
		echo "string";
	}
}
`))

	s.AddExpected([]byte(`
package test

import (
	"fmt"
)

func Foo() {
	a := int64(5)
	if Isint64Simple(a) {
		fmt.Print("integer")
	}
	b := 5.56
	if Isfloat64Simple(b) {
		fmt.Print("float")
	}
	c := true
	if IsboolSimple(c) {
		fmt.Print("bool")
	}
	d := "string"
	if IsstringSimple(d) {
		fmt.Print("string")
	}
}
`))

	s.RunTest()
}

func TestUnionTypesIsT(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5;
	$a = "string";
	if (is_int($a)) {
		echo "integer";
	}
	$b = 5.56;
	$b = 5;
	if (is_float($b)) {
		echo "float";
	}
	$c = true;
	$c = 5;
	if (is_bool($c)) {
		echo "bool";
	}
	$d = "string";
	$d = true;
	if (is_string($d)) {
		echo "string";
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
	a.Setstring("string")
	if Isint64Simple(a.Getstring()) {
		fmt.Print("integer")
	}
	b := NewVar()
	b.Setfloat64(5.56)
	b.Setint64(int64(5))
	if Isfloat64Simple(b.Getint64()) {
		fmt.Print("float")
	}
	c := NewVar()
	c.Setbool(true)
	c.Setint64(int64(5))
	if IsboolSimple(c.Getint64()) {
		fmt.Print("bool")
	}
	d := NewVar()
	d.Setstring("string")
	d.Setbool(true)
	if IsstringSimple(d.Getbool()) {
		fmt.Print("string")
	}
}
`))

	s.RunTest()
}

func TestUnionTypesIsTWithBranching(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5;
	if ($a > 100) {
		$a = "string";
	}
	if (is_int($a)) {
		echo "integer";
	}
	$b = 5.56;
	if ($b != 10) {
		$b = "string";
	}
	if (is_float($b)) {
		echo "float";
	}
	$c = true;
	if ($b != 10) {
		$c = false;
	}
	if (is_bool($c)) {
		echo "bool";
	}
	$d = "string";
	if ($d != "") {
		$d = 12;
	}
	if (is_string($d)) {
		echo "string";
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
	if a.Getint64() > int64(100) {
		a.Setstring("string")
	}
	if Isint64Simple(a.Getstring()) {
		fmt.Print("integer")
	}
	b := NewVar()
	b.Setfloat64(5.56)
	if b.Getfloat64() != int64(10) {
		b.Setstring("string")
	}
	if Isfloat64Simple(b.Getstring()) {
		fmt.Print("float")
	}
	c := true
	if b.Getstring() != int64(10) {
		c = false
	}
	if IsboolSimple(c) {
		fmt.Print("bool")
	}
	d := NewVar()
	d.Setstring("string")
	if d.Getstring() != "" {
		d.Setint64(int64(12))
	}
	if IsstringSimple(d.Getint64()) {
		fmt.Print("string")
	}
}
`))

	s.RunTest()
}
