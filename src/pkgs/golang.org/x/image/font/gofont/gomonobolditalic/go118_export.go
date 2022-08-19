// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package gomonobolditalic

import (
	q "golang.org/x/image/font/gofont/gomonobolditalic"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "gomonobolditalic",
		Path:       "golang.org/x/image/font/gofont/gomonobolditalic",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"TTF": reflect.ValueOf(&q.TTF),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
