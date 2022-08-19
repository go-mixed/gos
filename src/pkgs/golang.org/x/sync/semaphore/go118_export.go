// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package semaphore

import (
	q "golang.org/x/sync/semaphore"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "semaphore",
		Path: "golang.org/x/sync/semaphore",
		Deps: map[string]string{
			"container/list": "list",
			"context":        "context",
			"sync":           "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Weighted": reflect.TypeOf((*q.Weighted)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewWeighted": reflect.ValueOf(q.NewWeighted),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
