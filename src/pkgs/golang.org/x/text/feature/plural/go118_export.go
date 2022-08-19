// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package plural

import (
	q "golang.org/x/text/feature/plural"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "plural",
		Path: "golang.org/x/text/feature/plural",
		Deps: map[string]string{
			"fmt":                               "fmt",
			"golang.org/x/text/internal/catmsg": "catmsg",
			"golang.org/x/text/internal/language/compact": "compact",
			"golang.org/x/text/internal/number":           "number",
			"golang.org/x/text/language":                  "language",
			"golang.org/x/text/message/catalog":           "catalog",
			"io/ioutil":                                   "ioutil",
			"reflect":                                     "reflect",
			"strconv":                                     "strconv",
		},
		Interfaces: map[string]reflect.Type{
			"Interface": reflect.TypeOf((*q.Interface)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Form":  reflect.TypeOf((*q.Form)(nil)).Elem(),
			"Rules": reflect.TypeOf((*q.Rules)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Cardinal": reflect.ValueOf(&q.Cardinal),
			"Ordinal":  reflect.ValueOf(&q.Ordinal),
		},
		Funcs: map[string]reflect.Value{
			"Selectf": reflect.ValueOf(q.Selectf),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Few":   {reflect.TypeOf(q.Few), constant.MakeInt64(int64(q.Few))},
			"Many":  {reflect.TypeOf(q.Many), constant.MakeInt64(int64(q.Many))},
			"One":   {reflect.TypeOf(q.One), constant.MakeInt64(int64(q.One))},
			"Other": {reflect.TypeOf(q.Other), constant.MakeInt64(int64(q.Other))},
			"Two":   {reflect.TypeOf(q.Two), constant.MakeInt64(int64(q.Two))},
			"Zero":  {reflect.TypeOf(q.Zero), constant.MakeInt64(int64(q.Zero))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"CLDRVersion": {"untyped string", constant.MakeString(string(q.CLDRVersion))},
		},
	})
}
