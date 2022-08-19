// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package catalog

import (
	q "golang.org/x/text/message/catalog"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "catalog",
		Path: "golang.org/x/text/message/catalog",
		Deps: map[string]string{
			"errors":                            "errors",
			"fmt":                               "fmt",
			"golang.org/x/text/internal":        "internal",
			"golang.org/x/text/internal/catmsg": "catmsg",
			"golang.org/x/text/language":        "language",
			"sync":                              "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Catalog":    reflect.TypeOf((*q.Catalog)(nil)).Elem(),
			"Dictionary": reflect.TypeOf((*q.Dictionary)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Builder": reflect.TypeOf((*q.Builder)(nil)).Elem(),
			"Context": reflect.TypeOf((*q.Context)(nil)).Elem(),
			"Option":  reflect.TypeOf((*q.Option)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Message": reflect.TypeOf((*q.Message)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrNotFound": reflect.ValueOf(&q.ErrNotFound),
		},
		Funcs: map[string]reflect.Value{
			"Fallback":   reflect.ValueOf(q.Fallback),
			"NewBuilder": reflect.ValueOf(q.NewBuilder),
			"NewFromMap": reflect.ValueOf(q.NewFromMap),
			"String":     reflect.ValueOf(q.String),
			"Var":        reflect.ValueOf(q.Var),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
