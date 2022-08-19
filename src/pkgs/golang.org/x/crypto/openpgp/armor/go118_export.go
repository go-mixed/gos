// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package armor

import (
	q "golang.org/x/crypto/openpgp/armor"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "armor",
		Path: "golang.org/x/crypto/openpgp/armor",
		Deps: map[string]string{
			"bufio":                              "bufio",
			"bytes":                              "bytes",
			"encoding/base64":                    "base64",
			"golang.org/x/crypto/openpgp/errors": "errors",
			"io":                                 "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Block": reflect.TypeOf((*q.Block)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ArmorCorrupt": reflect.ValueOf(&q.ArmorCorrupt),
		},
		Funcs: map[string]reflect.Value{
			"Decode": reflect.ValueOf(q.Decode),
			"Encode": reflect.ValueOf(q.Encode),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
