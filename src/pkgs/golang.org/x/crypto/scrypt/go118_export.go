// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package scrypt

import (
	q "golang.org/x/crypto/scrypt"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "scrypt",
		Path: "golang.org/x/crypto/scrypt",
		Deps: map[string]string{
			"crypto/sha256":              "sha256",
			"encoding/binary":            "binary",
			"errors":                     "errors",
			"golang.org/x/crypto/pbkdf2": "pbkdf2",
			"math/bits":                  "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Key": reflect.ValueOf(q.Key),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
