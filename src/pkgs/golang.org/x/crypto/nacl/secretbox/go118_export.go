// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package secretbox

import (
	q "golang.org/x/crypto/nacl/secretbox"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "secretbox",
		Path: "golang.org/x/crypto/nacl/secretbox",
		Deps: map[string]string{
			"golang.org/x/crypto/internal/alias":    "alias",
			"golang.org/x/crypto/internal/poly1305": "poly1305",
			"golang.org/x/crypto/salsa20/salsa":     "salsa",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Open": reflect.ValueOf(q.Open),
			"Seal": reflect.ValueOf(q.Seal),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Overhead": {"untyped int", constant.MakeInt64(int64(q.Overhead))},
		},
	})
}
