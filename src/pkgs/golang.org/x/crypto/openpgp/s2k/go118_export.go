// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package s2k

import (
	q "golang.org/x/crypto/openpgp/s2k"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "s2k",
		Path: "golang.org/x/crypto/openpgp/s2k",
		Deps: map[string]string{
			"crypto":                             "crypto",
			"golang.org/x/crypto/openpgp/errors": "errors",
			"hash":                               "hash",
			"io":                                 "io",
			"strconv":                            "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"HashIdToHash":   reflect.ValueOf(q.HashIdToHash),
			"HashIdToString": reflect.ValueOf(q.HashIdToString),
			"HashToHashId":   reflect.ValueOf(q.HashToHashId),
			"Iterated":       reflect.ValueOf(q.Iterated),
			"Parse":          reflect.ValueOf(q.Parse),
			"Salted":         reflect.ValueOf(q.Salted),
			"Serialize":      reflect.ValueOf(q.Serialize),
			"Simple":         reflect.ValueOf(q.Simple),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
