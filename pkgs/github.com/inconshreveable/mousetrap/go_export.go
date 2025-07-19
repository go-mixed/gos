// export by github.com/goplus/ixgo/cmd/qexp

package mousetrap

import (
	q "github.com/inconshreveable/mousetrap"

	"reflect"

	"github.com/goplus/ixgo"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "mousetrap",
		Path: "github.com/inconshreveable/mousetrap",
		Deps: map[string]string{
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
		TypedConsts:   map[string]ixgo.TypedConst{},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
