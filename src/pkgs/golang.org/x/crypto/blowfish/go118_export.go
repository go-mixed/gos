// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package blowfish

import (
	q "golang.org/x/crypto/blowfish"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "blowfish",
		Path: "golang.org/x/crypto/blowfish",
		Deps: map[string]string{
			"strconv": "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Cipher":       reflect.TypeOf((*q.Cipher)(nil)).Elem(),
			"KeySizeError": reflect.TypeOf((*q.KeySizeError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"ExpandKey":       reflect.ValueOf(q.ExpandKey),
			"NewCipher":       reflect.ValueOf(q.NewCipher),
			"NewSaltedCipher": reflect.ValueOf(q.NewSaltedCipher),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
		},
	})
}
