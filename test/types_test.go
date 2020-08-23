package test

import (
	"testing"

	"github.com/i582/php2go/src/testsuite"
)

func TestSimpleTypes(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5; // int
	$b = 5.56; // float
	$c = true; // bool
	$d = "string"; // string
}
`))

	s.AddExpected([]byte(`
package test

func Foo() {
	a := int64(5)
	b := 5.56
	c := true
	d := "string"
}
`))

	s.RunTest()
}
