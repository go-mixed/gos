// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package terminal

import (
	q "golang.org/x/crypto/ssh/terminal"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "terminal",
		Path: "golang.org/x/crypto/ssh/terminal",
		Deps: map[string]string{
			"golang.org/x/term": "term",
			"io":                "io",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{
			"EscapeCodes": reflect.TypeOf((*q.EscapeCodes)(nil)).Elem(),
			"State":       reflect.TypeOf((*q.State)(nil)).Elem(),
			"Terminal":    reflect.TypeOf((*q.Terminal)(nil)).Elem(),
		},
		Vars: map[string]reflect.Value{
			"ErrPasteIndicator": reflect.ValueOf(&q.ErrPasteIndicator),
		},
		Funcs: map[string]reflect.Value{
			"GetSize":      reflect.ValueOf(q.GetSize),
			"GetState":     reflect.ValueOf(q.GetState),
			"IsTerminal":   reflect.ValueOf(q.IsTerminal),
			"MakeRaw":      reflect.ValueOf(q.MakeRaw),
			"NewTerminal":  reflect.ValueOf(q.NewTerminal),
			"ReadPassword": reflect.ValueOf(q.ReadPassword),
			"Restore":      reflect.ValueOf(q.Restore),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
