// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package bidi

import (
	q "golang.org/x/text/unicode/bidi"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "bidi",
		Path: "golang.org/x/text/unicode/bidi",
		Deps: map[string]string{
			"bytes":          "bytes",
			"container/list": "list",
			"fmt":            "fmt",
			"log":            "log",
			"sort":           "sort",
			"unicode/utf8":   "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Class":      reflect.TypeOf((*q.Class)(nil)).Elem(),
			"Direction":  reflect.TypeOf((*q.Direction)(nil)).Elem(),
			"Option":     reflect.TypeOf((*q.Option)(nil)).Elem(),
			"Ordering":   reflect.TypeOf((*q.Ordering)(nil)).Elem(),
			"Paragraph":  reflect.TypeOf((*q.Paragraph)(nil)).Elem(),
			"Properties": reflect.TypeOf((*q.Properties)(nil)).Elem(),
			"Run":        reflect.TypeOf((*q.Run)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"AppendReverse":    reflect.ValueOf(q.AppendReverse),
			"DefaultDirection": reflect.ValueOf(q.DefaultDirection),
			"Lookup":           reflect.ValueOf(q.Lookup),
			"LookupRune":       reflect.ValueOf(q.LookupRune),
			"LookupString":     reflect.ValueOf(q.LookupString),
			"ReverseString":    reflect.ValueOf(q.ReverseString),
		},
		TypedConsts: map[string]igop.TypedConst{
			"AL":          {reflect.TypeOf(q.AL), constant.MakeInt64(int64(q.AL))},
			"AN":          {reflect.TypeOf(q.AN), constant.MakeInt64(int64(q.AN))},
			"B":           {reflect.TypeOf(q.B), constant.MakeInt64(int64(q.B))},
			"BN":          {reflect.TypeOf(q.BN), constant.MakeInt64(int64(q.BN))},
			"CS":          {reflect.TypeOf(q.CS), constant.MakeInt64(int64(q.CS))},
			"Control":     {reflect.TypeOf(q.Control), constant.MakeInt64(int64(q.Control))},
			"EN":          {reflect.TypeOf(q.EN), constant.MakeInt64(int64(q.EN))},
			"ES":          {reflect.TypeOf(q.ES), constant.MakeInt64(int64(q.ES))},
			"ET":          {reflect.TypeOf(q.ET), constant.MakeInt64(int64(q.ET))},
			"FSI":         {reflect.TypeOf(q.FSI), constant.MakeInt64(int64(q.FSI))},
			"L":           {reflect.TypeOf(q.L), constant.MakeInt64(int64(q.L))},
			"LRE":         {reflect.TypeOf(q.LRE), constant.MakeInt64(int64(q.LRE))},
			"LRI":         {reflect.TypeOf(q.LRI), constant.MakeInt64(int64(q.LRI))},
			"LRO":         {reflect.TypeOf(q.LRO), constant.MakeInt64(int64(q.LRO))},
			"LeftToRight": {reflect.TypeOf(q.LeftToRight), constant.MakeInt64(int64(q.LeftToRight))},
			"Mixed":       {reflect.TypeOf(q.Mixed), constant.MakeInt64(int64(q.Mixed))},
			"NSM":         {reflect.TypeOf(q.NSM), constant.MakeInt64(int64(q.NSM))},
			"Neutral":     {reflect.TypeOf(q.Neutral), constant.MakeInt64(int64(q.Neutral))},
			"ON":          {reflect.TypeOf(q.ON), constant.MakeInt64(int64(q.ON))},
			"PDF":         {reflect.TypeOf(q.PDF), constant.MakeInt64(int64(q.PDF))},
			"PDI":         {reflect.TypeOf(q.PDI), constant.MakeInt64(int64(q.PDI))},
			"R":           {reflect.TypeOf(q.R), constant.MakeInt64(int64(q.R))},
			"RLE":         {reflect.TypeOf(q.RLE), constant.MakeInt64(int64(q.RLE))},
			"RLI":         {reflect.TypeOf(q.RLI), constant.MakeInt64(int64(q.RLI))},
			"RLO":         {reflect.TypeOf(q.RLO), constant.MakeInt64(int64(q.RLO))},
			"RightToLeft": {reflect.TypeOf(q.RightToLeft), constant.MakeInt64(int64(q.RightToLeft))},
			"S":           {reflect.TypeOf(q.S), constant.MakeInt64(int64(q.S))},
			"WS":          {reflect.TypeOf(q.WS), constant.MakeInt64(int64(q.WS))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"UnicodeVersion": {"untyped string", constant.MakeString(string(q.UnicodeVersion))},
		},
	})
}
