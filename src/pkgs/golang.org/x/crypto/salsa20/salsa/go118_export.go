// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package salsa

import (
	q "golang.org/x/crypto/salsa20/salsa"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "salsa",
		Path:       "golang.org/x/crypto/salsa20/salsa",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Sigma": reflect.ValueOf(&q.Sigma),
		},
		Funcs: map[string]reflect.Value{
			"Core208":      reflect.ValueOf(q.Core208),
			"HSalsa20":     reflect.ValueOf(q.HSalsa20),
			"XORKeyStream": reflect.ValueOf(q.XORKeyStream),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
