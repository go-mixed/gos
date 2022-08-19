// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package lzw

import (
	q "golang.org/x/image/tiff/lzw"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "lzw",
		Path: "golang.org/x/image/tiff/lzw",
		Deps: map[string]string{
			"bufio":  "bufio",
			"errors": "errors",
			"fmt":    "fmt",
			"io":     "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Order": reflect.TypeOf((*q.Order)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"NewReader": reflect.ValueOf(q.NewReader),
		},
		TypedConsts: map[string]igop.TypedConst{
			"LSB": {reflect.TypeOf(q.LSB), constant.MakeInt64(int64(q.LSB))},
			"MSB": {reflect.TypeOf(q.MSB), constant.MakeInt64(int64(q.MSB))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
