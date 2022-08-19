// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package vp8l

import (
	q "golang.org/x/image/vp8l"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "vp8l",
		Path: "golang.org/x/image/vp8l",
		Deps: map[string]string{
			"bufio":       "bufio",
			"errors":      "errors",
			"image":       "image",
			"image/color": "color",
			"io":          "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
