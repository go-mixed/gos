// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package chacha20

import (
	q "golang.org/x/crypto/chacha20"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "chacha20",
		Path: "golang.org/x/crypto/chacha20",
		Deps: map[string]string{
			"crypto/cipher":                      "cipher",
			"encoding/binary":                    "binary",
			"errors":                             "errors",
			"golang.org/x/crypto/internal/alias": "alias",
			"math/bits":                          "bits",
			"runtime":                            "runtime",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Cipher": reflect.TypeOf((*q.Cipher)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"HChaCha20":                reflect.ValueOf(q.HChaCha20),
			"NewUnauthenticatedCipher": reflect.ValueOf(q.NewUnauthenticatedCipher),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"KeySize":    {"untyped int", constant.MakeInt64(int64(q.KeySize))},
			"NonceSize":  {"untyped int", constant.MakeInt64(int64(q.NonceSize))},
			"NonceSizeX": {"untyped int", constant.MakeInt64(int64(q.NonceSizeX))},
		},
	})
}
