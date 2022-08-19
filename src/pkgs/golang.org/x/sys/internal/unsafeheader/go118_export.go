// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package unsafeheader

import (
	q "golang.org/x/sys/internal/unsafeheader"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "unsafeheader",
		Path: "golang.org/x/sys/internal/unsafeheader",
		Deps: map[string]string{
			"unsafe": "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Slice":  reflect.TypeOf((*q.Slice)(nil)).Elem(),
			"String": reflect.TypeOf((*q.String)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
