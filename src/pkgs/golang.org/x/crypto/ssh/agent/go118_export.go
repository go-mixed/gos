// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package agent

import (
	q "golang.org/x/crypto/ssh/agent"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "agent",
		Path: "golang.org/x/crypto/ssh/agent",
		Deps: map[string]string{
			"bytes":                       "bytes",
			"crypto/dsa":                  "dsa",
			"crypto/ecdsa":                "ecdsa",
			"crypto/elliptic":             "elliptic",
			"crypto/rand":                 "rand",
			"crypto/rsa":                  "rsa",
			"crypto/subtle":               "subtle",
			"encoding/base64":             "base64",
			"encoding/binary":             "binary",
			"errors":                      "errors",
			"fmt":                         "fmt",
			"golang.org/x/crypto/ed25519": "ed25519",
			"golang.org/x/crypto/ssh":     "ssh",
			"io":                          "io",
			"log":                         "log",
			"math/big":                    "big",
			"net":                         "net",
			"sync":                        "sync",
			"time":                        "time",
		},
		Interfaces: map[string]reflect.Type{
			"Agent":         reflect.TypeOf((*q.Agent)(nil)).Elem(),
			"ExtendedAgent": reflect.TypeOf((*q.ExtendedAgent)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"AddedKey":            reflect.TypeOf((*q.AddedKey)(nil)).Elem(),
			"ConstraintExtension": reflect.TypeOf((*q.ConstraintExtension)(nil)).Elem(),
			"Key":                 reflect.TypeOf((*q.Key)(nil)).Elem(),
			"SignatureFlags":      reflect.TypeOf((*q.SignatureFlags)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrExtensionUnsupported": reflect.ValueOf(&q.ErrExtensionUnsupported),
		},
		Funcs: map[string]reflect.Value{
			"ForwardToAgent":         reflect.ValueOf(q.ForwardToAgent),
			"ForwardToRemote":        reflect.ValueOf(q.ForwardToRemote),
			"NewClient":              reflect.ValueOf(q.NewClient),
			"NewKeyring":             reflect.ValueOf(q.NewKeyring),
			"RequestAgentForwarding": reflect.ValueOf(q.RequestAgentForwarding),
			"ServeAgent":             reflect.ValueOf(q.ServeAgent),
		},
		TypedConsts: map[string]igop.TypedConst{
			"SignatureFlagReserved":  {reflect.TypeOf(q.SignatureFlagReserved), constant.MakeInt64(int64(q.SignatureFlagReserved))},
			"SignatureFlagRsaSha256": {reflect.TypeOf(q.SignatureFlagRsaSha256), constant.MakeInt64(int64(q.SignatureFlagRsaSha256))},
			"SignatureFlagRsaSha512": {reflect.TypeOf(q.SignatureFlagRsaSha512), constant.MakeInt64(int64(q.SignatureFlagRsaSha512))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
