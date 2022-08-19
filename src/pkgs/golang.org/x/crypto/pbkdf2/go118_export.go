// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package pbkdf2

import (
	q "golang.org/x/crypto/pbkdf2"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pbkdf2",
		Path: "golang.org/x/crypto/pbkdf2",
		Deps: map[string]string{
			"crypto/hmac": "hmac",
			"hash":        "hash",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Key": reflect.ValueOf(q.Key),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
