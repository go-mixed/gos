// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ripemd160

import (
	q "golang.org/x/crypto/ripemd160"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ripemd160",
		Path: "golang.org/x/crypto/ripemd160",
		Deps: map[string]string{
			"crypto":    "crypto",
			"hash":      "hash",
			"math/bits": "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New": reflect.ValueOf(q.New),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"BlockSize": {"untyped int", constant.MakeInt64(int64(q.BlockSize))},
			"Size":      {"untyped int", constant.MakeInt64(int64(q.Size))},
		},
	})
}
