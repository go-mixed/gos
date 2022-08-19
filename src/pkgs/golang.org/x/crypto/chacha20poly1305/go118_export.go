// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package chacha20poly1305

import (
	q "golang.org/x/crypto/chacha20poly1305"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "chacha20poly1305",
		Path: "golang.org/x/crypto/chacha20poly1305",
		Deps: map[string]string{
			"crypto/cipher":                         "cipher",
			"encoding/binary":                       "binary",
			"errors":                                "errors",
			"golang.org/x/crypto/chacha20":          "chacha20",
			"golang.org/x/crypto/internal/alias":    "alias",
			"golang.org/x/crypto/internal/poly1305": "poly1305",
			"golang.org/x/sys/cpu":                  "cpu",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"New":  reflect.ValueOf(q.New),
			"NewX": reflect.ValueOf(q.NewX),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"KeySize":    {"untyped int", constant.MakeInt64(int64(q.KeySize))},
			"NonceSize":  {"untyped int", constant.MakeInt64(int64(q.NonceSize))},
			"NonceSizeX": {"untyped int", constant.MakeInt64(int64(q.NonceSizeX))},
			"Overhead":   {"untyped int", constant.MakeInt64(int64(q.Overhead))},
		},
	})
}
