// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package hpack

import (
	q "golang.org/x/net/http2/hpack"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "hpack",
		Path: "golang.org/x/net/http2/hpack",
		Deps: map[string]string{
			"bytes":  "bytes",
			"errors": "errors",
			"fmt":    "fmt",
			"io":     "io",
			"sync":   "sync",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Decoder":           reflect.TypeOf((*q.Decoder)(nil)).Elem(),
			"DecodingError":     reflect.TypeOf((*q.DecodingError)(nil)).Elem(),
			"Encoder":           reflect.TypeOf((*q.Encoder)(nil)).Elem(),
			"HeaderField":       reflect.TypeOf((*q.HeaderField)(nil)).Elem(),
			"InvalidIndexError": reflect.TypeOf((*q.InvalidIndexError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrInvalidHuffman": reflect.ValueOf(&q.ErrInvalidHuffman),
			"ErrStringLength":   reflect.ValueOf(&q.ErrStringLength),
		},
		Funcs: map[string]reflect.Value{
			"AppendHuffmanString":   reflect.ValueOf(q.AppendHuffmanString),
			"HuffmanDecode":         reflect.ValueOf(q.HuffmanDecode),
			"HuffmanDecodeToString": reflect.ValueOf(q.HuffmanDecodeToString),
			"HuffmanEncodeLength":   reflect.ValueOf(q.HuffmanEncodeLength),
			"NewDecoder":            reflect.ValueOf(q.NewDecoder),
			"NewEncoder":            reflect.ValueOf(q.NewEncoder),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
