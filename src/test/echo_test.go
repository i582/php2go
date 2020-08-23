package test

import (
	"testing"

	"github.com/i582/php2go/src/testsuite"
)

func TestEchoSimpleType(t *testing.T) {
	s := testsuite.NewSuite(t)
	s.AddFile([]byte(`<?php
function Foo() {
	$a = 5;
	echo $a;
	echo 12.45;
	echo "Hello";
	echo true;
	
	echo $a, 1, 5.56, "Hello";
}
`))

	s.AddExpected([]byte(`
package test

import (
	"fmt"
)

func Foo() {
	a := int64(5)
	fmt.Print(a)
	fmt.Print(12.45)
	fmt.Print("Hello")
	fmt.Print(true)
	fmt.Print(a, int64(1), 5.56, "Hello")
}
`))

	s.RunTest()
}
