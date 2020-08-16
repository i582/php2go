package root

import (
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/node/expr"
	"github.com/i582/php2go/src/php/node/expr/assign"
	"github.com/i582/php2go/src/php/node/stmt"
	"github.com/i582/php2go/src/php/walker"

	"github.com/i582/php2go/src/block"
	"github.com/i582/php2go/src/ctx"
	"github.com/i582/php2go/src/function"
	"github.com/i582/php2go/src/meta"
	"github.com/i582/php2go/src/variable"
)

type RootWalker struct {
	Ctx ctx.Context
}

func (r RootWalker) EnterChildNode(key string, w walker.Walkable) {}
func (r RootWalker) LeaveChildNode(key string, w walker.Walkable) {}
func (r RootWalker) EnterChildList(key string, w walker.Walkable) {}
func (r RootWalker) LeaveChildList(key string, w walker.Walkable) {}
func (r *RootWalker) LeaveNode(w walker.Walkable)                 {}

func (r *RootWalker) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	switch n := n.(type) {
	case *node.Root:

	case *stmt.Expression:

	case *assign.Assign:
		r.handleAssign(n)

	case *stmt.Function:
		r.handleFunction(n)
	}

	return true
}

func (r *RootWalker) handleAssign(a *assign.Assign) {

}

func (r *RootWalker) handleFunction(f *stmt.Function) {
	fn := function.Function{}

	fn.Name = f.FunctionName.(*node.Identifier).Value

	for _, param := range f.Params {
		fn.Params = append(fn.Params, r.handleFunctionParam(param.(*node.Parameter)))
	}

	r.handleFunctionStmts(f.Stmts, &fn)

	meta.AddFunction(&fn)

	f.Func = &fn
}

func (r *RootWalker) handleFunctionParam(p *node.Parameter) function.Param {
	name := p.Variable.(*expr.Variable).VarName.(*node.Identifier).Value

	return function.Param{
		Name: name,
	}
}

func (r *RootWalker) handleFunctionStmts(stmts []node.Node, fn *function.Function) {
	w := &block.BlockWalker{
		Ctx: ctx.Context{
			Parent:          &r.Ctx,
			Variables:       variable.NewTable(),
			CurrentFunction: fn,
		},
	}

	for _, st := range stmts {
		st.Walk(w)
	}

	fn.Variables = w.Context().Variables
}
