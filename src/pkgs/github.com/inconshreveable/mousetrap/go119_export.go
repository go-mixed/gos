// export by github.com/goplus/igop/cmd/qexp

//go:build go1.19
// +build go1.19

package mousetrap

import (
	q "github.com/inconshreveable/mousetrap"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "mousetrap",
		Path: "github.com/inconshreveable/mousetrap",
		Deps: map[string]string{
			"os":      "os",
			"syscall": "syscall",
			"unsafe":  "unsafe",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"StartedByExplorer": reflect.ValueOf(q.StartedByExplorer),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
