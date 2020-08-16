package block

import (
	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/node/expr"
	"github.com/i582/php2go/src/php/node/expr/assign"
	"github.com/i582/php2go/src/php/node/stmt"
	"github.com/i582/php2go/src/php/walker"
	"github.com/i582/php2go/src/types"
	"github.com/i582/php2go/src/variable"

	"github.com/i582/php2go/src/ctx"
	"github.com/i582/php2go/src/solver"
)

type BlockWalker struct {
	Ctx ctx.Context
}

func (b BlockWalker) EnterChildNode(key string, w walker.Walkable) {}
func (b BlockWalker) LeaveChildNode(key string, w walker.Walkable) {}
func (b BlockWalker) EnterChildList(key string, w walker.Walkable) {}
func (b BlockWalker) LeaveChildList(key string, w walker.Walkable) {}
func (b *BlockWalker) LeaveNode(w walker.Walkable)                 {}

func (b *BlockWalker) EnterNode(w walker.Walkable) bool {
	n := w.(node.Node)

	switch n := n.(type) {
	case *node.Root:

	case *stmt.Expression:

	case *expr.ShortArray:
		return b.handleArray(n)
	case *assign.Assign:
		return b.handleAssign(n)
	case *stmt.For:
		return b.handleFor(n)
	case *stmt.While:
		return b.handleWhile(n)
	case *stmt.If:
		return b.handleIf(n)
	case *stmt.Return:
		return b.handleReturn(n)
	case *expr.Variable:
		return b.handleVariable(n)
	}

	return true
}

func (b *BlockWalker) Context() ctx.Context {
	return b.Ctx
}

func (b *BlockWalker) handleArray(a *expr.ShortArray) bool {
	for _, item := range a.Items {
		item.Walk(b)
	}

	return false
}

func (b *BlockWalker) handleFor(f *stmt.For) bool {
	w := &BlockWalker{
		Ctx: ctx.Context{
			Parent:          &b.Ctx,
			Variables:       variable.NewTable(),
			CurrentFunction: nil,
		},
	}
	for _, init := range f.Init {
		init.Walk(w)
	}

	for _, cond := range f.Cond {
		cond.Walk(w)
	}

	for _, aftereffect := range f.Loop {
		aftereffect.Walk(w)
	}

	f.Stmt.Walk(w)

	f.Ctx = w.Ctx

	return false
}

func (b *BlockWalker) handleWhile(wl *stmt.While) bool {
	w := &BlockWalker{
		Ctx: ctx.Context{
			Parent:          &b.Ctx,
			Variables:       variable.NewTable(),
			CurrentFunction: nil,
		},
	}

	wl.Cond.Walk(w)
	wl.Stmt.Walk(w)

	wl.Ctx = w.Ctx

	return false
}

func (b *BlockWalker) handleIf(i *stmt.If) bool {
	w := &BlockWalker{
		Ctx: ctx.Context{
			Parent:          &b.Ctx,
			Variables:       variable.NewTable(),
			CurrentFunction: b.Ctx.CurrentFunction,
		},
	}
	i.Cond.Walk(b)
	i.Stmt.Walk(w)

	ww := &BlockWalker{
		Ctx: ctx.Context{
			Parent:          &b.Ctx,
			Variables:       variable.NewTable(),
			CurrentFunction: b.Ctx.CurrentFunction,
		},
	}

	if i.Else != nil {
		i.Else.Walk(ww)
	}

	for _, v := range w.Ctx.Variables.Vars {
		if vv, ok := ww.Ctx.Variables.Get(v.Name); ok {
			v.Type.Merge(vv.Type)
			vv.Type.Merge(v.Type)
			b.Ctx.Variables.Add(v.Name, v.Type)
			newVar, _ := b.Ctx.Variables.Get(v.Name)
			newVar.FromIfElse = true
			v.WasInitialize = true
			vv.WasInitialize = true
		}
	}

	i.IfCtx = w.Ctx
	i.ElseCtx = ww.Ctx

	return false
}

func (b *BlockWalker) handleReturn(ret *stmt.Return) bool {
	tp := solver.ExprTypeLocal(&b.Ctx, ret.Expr)

	if ret.Expr == nil {
		tp = types.NewBaseTypes(types.Void)
	}
	if b.Ctx.CurrentFunction != nil {
		b.Ctx.CurrentFunction.ReturnType.Merge(tp)
	}

	return true
}

func (b *BlockWalker) handleAssign(a *assign.Assign) bool {
	e := a.Expression
	e.Walk(b)
	switch a := a.Variable.(type) {
	case *expr.Variable:
		var varName string
		switch name := a.VarName.(type) {
		case *node.Identifier:
			varName = name.Value
		}
		var ok bool
		a.Var, ok = b.Ctx.GetVariable(varName)
		if ok {
			a.Var.AddType(solver.ExprTypeLocal(&b.Ctx, e))
		} else {
			b.Ctx.Variables.Add(varName, solver.ExprTypeLocal(&b.Ctx, e))
			a.Var, _ = b.Ctx.Variables.Get(varName)
		}
	}
	a.Variable.Walk(b)
	return false
}

func (b *BlockWalker) handleVariable(v *expr.Variable) bool {
	var varName string
	switch name := v.VarName.(type) {
	case *node.Identifier:
		varName = name.Value
	}

	var ok bool
	v.Var, ok = b.Ctx.GetVariable(varName)
	if !ok {
		panic("var not found")
	}

	return true
}
