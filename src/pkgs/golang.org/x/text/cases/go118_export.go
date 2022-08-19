// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package cases

import (
	q "golang.org/x/text/cases"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cases",
		Path: "golang.org/x/text/cases",
		Deps: map[string]string{
			"golang.org/x/text/internal":     "internal",
			"golang.org/x/text/language":     "language",
			"golang.org/x/text/transform":    "transform",
			"golang.org/x/text/unicode/norm": "norm",
			"strings":                        "strings",
			"unicode":                        "unicode",
			"unicode/utf8":                   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Caser":  reflect.TypeOf((*q.Caser)(nil)).Elem(),
			"Option": reflect.TypeOf((*q.Option)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Compact":   reflect.ValueOf(&q.Compact),
			"NoLower":   reflect.ValueOf(&q.NoLower),
			"Supported": reflect.ValueOf(&q.Supported),
		},
		Funcs: map[string]reflect.Value{
			"Fold":             reflect.ValueOf(q.Fold),
			"HandleFinalSigma": reflect.ValueOf(q.HandleFinalSigma),
			"Lower":            reflect.ValueOf(q.Lower),
			"Title":            reflect.ValueOf(q.Title),
			"Upper":            reflect.ValueOf(q.Upper),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
