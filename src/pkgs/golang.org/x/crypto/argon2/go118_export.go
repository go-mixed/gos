// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package argon2

import (
	q "golang.org/x/crypto/argon2"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "argon2",
		Path: "golang.org/x/crypto/argon2",
		Deps: map[string]string{
			"encoding/binary":             "binary",
			"golang.org/x/crypto/blake2b": "blake2b",
			"golang.org/x/sys/cpu":        "cpu",
			"hash":                        "hash",
			"sync":                        "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"IDKey": reflect.ValueOf(q.IDKey),
			"Key":   reflect.ValueOf(q.Key),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"Version": {"untyped int", constant.MakeInt64(int64(q.Version))},
		},
	})
}
