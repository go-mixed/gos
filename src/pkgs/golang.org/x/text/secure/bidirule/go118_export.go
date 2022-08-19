// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package bidirule

import (
	q "golang.org/x/text/secure/bidirule"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bidirule",
		Path: "golang.org/x/text/secure/bidirule",
		Deps: map[string]string{
			"errors":                         "errors",
			"golang.org/x/text/transform":    "transform",
			"golang.org/x/text/unicode/bidi": "bidi",
			"unicode/utf8":                   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Transformer": reflect.TypeOf((*q.Transformer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrInvalid": reflect.ValueOf(&q.ErrInvalid),
		},
		Funcs: map[string]reflect.Value{
			"Direction":       reflect.ValueOf(q.Direction),
			"DirectionString": reflect.ValueOf(q.DirectionString),
			"New":             reflect.ValueOf(q.New),
			"Valid":           reflect.ValueOf(q.Valid),
			"ValidString":     reflect.ValueOf(q.ValidString),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
