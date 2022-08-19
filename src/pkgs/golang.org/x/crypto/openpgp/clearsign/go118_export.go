// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package clearsign

import (
	q "golang.org/x/crypto/openpgp/clearsign"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "clearsign",
		Path: "golang.org/x/crypto/openpgp/clearsign",
		Deps: map[string]string{
			"bufio":                              "bufio",
			"bytes":                              "bytes",
			"crypto":                             "crypto",
			"fmt":                                "fmt",
			"golang.org/x/crypto/openpgp/armor":  "armor",
			"golang.org/x/crypto/openpgp/errors": "errors",
			"golang.org/x/crypto/openpgp/packet": "packet",
			"hash":                               "hash",
			"io":                                 "io",
			"net/textproto":                      "textproto",
			"strconv":                            "strconv",
			"strings":                            "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Block": reflect.TypeOf((*q.Block)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":      reflect.ValueOf(q.Decode),
			"Encode":      reflect.ValueOf(q.Encode),
			"EncodeMulti": reflect.ValueOf(q.EncodeMulti),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
