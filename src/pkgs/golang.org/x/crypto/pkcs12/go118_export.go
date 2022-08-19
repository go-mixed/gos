// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package pkcs12

import (
	q "golang.org/x/crypto/pkcs12"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pkcs12",
		Path: "golang.org/x/crypto/pkcs12",
		Deps: map[string]string{
			"bytes":            "bytes",
			"crypto/cipher":    "cipher",
			"crypto/des":       "des",
			"crypto/ecdsa":     "ecdsa",
			"crypto/hmac":      "hmac",
			"crypto/rsa":       "rsa",
			"crypto/sha1":      "sha1",
			"crypto/x509":      "x509",
			"crypto/x509/pkix": "pkix",
			"encoding/asn1":    "asn1",
			"encoding/hex":     "hex",
			"encoding/pem":     "pem",
			"errors":           "errors",
			"golang.org/x/crypto/pkcs12/internal/rc2": "rc2",
			"math/big":      "big",
			"unicode/utf16": "utf16",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"NotImplementedError": reflect.TypeOf((*q.NotImplementedError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrDecryption":        reflect.ValueOf(&q.ErrDecryption),
			"ErrIncorrectPassword": reflect.ValueOf(&q.ErrIncorrectPassword),
		},
		Funcs: map[string]reflect.Value{
			"Decode": reflect.ValueOf(q.Decode),
			"ToPEM":  reflect.ValueOf(q.ToPEM),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
