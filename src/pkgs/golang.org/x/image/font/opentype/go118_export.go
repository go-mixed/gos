// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package opentype

import (
	q "golang.org/x/image/font/opentype"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "opentype",
		Path: "golang.org/x/image/font/opentype",
		Deps: map[string]string{
			"golang.org/x/image/font":       "font",
			"golang.org/x/image/font/sfnt":  "sfnt",
			"golang.org/x/image/math/fixed": "fixed",
			"golang.org/x/image/vector":     "vector",
			"image":                         "image",
			"image/draw":                    "draw",
			"io":                            "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Face":        reflect.TypeOf((*q.Face)(nil)).Elem(),
			"FaceOptions": reflect.TypeOf((*q.FaceOptions)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{
			"Collection": reflect.TypeOf((*q.Collection)(nil)).Elem(),
			"Font":       reflect.TypeOf((*q.Font)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewFace":                 reflect.ValueOf(q.NewFace),
			"Parse":                   reflect.ValueOf(q.Parse),
			"ParseCollection":         reflect.ValueOf(q.ParseCollection),
			"ParseCollectionReaderAt": reflect.ValueOf(q.ParseCollectionReaderAt),
			"ParseReaderAt":           reflect.ValueOf(q.ParseReaderAt),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
