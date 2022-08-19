// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package internal

import (
	q "golang.org/x/text/encoding/internal"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "internal",
		Path: "golang.org/x/text/encoding/internal",
		Deps: map[string]string{
			"golang.org/x/text/encoding":                     "encoding",
			"golang.org/x/text/encoding/internal/identifier": "identifier",
			"golang.org/x/text/transform":                    "transform",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Encoding":        reflect.TypeOf((*q.Encoding)(nil)).Elem(),
			"FuncEncoding":    reflect.TypeOf((*q.FuncEncoding)(nil)).Elem(),
			"RepertoireError": reflect.TypeOf((*q.RepertoireError)(nil)).Elem(),
			"SimpleEncoding":  reflect.TypeOf((*q.SimpleEncoding)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrASCIIReplacement": reflect.ValueOf(&q.ErrASCIIReplacement),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
