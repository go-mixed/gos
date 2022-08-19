// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package rangetable

import (
	q "golang.org/x/text/unicode/rangetable"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "rangetable",
		Path: "golang.org/x/text/unicode/rangetable",
		Deps: map[string]string{
			"sort":    "sort",
			"unicode": "unicode",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Assigned": reflect.ValueOf(q.Assigned),
			"Merge":    reflect.ValueOf(q.Merge),
			"New":      reflect.ValueOf(q.New),
			"Visit":    reflect.ValueOf(q.Visit),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
