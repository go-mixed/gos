// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package salsa20

import (
	q "golang.org/x/crypto/salsa20"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "salsa20",
		Path: "golang.org/x/crypto/salsa20",
		Deps: map[string]string{
			"golang.org/x/crypto/internal/alias": "alias",
			"golang.org/x/crypto/salsa20/salsa":  "salsa",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"XORKeyStream": reflect.ValueOf(q.XORKeyStream),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
