// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package runenames

import (
	q "golang.org/x/text/unicode/runenames"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "runenames",
		Path: "golang.org/x/text/unicode/runenames",
		Deps: map[string]string{
			"sort": "sort",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Name": reflect.ValueOf(q.Name),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
