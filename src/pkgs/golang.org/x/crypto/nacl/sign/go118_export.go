// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package sign

import (
	q "golang.org/x/crypto/nacl/sign"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "sign",
		Path: "golang.org/x/crypto/nacl/sign",
		Deps: map[string]string{
			"golang.org/x/crypto/ed25519":        "ed25519",
			"golang.org/x/crypto/internal/alias": "alias",
			"io":                                 "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"GenerateKey": reflect.ValueOf(q.GenerateKey),
			"Open":        reflect.ValueOf(q.Open),
			"Sign":        reflect.ValueOf(q.Sign),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Overhead": {"untyped int", constant.MakeInt64(int64(q.Overhead))},
		},
	})
}
