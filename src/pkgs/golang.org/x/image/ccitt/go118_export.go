// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ccitt

import (
	q "golang.org/x/image/ccitt"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ccitt",
		Path: "golang.org/x/image/ccitt",
		Deps: map[string]string{
			"encoding/binary": "binary",
			"errors":          "errors",
			"image":           "image",
			"io":              "io",
			"math/bits":       "bits",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Options":   reflect.TypeOf((*q.Options)(nil)).Elem(),
			"Order":     reflect.TypeOf((*q.Order)(nil)).Elem(),
			"SubFormat": reflect.TypeOf((*q.SubFormat)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"DecodeIntoGray": reflect.ValueOf(q.DecodeIntoGray),
			"NewReader":      reflect.ValueOf(q.NewReader),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Group3": {reflect.TypeOf(q.Group3), constant.MakeInt64(int64(q.Group3))},
			"Group4": {reflect.TypeOf(q.Group4), constant.MakeInt64(int64(q.Group4))},
			"LSB":    {reflect.TypeOf(q.LSB), constant.MakeInt64(int64(q.LSB))},
			"MSB":    {reflect.TypeOf(q.MSB), constant.MakeInt64(int64(q.MSB))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"AutoDetectHeight": {"untyped int", constant.MakeInt64(int64(q.AutoDetectHeight))},
		},
	})
}
