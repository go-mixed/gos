// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package f32

import (
	q "golang.org/x/image/math/f32"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "f32",
		Path:       "golang.org/x/image/math/f32",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Aff3": reflect.TypeOf((*q.Aff3)(nil)).Elem(),
			"Aff4": reflect.TypeOf((*q.Aff4)(nil)).Elem(),
			"Mat3": reflect.TypeOf((*q.Mat3)(nil)).Elem(),
			"Mat4": reflect.TypeOf((*q.Mat4)(nil)).Elem(),
			"Vec2": reflect.TypeOf((*q.Vec2)(nil)).Elem(),
			"Vec3": reflect.TypeOf((*q.Vec3)(nil)).Elem(),
			"Vec4": reflect.TypeOf((*q.Vec4)(nil)).Elem(),
		},
		AliasTypes:    map[string]reflect.Type{},
		Vars:          map[string]reflect.Value{},
		Funcs:         map[string]reflect.Value{},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
