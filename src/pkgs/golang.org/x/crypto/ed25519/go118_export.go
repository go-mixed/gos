// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ed25519

import (
	q "golang.org/x/crypto/ed25519"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ed25519",
		Path: "golang.org/x/crypto/ed25519",
		Deps: map[string]string{
			"crypto/ed25519": "ed25519",
			"io":             "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"PrivateKey": reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":  reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey":    reflect.ValueOf(q.GenerateKey),
			"NewKeyFromSeed": reflect.ValueOf(q.NewKeyFromSeed),
			"Sign":           reflect.ValueOf(q.Sign),
			"Verify":         reflect.ValueOf(q.Verify),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"PrivateKeySize": {"untyped int", constant.MakeInt64(int64(q.PrivateKeySize))},
			"PublicKeySize":  {"untyped int", constant.MakeInt64(int64(q.PublicKeySize))},
			"SeedSize":       {"untyped int", constant.MakeInt64(int64(q.SeedSize))},
			"SignatureSize":  {"untyped int", constant.MakeInt64(int64(q.SignatureSize))},
		},
	})
}
