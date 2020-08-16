package testsuite

import (
	"bytes"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"

	"github.com/i582/php2go/src/php/php7"
	"github.com/i582/php2go/src/root"

	"github.com/i582/php2go/src/generator"
)

type Test struct {
	t        *testing.T
	Content  []byte
	Expected []byte
}

func NewTest(t *testing.T, content, expected string) Test {
	return Test{
		t:        t,
		Content:  []byte(content),
		Expected: []byte(strings.TrimPrefix(expected, "\n")),
	}
}

func (t *Test) RunTest() {
	parser := php7.NewParser(t.Content, "7.4")
	parser.Parse()

	for _, e := range parser.GetErrors() {
		t.t.Error(e)
	}

	rw := root.RootWalker{}

	rootNode := parser.GetRootNode()
	rootNode.Walk(&rw)

	var b []byte
	buf := bytes.NewBuffer(b)
	gw := generator.NewGeneratorWalker(buf, "test.php")
	rootNode.Walk(&gw)
	gw.Final()

	if !cmp.Equal(buf.String(), string(t.Expected)) {
		t.t.Error(cmp.Diff(buf.String(), string(t.Expected)))
	}
}
