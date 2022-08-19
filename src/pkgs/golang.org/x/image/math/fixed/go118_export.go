// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package fixed

import (
	q "golang.org/x/image/math/fixed"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "fixed",
		Path: "golang.org/x/image/math/fixed",
		Deps: map[string]string{
			"fmt": "fmt",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Int26_6":        reflect.TypeOf((*q.Int26_6)(nil)).Elem(),
			"Int52_12":       reflect.TypeOf((*q.Int52_12)(nil)).Elem(),
			"Point26_6":      reflect.TypeOf((*q.Point26_6)(nil)).Elem(),
			"Point52_12":     reflect.TypeOf((*q.Point52_12)(nil)).Elem(),
			"Rectangle26_6":  reflect.TypeOf((*q.Rectangle26_6)(nil)).Elem(),
			"Rectangle52_12": reflect.TypeOf((*q.Rectangle52_12)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"I": reflect.ValueOf(q.I),
			"P": reflect.ValueOf(q.P),
			"R": reflect.ValueOf(q.R),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
