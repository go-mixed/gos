// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package otr

import (
	q "golang.org/x/crypto/otr"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "otr",
		Path: "golang.org/x/crypto/otr",
		Deps: map[string]string{
			"bytes":           "bytes",
			"crypto/aes":      "aes",
			"crypto/cipher":   "cipher",
			"crypto/dsa":      "dsa",
			"crypto/hmac":     "hmac",
			"crypto/rand":     "rand",
			"crypto/sha1":     "sha1",
			"crypto/sha256":   "sha256",
			"crypto/subtle":   "subtle",
			"encoding/base64": "base64",
			"encoding/hex":    "hex",
			"errors":          "errors",
			"hash":            "hash",
			"io":              "io",
			"math/big":        "big",
			"strconv":         "strconv",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Conversation":   reflect.TypeOf((*q.Conversation)(nil)).Elem(),
			"PrivateKey":     reflect.TypeOf((*q.PrivateKey)(nil)).Elem(),
			"PublicKey":      reflect.TypeOf((*q.PublicKey)(nil)).Elem(),
			"SecurityChange": reflect.TypeOf((*q.SecurityChange)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrorPrefix":  reflect.ValueOf(&q.ErrorPrefix),
			"QueryMessage": reflect.ValueOf(&q.QueryMessage),
		},
		Funcs: map[string]reflect.Value{},
		TypedConsts: map[string]igop.TypedConst{
			"ConversationEnded": {reflect.TypeOf(q.ConversationEnded), constant.MakeInt64(int64(q.ConversationEnded))},
			"NewKeys":           {reflect.TypeOf(q.NewKeys), constant.MakeInt64(int64(q.NewKeys))},
			"NoChange":          {reflect.TypeOf(q.NoChange), constant.MakeInt64(int64(q.NoChange))},
			"SMPComplete":       {reflect.TypeOf(q.SMPComplete), constant.MakeInt64(int64(q.SMPComplete))},
			"SMPFailed":         {reflect.TypeOf(q.SMPFailed), constant.MakeInt64(int64(q.SMPFailed))},
			"SMPSecretNeeded":   {reflect.TypeOf(q.SMPSecretNeeded), constant.MakeInt64(int64(q.SMPSecretNeeded))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
