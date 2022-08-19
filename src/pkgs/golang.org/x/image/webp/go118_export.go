// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package webp

import (
	q "golang.org/x/image/webp"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "webp",
		Path: "golang.org/x/image/webp",
		Deps: map[string]string{
			"bytes":                   "bytes",
			"errors":                  "errors",
			"golang.org/x/image/riff": "riff",
			"golang.org/x/image/vp8":  "vp8",
			"golang.org/x/image/vp8l": "vp8l",
			"image":                   "image",
			"image/color":             "color",
			"io":                      "io",
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
