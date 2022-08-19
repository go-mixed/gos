// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package tea

import (
	q "golang.org/x/crypto/tea"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "tea",
		Path: "golang.org/x/crypto/tea",
		Deps: map[string]string{
			"crypto/cipher":   "cipher",
			"encoding/binary": "binary",
			"errors":          "errors",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewCipher":           reflect.ValueOf(q.NewCipher),
			"NewCipherWithRounds": reflect.ValueOf(q.NewCipherWithRounds),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"KeySize":   {"untyped int", constant.MakeInt64(int64(q.KeySize))},
		},
	})
}
