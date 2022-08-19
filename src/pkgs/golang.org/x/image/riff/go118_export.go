// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package riff

import (
	q "golang.org/x/image/riff"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "riff",
		Path: "golang.org/x/image/riff",
		Deps: map[string]string{
			"errors":    "errors",
			"io":        "io",
			"io/ioutil": "ioutil",
			"math":      "math",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"FourCC": reflect.TypeOf((*q.FourCC)(nil)).Elem(),
			"Reader": reflect.TypeOf((*q.Reader)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"LIST": reflect.ValueOf(&q.LIST),
		},
		Funcs: map[string]reflect.Value{
			"NewListReader": reflect.ValueOf(q.NewListReader),
			"NewReader":     reflect.ValueOf(q.NewReader),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
