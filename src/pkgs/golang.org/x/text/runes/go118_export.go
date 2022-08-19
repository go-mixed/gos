// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package runes

import (
	q "golang.org/x/text/runes"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "runes",
		Path: "golang.org/x/text/runes",
		Deps: map[string]string{
			"golang.org/x/text/transform": "transform",
			"unicode":                     "unicode",
			"unicode/utf8":                "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Set": reflect.TypeOf((*q.Set)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Transformer": reflect.TypeOf((*q.Transformer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"If":               reflect.ValueOf(q.If),
			"In":               reflect.ValueOf(q.In),
			"Map":              reflect.ValueOf(q.Map),
			"NotIn":            reflect.ValueOf(q.NotIn),
			"Predicate":        reflect.ValueOf(q.Predicate),
			"Remove":           reflect.ValueOf(q.Remove),
			"ReplaceIllFormed": reflect.ValueOf(q.ReplaceIllFormed),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
