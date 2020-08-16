package solver

import (
	"github.com/i582/php2go/src/ctx"
	"github.com/i582/php2go/src/meta"
	"github.com/i582/php2go/src/types"
)

func ResolveTypes(ctx *ctx.Context, tp types.Types) types.Types {
	res := types.Types{}
	for _, t := range tp.Types {
		res.Merge(ResolveType(ctx, t))
	}
	return res
}

func ResolveType(ctx *ctx.Context, t types.Type) types.Types {
	switch t.LazyType {
	case types.FunctionCall:
		fnInfo, ok := meta.GetFunction(t.FunctionName)
		if !ok {
			return types.Types{}
		}
		return fnInfo.ReturnType
	}

	return types.NewTypes(t)
}
