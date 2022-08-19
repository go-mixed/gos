// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package basicfont

import (
	q "golang.org/x/image/font/basicfont"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "basicfont",
		Path: "golang.org/x/image/font/basicfont",
		Deps: map[string]string{
			"golang.org/x/image/font":       "font",
			"golang.org/x/image/math/fixed": "fixed",
			"image":                         "image",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Face":  reflect.TypeOf((*q.Face)(nil)).Elem(),
			"Range": reflect.TypeOf((*q.Range)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Face7x13": reflect.ValueOf(&q.Face7x13),
		},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
