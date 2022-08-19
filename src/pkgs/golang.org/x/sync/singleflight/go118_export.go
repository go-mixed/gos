// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package singleflight

import (
	q "golang.org/x/sync/singleflight"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "singleflight",
		Path: "golang.org/x/sync/singleflight",
		Deps: map[string]string{
			"bytes":         "bytes",
			"errors":        "errors",
			"fmt":           "fmt",
			"runtime":       "runtime",
			"runtime/debug": "debug",
			"sync":          "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Group":  reflect.TypeOf((*q.Group)(nil)).Elem(),
			"Result": reflect.TypeOf((*q.Result)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
