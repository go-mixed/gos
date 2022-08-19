// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package knownhosts

import (
	q "golang.org/x/crypto/ssh/knownhosts"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "knownhosts",
		Path: "golang.org/x/crypto/ssh/knownhosts",
		Deps: map[string]string{
			"bufio":                   "bufio",
			"bytes":                   "bytes",
			"crypto/hmac":             "hmac",
			"crypto/rand":             "rand",
			"crypto/sha1":             "sha1",
			"encoding/base64":         "base64",
			"errors":                  "errors",
			"fmt":                     "fmt",
			"golang.org/x/crypto/ssh": "ssh",
			"io":                      "io",
			"net":                     "net",
			"os":                      "os",
			"strings":                 "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"KeyError":     reflect.TypeOf((*q.KeyError)(nil)).Elem(),
			"KnownKey":     reflect.TypeOf((*q.KnownKey)(nil)).Elem(),
			"RevokedError": reflect.TypeOf((*q.RevokedError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"HashHostname": reflect.ValueOf(q.HashHostname),
			"Line":         reflect.ValueOf(q.Line),
			"New":          reflect.ValueOf(q.New),
			"Normalize":    reflect.ValueOf(q.Normalize),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
