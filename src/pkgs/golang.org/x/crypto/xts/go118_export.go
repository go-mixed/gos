// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package xts

import (
	q "golang.org/x/crypto/xts"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "xts",
		Path: "golang.org/x/crypto/xts",
		Deps: map[string]string{
			"crypto/cipher":                      "cipher",
			"encoding/binary":                    "binary",
			"errors":                             "errors",
			"golang.org/x/crypto/internal/alias": "alias",
			"sync":                               "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Cipher": reflect.TypeOf((*q.Cipher)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher": reflect.ValueOf(q.NewCipher),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
