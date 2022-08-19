// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package bmp

import (
	q "golang.org/x/image/bmp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bmp",
		Path: "golang.org/x/image/bmp",
		Deps: map[string]string{
			"encoding/binary": "binary",
			"errors":          "errors",
			"image":           "image",
			"image/color":     "color",
			"io":              "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrUnsupported": reflect.ValueOf(&q.ErrUnsupported),
		},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
			"Encode":       reflect.ValueOf(q.Encode),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
