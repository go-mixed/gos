// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package search

import (
	q "golang.org/x/text/search"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "search",
		Path: "golang.org/x/text/search",
		Deps: map[string]string{
			"golang.org/x/text/internal/colltab": "colltab",
			"golang.org/x/text/language":         "language",
			"strings":                            "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"IndexOption": reflect.TypeOf((*q.IndexOption)(nil)).Elem(),
			"Matcher":     reflect.TypeOf((*q.Matcher)(nil)).Elem(),
			"Option":      reflect.TypeOf((*q.Option)(nil)).Elem(),
			"Pattern":     reflect.TypeOf((*q.Pattern)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Exact":            reflect.ValueOf(&q.Exact),
			"IgnoreCase":       reflect.ValueOf(&q.IgnoreCase),
			"IgnoreDiacritics": reflect.ValueOf(&q.IgnoreDiacritics),
			"IgnoreWidth":      reflect.ValueOf(&q.IgnoreWidth),
			"Loose":            reflect.ValueOf(&q.Loose),
			"Supported":        reflect.ValueOf(&q.Supported),
			"WholeWord":        reflect.ValueOf(&q.WholeWord),
		},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Anchor":    {reflect.TypeOf(q.Anchor), constant.MakeInt64(int64(q.Anchor))},
			"Backwards": {reflect.TypeOf(q.Backwards), constant.MakeInt64(int64(q.Backwards))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"CLDRVersion":    {"untyped string", constant.MakeString(string(q.CLDRVersion))},
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
