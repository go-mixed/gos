// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package box

import (
	q "golang.org/x/crypto/nacl/box"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "box",
		Path: "golang.org/x/crypto/nacl/box",
		Deps: map[string]string{
			"crypto/rand":                        "rand",
			"golang.org/x/crypto/blake2b":        "blake2b",
			"golang.org/x/crypto/curve25519":     "curve25519",
			"golang.org/x/crypto/nacl/secretbox": "secretbox",
			"golang.org/x/crypto/salsa20/salsa":  "salsa",
			"io":                                 "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey":             reflect.ValueOf(q.GenerateKey),
			"Open":                    reflect.ValueOf(q.Open),
			"OpenAfterPrecomputation": reflect.ValueOf(q.OpenAfterPrecomputation),
			"OpenAnonymous":           reflect.ValueOf(q.OpenAnonymous),
			"Precompute":              reflect.ValueOf(q.Precompute),
			"Seal":                    reflect.ValueOf(q.Seal),
			"SealAfterPrecomputation": reflect.ValueOf(q.SealAfterPrecomputation),
			"SealAnonymous":           reflect.ValueOf(q.SealAnonymous),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"AnonymousOverhead": {"untyped int", constant.MakeInt64(int64(q.AnonymousOverhead))},
			"Overhead":          {"untyped int", constant.MakeInt64(int64(q.Overhead))},
		},
	})
}
