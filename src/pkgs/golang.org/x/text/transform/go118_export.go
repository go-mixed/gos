// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package transform

import (
	q "golang.org/x/text/transform"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "transform",
		Path: "golang.org/x/text/transform",
		Deps: map[string]string{
			"bytes":        "bytes",
			"errors":       "errors",
			"io":           "io",
			"unicode/utf8": "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"SpanningTransformer": reflect.TypeOf((*q.SpanningTransformer)(nil)).Elem(),
			"Transformer":         reflect.TypeOf((*q.Transformer)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"NopResetter": reflect.TypeOf((*q.NopResetter)(nil)).Elem(),
			"Reader":      reflect.TypeOf((*q.Reader)(nil)).Elem(),
			"Writer":      reflect.TypeOf((*q.Writer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"Discard":      reflect.ValueOf(&q.Discard),
			"ErrEndOfSpan": reflect.ValueOf(&q.ErrEndOfSpan),
			"ErrShortDst":  reflect.ValueOf(&q.ErrShortDst),
			"ErrShortSrc":  reflect.ValueOf(&q.ErrShortSrc),
			"Nop":          reflect.ValueOf(&q.Nop),
		},
		Funcs: map[string]reflect.Value{
			"Append":     reflect.ValueOf(q.Append),
			"Bytes":      reflect.ValueOf(q.Bytes),
			"Chain":      reflect.ValueOf(q.Chain),
			"NewReader":  reflect.ValueOf(q.NewReader),
			"NewWriter":  reflect.ValueOf(q.NewWriter),
			"RemoveFunc": reflect.ValueOf(q.RemoveFunc),
			"String":     reflect.ValueOf(q.String),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
