// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package httpguts

import (
	q "golang.org/x/net/http/httpguts"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "httpguts",
		Path: "golang.org/x/net/http/httpguts",
		Deps: map[string]string{
			"golang.org/x/net/idna": "idna",
			"net":                   "net",
			"net/textproto":         "textproto",
			"strings":               "strings",
			"unicode/utf8":          "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"HeaderValuesContainsToken": reflect.ValueOf(q.HeaderValuesContainsToken),
			"IsTokenRune":               reflect.ValueOf(q.IsTokenRune),
			"PunycodeHostPort":          reflect.ValueOf(q.PunycodeHostPort),
			"ValidHeaderFieldName":      reflect.ValueOf(q.ValidHeaderFieldName),
			"ValidHeaderFieldValue":     reflect.ValueOf(q.ValidHeaderFieldValue),
			"ValidHostHeader":           reflect.ValueOf(q.ValidHostHeader),
			"ValidTrailerHeader":        reflect.ValueOf(q.ValidTrailerHeader),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
