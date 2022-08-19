// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package syncmap

import (
	q "golang.org/x/sync/syncmap"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "syncmap",
		Path: "golang.org/x/sync/syncmap",
		Deps: map[string]string{
			"sync": "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"Map": reflect.TypeOf((*q.Map)(nil)).Elem(),
		},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
