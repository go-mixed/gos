// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package tiff

import (
	q "golang.org/x/image/tiff"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "tiff",
		Path: "golang.org/x/image/tiff",
		Deps: map[string]string{
			"bufio":                       "bufio",
			"bytes":                       "bytes",
			"compress/zlib":               "zlib",
			"encoding/binary":             "binary",
			"errors":                      "errors",
			"fmt":                         "fmt",
			"golang.org/x/image/ccitt":    "ccitt",
			"golang.org/x/image/tiff/lzw": "lzw",
			"image":                       "image",
			"image/color":                 "color",
			"io":                          "io",
			"io/ioutil":                   "ioutil",
			"math":                        "math",
			"sort":                        "sort",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"CompressionType":  reflect.TypeOf((*q.CompressionType)(nil)).Elem(),
			"FormatError":      reflect.TypeOf((*q.FormatError)(nil)).Elem(),
			"Options":          reflect.TypeOf((*q.Options)(nil)).Elem(),
			"UnsupportedError": reflect.TypeOf((*q.UnsupportedError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Decode":       reflect.ValueOf(q.Decode),
			"DecodeConfig": reflect.ValueOf(q.DecodeConfig),
			"Encode":       reflect.ValueOf(q.Encode),
		},
		TypedConsts: map[string]igop.TypedConst{
			"CCITTGroup3":  {reflect.TypeOf(q.CCITTGroup3), constant.MakeInt64(int64(q.CCITTGroup3))},
			"CCITTGroup4":  {reflect.TypeOf(q.CCITTGroup4), constant.MakeInt64(int64(q.CCITTGroup4))},
			"Deflate":      {reflect.TypeOf(q.Deflate), constant.MakeInt64(int64(q.Deflate))},
			"LZW":          {reflect.TypeOf(q.LZW), constant.MakeInt64(int64(q.LZW))},
			"Uncompressed": {reflect.TypeOf(q.Uncompressed), constant.MakeInt64(int64(q.Uncompressed))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
