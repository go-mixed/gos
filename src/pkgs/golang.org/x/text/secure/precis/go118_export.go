// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package precis

import (
	q "golang.org/x/text/secure/precis"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "precis",
		Path: "golang.org/x/text/secure/precis",
		Deps: map[string]string{
			"bytes":                             "bytes",
			"errors":                            "errors",
			"golang.org/x/text/cases":           "cases",
			"golang.org/x/text/language":        "language",
			"golang.org/x/text/runes":           "runes",
			"golang.org/x/text/secure/bidirule": "bidirule",
			"golang.org/x/text/transform":       "transform",
			"golang.org/x/text/unicode/norm":    "norm",
			"golang.org/x/text/width":           "width",
			"unicode":                           "unicode",
			"unicode/utf8":                      "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Option":      reflect.TypeOf((*q.Option)(nil)).Elem(),
			"Profile":     reflect.TypeOf((*q.Profile)(nil)).Elem(),
			"Transformer": reflect.TypeOf((*q.Transformer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"BidiRule":              reflect.ValueOf(&q.BidiRule),
			"DisallowEmpty":         reflect.ValueOf(&q.DisallowEmpty),
			"FoldWidth":             reflect.ValueOf(&q.FoldWidth),
			"IgnoreCase":            reflect.ValueOf(&q.IgnoreCase),
			"Nickname":              reflect.ValueOf(&q.Nickname),
			"OpaqueString":          reflect.ValueOf(&q.OpaqueString),
			"UsernameCaseMapped":    reflect.ValueOf(&q.UsernameCaseMapped),
			"UsernameCasePreserved": reflect.ValueOf(&q.UsernameCasePreserved),
		},
		Funcs: map[string]reflect.Value{
			"AdditionalMapping":    reflect.ValueOf(q.AdditionalMapping),
			"Disallow":             reflect.ValueOf(q.Disallow),
			"FoldCase":             reflect.ValueOf(q.FoldCase),
			"LowerCase":            reflect.ValueOf(q.LowerCase),
			"NewFreeform":          reflect.ValueOf(q.NewFreeform),
			"NewIdentifier":        reflect.ValueOf(q.NewIdentifier),
			"NewRestrictedProfile": reflect.ValueOf(q.NewRestrictedProfile),
			"Norm":                 reflect.ValueOf(q.Norm),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
