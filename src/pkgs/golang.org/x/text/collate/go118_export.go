// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package collate

import (
	q "golang.org/x/text/collate"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "collate",
		Path: "golang.org/x/text/collate",
		Deps: map[string]string{
			"bytes":                              "bytes",
			"golang.org/x/text/internal/colltab": "colltab",
			"golang.org/x/text/language":         "language",
			"golang.org/x/text/unicode/norm":     "norm",
			"sort":                               "sort",
			"strings":                            "strings",
		},
		Interfaces: map[string]reflect.Type{
			"Lister": reflect.TypeOf((*q.Lister)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Buffer":   reflect.TypeOf((*q.Buffer)(nil)).Elem(),
			"Collator": reflect.TypeOf((*q.Collator)(nil)).Elem(),
			"Option":   reflect.TypeOf((*q.Option)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Force":            reflect.ValueOf(&q.Force),
			"IgnoreCase":       reflect.ValueOf(&q.IgnoreCase),
			"IgnoreDiacritics": reflect.ValueOf(&q.IgnoreDiacritics),
			"IgnoreWidth":      reflect.ValueOf(&q.IgnoreWidth),
			"Loose":            reflect.ValueOf(&q.Loose),
			"Numeric":          reflect.ValueOf(&q.Numeric),
		},
		Funcs: map[string]reflect.Value{
			"New":            reflect.ValueOf(q.New),
			"NewFromTable":   reflect.ValueOf(q.NewFromTable),
			"OptionsFromTag": reflect.ValueOf(q.OptionsFromTag),
			"Reorder":        reflect.ValueOf(q.Reorder),
			"Supported":      reflect.ValueOf(q.Supported),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"CLDRVersion":    {"untyped string", constant.MakeString(string(q.CLDRVersion))},
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
