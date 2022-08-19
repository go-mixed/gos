// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package cryptobyte

import (
	q "golang.org/x/crypto/cryptobyte"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cryptobyte",
		Path: "golang.org/x/crypto/cryptobyte",
		Deps: map[string]string{
			"encoding/asn1":                       "asn1",
			"errors":                              "errors",
			"fmt":                                 "fmt",
			"golang.org/x/crypto/cryptobyte/asn1": "asn1",
			"math/big":                            "big",
			"reflect":                             "reflect",
			"time":                                "time",
		},
		Interfaces: map[string]reflect.Type{
			"MarshalingValue": reflect.TypeOf((*q.MarshalingValue)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"BuildError":          reflect.TypeOf((*q.BuildError)(nil)).Elem(),
			"Builder":             reflect.TypeOf((*q.Builder)(nil)).Elem(),
			"BuilderContinuation": reflect.TypeOf((*q.BuilderContinuation)(nil)).Elem(),
			"String":              reflect.TypeOf((*q.String)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewBuilder":      reflect.ValueOf(q.NewBuilder),
			"NewFixedBuilder": reflect.ValueOf(q.NewFixedBuilder),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
