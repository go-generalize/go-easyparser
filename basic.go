// package easyparser provides a Go module parser for TypeScript AST
package easyparser

import (
	"go/types"
	"reflect"

	tstypes "github.com/go-generalize/go-easyparser/types"
)

func (p *pkgParser) parseBasic(t *types.Basic) tstypes.Type {
	switch {
	case t.Info()&(types.IsInteger|types.IsFloat) != 0:
		return &tstypes.Number{
			RawType: t.Kind(),
		}
	case t.Info()&types.IsBoolean != 0:
		return &tstypes.Boolean{}
	case t.Info()&types.IsString != 0:
		return &tstypes.String{}
	default:
		panic("unsupported type: " + reflect.TypeOf(t).String())
	}
}
