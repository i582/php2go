package utils

import (
	"strings"

	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/node/name"
)

func NamePartsToString(parts []node.Node) string {
	s := make([]string, 0, len(parts))
	for _, v := range parts {
		s = append(s, v.(*name.NamePart).Value)
	}
	return strings.Join(s, `\`)
}

func FirstLetterUpperCase(s string) string {
	return strings.ToUpper(s[0:1]) + s[1:]
}

func TransformType(t string) string {
	isMap := strings.HasPrefix(t, "map")

	t = strings.ReplaceAll(t, "[]", "ElementType")

	if isMap {
		t = strings.Replace(t, "[", "WithKey", 1)
		t = strings.Replace(t, "]", "WithValue", 1)
	}

	return t
}

func WithTypeCast(t string, need bool, w func(string), f func()) {
	if need {
		w(t)
		w("(")
	}

	f()

	if need {
		w(")")
	}
}
