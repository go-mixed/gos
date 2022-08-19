// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package hkdf

import (
	q "golang.org/x/crypto/hkdf"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "hkdf",
		Path: "golang.org/x/crypto/hkdf",
		Deps: map[string]string{
			"crypto/hmac": "hmac",
			"errors":      "errors",
			"hash":        "hash",
			"io":          "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Expand":  reflect.ValueOf(q.Expand),
			"Extract": reflect.ValueOf(q.Extract),
			"New":     reflect.ValueOf(q.New),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
