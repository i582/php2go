package solver

import (
	"strings"

	"github.com/i582/php2go/src/php/node"
	"github.com/i582/php2go/src/php/node/expr"
	"github.com/i582/php2go/src/php/node/expr/binary"
	"github.com/i582/php2go/src/php/node/name"
	"github.com/i582/php2go/src/php/node/scalar"
	"github.com/i582/php2go/src/utils"

	"github.com/i582/php2go/src/ctx"
	"github.com/i582/php2go/src/types"
)

func binaryOpType(ctx *ctx.Context, left node.Node, right node.Node) types.Types {
	lt := ExprTypeLocal(ctx, left)
	rt := ExprTypeLocal(ctx, right)

	switch {
	case lt.Is(types.Integer) && rt.Is(types.Integer):
		return types.NewBaseTypes(types.Integer)
	case lt.Is(types.Integer) && rt.Is(types.Float) || lt.Is(types.Float) && rt.Is(types.Integer):
		return types.NewBaseTypes(types.Float)
	case lt.Is(types.Float) && rt.Is(types.Float):
		return types.NewBaseTypes(types.Float)
	case lt.Is(types.String) && rt.Is(types.String):
		return types.NewBaseTypes(types.String)
	default:
		panic("error operand types")
	}
}

func ExprType(ctx *ctx.Context, n node.Node) types.Types {
	tp := ExprTypeLocal(ctx, n)

	if tp.Resolved() {
		return tp
	}

	return ResolveTypes(ctx, tp)
}

func ExprTypeLocal(ctx *ctx.Context, n node.Node) types.Types {
	switch n := n.(type) {

	case *expr.FunctionCall:
		fnName := utils.NamePartsToString(n.Function.(*name.Name).Parts)
		return types.NewTypes(types.NewLazyFunctionCallType(fnName))

	case *expr.Variable:
		nm := n.VarName.(*node.Identifier).Value
		v, ok := ctx.GetVariable(nm)
		if !ok {
			panic("variable not found")
		}
		return v.Type

	case *expr.ShortArray:
		return arrayType(ctx, n)
	case *expr.ArrayItem:
		return ExprTypeLocal(ctx, n.Val)

	case *scalar.Lnumber:
		return types.NewBaseTypes(types.Integer)
	case *scalar.Dnumber:
		return types.NewBaseTypes(types.Float)
	case *scalar.String:
		return types.NewBaseTypes(types.String)
	case *name.Name:
		nm := utils.NamePartsToString(n.Parts)
		if nm == "true" || nm == "false" {
			return types.NewBaseTypes(types.Bool)
		} else if strings.EqualFold(nm, "null") {
			return types.NewBaseTypes(types.Null)
		}

	case *expr.ConstFetch:
		return ExprTypeLocal(ctx, n.Constant)

	case *binary.Plus:
		return binaryOpType(ctx, n.Left, n.Right)
	case *binary.Minus:
		return binaryOpType(ctx, n.Left, n.Right)
	case *binary.Mul:
		return binaryOpType(ctx, n.Left, n.Right)
	case *binary.Div:
		return binaryOpType(ctx, n.Left, n.Right)
	case *binary.Concat:
		return types.NewBaseTypes(types.String)

	case *binary.Equal:
		return types.NewBaseTypes(types.Bool)
	case *binary.NotEqual:
		return types.NewBaseTypes(types.Bool)
	case *binary.Smaller:
		return types.NewBaseTypes(types.Bool)
	case *binary.SmallerOrEqual:
		return types.NewBaseTypes(types.Bool)
	case *binary.Greater:
		return types.NewBaseTypes(types.Bool)
	case *binary.GreaterOrEqual:
		return types.NewBaseTypes(types.Bool)

	case *binary.LogicalAnd:
		return types.NewBaseTypes(types.Bool)
	case *binary.LogicalOr:
		return types.NewBaseTypes(types.Bool)
	case *binary.LogicalXor:
		return types.NewBaseTypes(types.Bool)

	case *binary.BooleanAnd:
		return types.NewBaseTypes(types.Bool)
	case *binary.BooleanOr:
		return types.NewBaseTypes(types.Bool)
	}

	return types.Types{}
}

func arrayType(ctx *ctx.Context, a *expr.ShortArray) types.Types {
	if len(a.Items) == 0 {
		return types.NewTypes(types.NewArrayType(types.Integer))
	}

	isAssoc := a.Items[0].(*expr.ArrayItem).Key != nil

	keyType := ExprTypeLocal(ctx, a.Items[0].(*expr.ArrayItem).Key)
	valType := ExprTypeLocal(ctx, a.Items[0].(*expr.ArrayItem).Val)

	var t types.Types

	if isAssoc {
		t = types.NewTypes(types.NewAssociativeArrayType(keyType, valType, 1))
	} else {
		t = types.NewTypes(types.NewPlainArrayType(valType, 1))
	}

	return t
}
