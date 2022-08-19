// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package h2c

import (
	q "golang.org/x/net/http2/h2c"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "h2c",
		Path: "golang.org/x/net/http2/h2c",
		Deps: map[string]string{
			"bufio":                          "bufio",
			"bytes":                          "bytes",
			"encoding/base64":                "base64",
			"errors":                         "errors",
			"fmt":                            "fmt",
			"golang.org/x/net/http/httpguts": "httpguts",
			"golang.org/x/net/http2":         "http2",
			"io":                             "io",
			"log":                            "log",
			"net":                            "net",
			"net/http":                       "http",
			"net/textproto":                  "textproto",
			"os":                             "os",
			"strings":                        "strings",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewHandler": reflect.ValueOf(q.NewHandler),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
