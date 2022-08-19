// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package auth

import (
	q "golang.org/x/crypto/nacl/auth"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "auth",
		Path: "golang.org/x/crypto/nacl/auth",
		Deps: map[string]string{
			"crypto/hmac":   "hmac",
			"crypto/sha512": "sha512",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Sum":    reflect.ValueOf(q.Sum),
			"Verify": reflect.ValueOf(q.Verify),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"KeySize": {"untyped int", constant.MakeInt64(int64(q.KeySize))},
			"Size":    {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
