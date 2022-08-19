// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package cast5

import (
	q "golang.org/x/crypto/cast5"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cast5",
		Path: "golang.org/x/crypto/cast5",
		Deps: map[string]string{
			"errors": "errors",
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
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"KeySize":   {"untyped int", constant.MakeInt64(int64(q.KeySize))},
		},
	})
}
