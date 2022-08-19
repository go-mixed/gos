// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package draw

import (
	q "golang.org/x/image/draw"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "draw",
		Path: "golang.org/x/image/draw",
		Deps: map[string]string{
			"golang.org/x/image/math/f64": "f64",
			"image":                       "image",
			"image/color":                 "color",
			"image/draw":                  "draw",
			"math":                        "math",
			"sync":                        "sync",
		},
		Interfaces: map[string]reflect.Type{
			"Interpolator": reflect.TypeOf((*q.Interpolator)(nil)).Elem(),
			"Scaler":       reflect.TypeOf((*q.Scaler)(nil)).Elem(),
			"Transformer":  reflect.TypeOf((*q.Transformer)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Kernel":  reflect.TypeOf((*q.Kernel)(nil)).Elem(),
			"Options": reflect.TypeOf((*q.Options)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Drawer":      reflect.TypeOf((*q.Drawer)(nil)).Elem(),
			"Image":       reflect.TypeOf((*q.Image)(nil)).Elem(),
			"Op":          reflect.TypeOf((*q.Op)(nil)).Elem(),
			"Quantizer":   reflect.TypeOf((*q.Quantizer)(nil)).Elem(),
			"RGBA64Image": reflect.TypeOf((*q.RGBA64Image)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ApproxBiLinear":  reflect.ValueOf(&q.ApproxBiLinear),
			"BiLinear":        reflect.ValueOf(&q.BiLinear),
			"CatmullRom":      reflect.ValueOf(&q.CatmullRom),
			"FloydSteinberg":  reflect.ValueOf(&q.FloydSteinberg),
			"NearestNeighbor": reflect.ValueOf(&q.NearestNeighbor),
		},
		Funcs: map[string]reflect.Value{
			"Copy":     reflect.ValueOf(q.Copy),
			"Draw":     reflect.ValueOf(q.Draw),
			"DrawMask": reflect.ValueOf(q.DrawMask),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Over": {reflect.TypeOf(q.Over), constant.MakeInt64(int64(q.Over))},
			"Src":  {reflect.TypeOf(q.Src), constant.MakeInt64(int64(q.Src))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
