// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package vp8

import (
	q "golang.org/x/image/vp8"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "vp8",
		Path: "golang.org/x/image/vp8",
		Deps: map[string]string{
			"errors": "errors",
			"image":  "image",
			"io":     "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Decoder":     reflect.TypeOf((*q.Decoder)(nil)).Elem(),
			"FrameHeader": reflect.TypeOf((*q.FrameHeader)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewDecoder": reflect.ValueOf(q.NewDecoder),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
