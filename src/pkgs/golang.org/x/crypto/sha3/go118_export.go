// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package sha3

import (
	q "golang.org/x/crypto/sha3"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "sha3",
		Path: "golang.org/x/crypto/sha3",
		Deps: map[string]string{
			"crypto":          "crypto",
			"encoding/binary": "binary",
			"hash":            "hash",
			"io":              "io",
			"unsafe":          "unsafe",
		},
		Interfaces: map[string]reflect.Type{
			"ShakeHash": reflect.TypeOf((*q.ShakeHash)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New224":             reflect.ValueOf(q.New224),
			"New256":             reflect.ValueOf(q.New256),
			"New384":             reflect.ValueOf(q.New384),
			"New512":             reflect.ValueOf(q.New512),
			"NewCShake128":       reflect.ValueOf(q.NewCShake128),
			"NewCShake256":       reflect.ValueOf(q.NewCShake256),
			"NewLegacyKeccak256": reflect.ValueOf(q.NewLegacyKeccak256),
			"NewLegacyKeccak512": reflect.ValueOf(q.NewLegacyKeccak512),
			"NewShake128":        reflect.ValueOf(q.NewShake128),
			"NewShake256":        reflect.ValueOf(q.NewShake256),
			"ShakeSum128":        reflect.ValueOf(q.ShakeSum128),
			"ShakeSum256":        reflect.ValueOf(q.ShakeSum256),
			"Sum224":             reflect.ValueOf(q.Sum224),
			"Sum256":             reflect.ValueOf(q.Sum256),
			"Sum384":             reflect.ValueOf(q.Sum384),
			"Sum512":             reflect.ValueOf(q.Sum512),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
