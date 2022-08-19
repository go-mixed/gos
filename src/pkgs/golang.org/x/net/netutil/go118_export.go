// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package netutil

import (
	q "golang.org/x/net/netutil"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "netutil",
		Path: "golang.org/x/net/netutil",
		Deps: map[string]string{
			"net":  "net",
			"sync": "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"LimitListener": reflect.ValueOf(q.LimitListener),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
