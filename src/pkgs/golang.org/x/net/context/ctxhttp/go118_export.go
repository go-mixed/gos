// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ctxhttp

import (
	q "golang.org/x/net/context/ctxhttp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ctxhttp",
		Path: "golang.org/x/net/context/ctxhttp",
		Deps: map[string]string{
			"context":  "context",
			"io":       "io",
			"net/http": "http",
			"net/url":  "url",
			"strings":  "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Do":       reflect.ValueOf(q.Do),
			"Get":      reflect.ValueOf(q.Get),
			"Head":     reflect.ValueOf(q.Head),
			"Post":     reflect.ValueOf(q.Post),
			"PostForm": reflect.ValueOf(q.PostForm),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
