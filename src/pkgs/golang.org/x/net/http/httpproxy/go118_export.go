// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package httpproxy

import (
	q "golang.org/x/net/http/httpproxy"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "httpproxy",
		Path: "golang.org/x/net/http/httpproxy",
		Deps: map[string]string{
			"errors":                "errors",
			"fmt":                   "fmt",
			"golang.org/x/net/idna": "idna",
			"net":                   "net",
			"net/url":               "url",
			"os":                    "os",
			"strings":               "strings",
			"unicode/utf8":          "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Config": reflect.TypeOf((*q.Config)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"FromEnvironment": reflect.ValueOf(q.FromEnvironment),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
