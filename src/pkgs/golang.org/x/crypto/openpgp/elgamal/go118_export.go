// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package elgamal

import (
	q "golang.org/x/crypto/openpgp/elgamal"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "elgamal",
		Path: "golang.org/x/crypto/openpgp/elgamal",
		Deps: map[string]string{
			"crypto/rand":   "rand",
			"crypto/subtle": "subtle",
			"errors":        "errors",
			"io":            "io",
			"math/big":      "big",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decrypt": reflect.ValueOf(q.Decrypt),
			"Encrypt": reflect.ValueOf(q.Encrypt),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
