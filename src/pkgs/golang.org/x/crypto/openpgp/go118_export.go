// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package openpgp

import (
	q "golang.org/x/crypto/openpgp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "openpgp",
		Path: "golang.org/x/crypto/openpgp",
		Deps: map[string]string{
			"crypto":                             "crypto",
			"crypto/rsa":                         "rsa",
			"crypto/sha256":                      "sha256",
			"golang.org/x/crypto/openpgp/armor":  "armor",
			"golang.org/x/crypto/openpgp/errors": "errors",
			"golang.org/x/crypto/openpgp/packet": "packet",
			"golang.org/x/crypto/openpgp/s2k":    "s2k",
			"hash":                               "hash",
			"io":                                 "io",
			"strconv":                            "strconv",
			"time":                               "time",
		},
		Interfaces: map[string]reflect.Type{
			"KeyRing": reflect.TypeOf((*q.KeyRing)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Entity":         reflect.TypeOf((*q.Entity)(nil)).Elem(),
			"EntityList":     reflect.TypeOf((*q.EntityList)(nil)).Elem(),
			"FileHints":      reflect.TypeOf((*q.FileHints)(nil)).Elem(),
			"Identity":       reflect.TypeOf((*q.Identity)(nil)).Elem(),
			"Key":            reflect.TypeOf((*q.Key)(nil)).Elem(),
			"MessageDetails": reflect.TypeOf((*q.MessageDetails)(nil)).Elem(),
			"PromptFunction": reflect.TypeOf((*q.PromptFunction)(nil)).Elem(),
			"Subkey":         reflect.TypeOf((*q.Subkey)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"PrivateKeyType": reflect.ValueOf(&q.PrivateKeyType),
			"PublicKeyType":  reflect.ValueOf(&q.PublicKeyType),
			"SignatureType":  reflect.ValueOf(&q.SignatureType),
		},
		Funcs: map[string]reflect.Value{
			"ArmoredDetachSign":             reflect.ValueOf(q.ArmoredDetachSign),
			"ArmoredDetachSignText":         reflect.ValueOf(q.ArmoredDetachSignText),
			"CheckArmoredDetachedSignature": reflect.ValueOf(q.CheckArmoredDetachedSignature),
			"CheckDetachedSignature":        reflect.ValueOf(q.CheckDetachedSignature),
			"DetachSign":                    reflect.ValueOf(q.DetachSign),
			"DetachSignText":                reflect.ValueOf(q.DetachSignText),
			"Encrypt":                       reflect.ValueOf(q.Encrypt),
			"NewCanonicalTextHash":          reflect.ValueOf(q.NewCanonicalTextHash),
			"NewEntity":                     reflect.ValueOf(q.NewEntity),
			"ReadArmoredKeyRing":            reflect.ValueOf(q.ReadArmoredKeyRing),
			"ReadEntity":                    reflect.ValueOf(q.ReadEntity),
			"ReadKeyRing":                   reflect.ValueOf(q.ReadKeyRing),
			"ReadMessage":                   reflect.ValueOf(q.ReadMessage),
			"Sign":                          reflect.ValueOf(q.Sign),
			"SymmetricallyEncrypt":          reflect.ValueOf(q.SymmetricallyEncrypt),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
