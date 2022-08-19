// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package width

import (
	q "golang.org/x/text/width"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "width",
		Path: "golang.org/x/text/width",
		Deps: map[string]string{
			"golang.org/x/text/transform": "transform",
			"strconv":                     "strconv",
			"unicode/utf8":                "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Kind":        reflect.TypeOf((*q.Kind)(nil)).Elem(),
			"Properties":  reflect.TypeOf((*q.Properties)(nil)).Elem(),
			"Transformer": reflect.TypeOf((*q.Transformer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Fold":   reflect.ValueOf(&q.Fold),
			"Narrow": reflect.ValueOf(&q.Narrow),
			"Widen":  reflect.ValueOf(&q.Widen),
		},
		Funcs: map[string]reflect.Value{
			"Lookup":       reflect.ValueOf(q.Lookup),
			"LookupRune":   reflect.ValueOf(q.LookupRune),
			"LookupString": reflect.ValueOf(q.LookupString),
		},
		TypedConsts: map[string]igop.TypedConst{
			"EastAsianAmbiguous": {reflect.TypeOf(q.EastAsianAmbiguous), constant.MakeInt64(int64(q.EastAsianAmbiguous))},
			"EastAsianFullwidth": {reflect.TypeOf(q.EastAsianFullwidth), constant.MakeInt64(int64(q.EastAsianFullwidth))},
			"EastAsianHalfwidth": {reflect.TypeOf(q.EastAsianHalfwidth), constant.MakeInt64(int64(q.EastAsianHalfwidth))},
			"EastAsianNarrow":    {reflect.TypeOf(q.EastAsianNarrow), constant.MakeInt64(int64(q.EastAsianNarrow))},
			"EastAsianWide":      {reflect.TypeOf(q.EastAsianWide), constant.MakeInt64(int64(q.EastAsianWide))},
			"Neutral":            {reflect.TypeOf(q.Neutral), constant.MakeInt64(int64(q.Neutral))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
