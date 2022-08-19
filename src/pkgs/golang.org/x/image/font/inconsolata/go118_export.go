// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package inconsolata

import (
	q "golang.org/x/image/font/inconsolata"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "inconsolata",
		Path: "golang.org/x/image/font/inconsolata",
		Deps: map[string]string{
			"golang.org/x/image/font/basicfont": "basicfont",
			"image":                             "image",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Bold8x16":    reflect.ValueOf(&q.Bold8x16),
			"Regular8x16": reflect.ValueOf(&q.Regular8x16),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
