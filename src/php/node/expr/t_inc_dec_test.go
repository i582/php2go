package expr_test

import (
	"testing"

	"gotest.tools/assert"

	"github.com/i582/php2go/src/php/node/expr"
	"github.com/i582/php2go/src/php/position"

	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/node/stmt"
	"github.com/i582/php2go/src/php/php5"
	"github.com/i582/php2go/src/php/php7"
)

func TestPostDec(t *testing.T) {
	src := `<? $a--;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    8,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    8,
				},
				Expr: &expr.PostDec{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    5,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
							Value: "a",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPostInc(t *testing.T) {
	src := `<? $a++;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    8,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    8,
				},
				Expr: &expr.PostInc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  3,
							EndPos:    5,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  3,
								EndPos:    5,
							},
							Value: "a",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPreDec(t *testing.T) {
	src := `<? --$a;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    8,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    8,
				},
				Expr: &expr.PreDec{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  5,
							EndPos:    7,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    7,
							},
							Value: "a",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestPreInc(t *testing.T) {
	src := `<? ++$a;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  3,
			EndPos:    8,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  3,
					EndPos:    8,
				},
				Expr: &expr.PreInc{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  3,
						EndPos:    7,
					},
					Variable: &expr.Variable{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  5,
							EndPos:    7,
						},
						VarName: &node.Identifier{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  5,
								EndPos:    7,
							},
							Value: "a",
						},
					},
				},
			},
		},
	}

	php7parser := php7.NewParser([]byte(src), "7.4")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)

	php5parser := php5.NewParser([]byte(src), "5.6")
	php5parser.Parse()
	actual = php5parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
