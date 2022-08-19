// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package bn256

import (
	q "golang.org/x/crypto/bn256"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bn256",
		Path: "golang.org/x/crypto/bn256",
		Deps: map[string]string{
			"crypto/rand": "rand",
			"io":          "io",
			"math/big":    "big",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"G1": reflect.TypeOf((*q.G1)(nil)).Elem(),
			"G2": reflect.TypeOf((*q.G2)(nil)).Elem(),
			"GT": reflect.TypeOf((*q.GT)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Order": reflect.ValueOf(&q.Order),
		},
		Funcs: map[string]reflect.Value{
			"Pair":     reflect.ValueOf(q.Pair),
			"RandomG1": reflect.ValueOf(q.RandomG1),
			"RandomG2": reflect.ValueOf(q.RandomG2),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
