// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package norm

import (
	q "golang.org/x/text/unicode/norm"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "norm",
		Path: "golang.org/x/text/unicode/norm",
		Deps: map[string]string{
			"encoding/binary":             "binary",
			"fmt":                         "fmt",
			"golang.org/x/text/transform": "transform",
			"io":                          "io",
			"sync":                        "sync",
			"unicode/utf8":                "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Form":       reflect.TypeOf((*q.Form)(nil)).Elem(),
			"Iter":       reflect.TypeOf((*q.Iter)(nil)).Elem(),
			"Properties": reflect.TypeOf((*q.Properties)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs:      map[string]reflect.Value{},
		TypedConsts: map[string]igop.TypedConst{
			"NFC":  {reflect.TypeOf(q.NFC), constant.MakeInt64(int64(q.NFC))},
			"NFD":  {reflect.TypeOf(q.NFD), constant.MakeInt64(int64(q.NFD))},
			"NFKC": {reflect.TypeOf(q.NFKC), constant.MakeInt64(int64(q.NFKC))},
			"NFKD": {reflect.TypeOf(q.NFKD), constant.MakeInt64(int64(q.NFKD))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"GraphemeJoiner":        {"untyped string", constant.MakeString(string(q.GraphemeJoiner))},
			"MaxSegmentSize":        {"untyped int", constant.MakeInt64(int64(q.MaxSegmentSize))},
			"MaxTransformChunkSize": {"untyped int", constant.MakeInt64(int64(q.MaxTransformChunkSize))},
			"Version":               {"untyped string", constant.MakeString(string(q.Version))},
		},
	})
}
