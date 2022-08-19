// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package html

import (
	q "golang.org/x/net/html"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "html",
		Path: "golang.org/x/net/html",
		Deps: map[string]string{
			"bufio":                      "bufio",
			"bytes":                      "bytes",
			"errors":                     "errors",
			"fmt":                        "fmt",
			"golang.org/x/net/html/atom": "atom",
			"io":                         "io",
			"strconv":                    "strconv",
			"strings":                    "strings",
			"unicode/utf8":               "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Attribute":   reflect.TypeOf((*q.Attribute)(nil)).Elem(),
			"Node":        reflect.TypeOf((*q.Node)(nil)).Elem(),
			"NodeType":    reflect.TypeOf((*q.NodeType)(nil)).Elem(),
			"ParseOption": reflect.TypeOf((*q.ParseOption)(nil)).Elem(),
			"Token":       reflect.TypeOf((*q.Token)(nil)).Elem(),
			"TokenType":   reflect.TypeOf((*q.TokenType)(nil)).Elem(),
			"Tokenizer":   reflect.TypeOf((*q.Tokenizer)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrBufferExceeded": reflect.ValueOf(&q.ErrBufferExceeded),
		},
		Funcs: map[string]reflect.Value{
			"EscapeString":               reflect.ValueOf(q.EscapeString),
			"NewTokenizer":               reflect.ValueOf(q.NewTokenizer),
			"NewTokenizerFragment":       reflect.ValueOf(q.NewTokenizerFragment),
			"Parse":                      reflect.ValueOf(q.Parse),
			"ParseFragment":              reflect.ValueOf(q.ParseFragment),
			"ParseFragmentWithOptions":   reflect.ValueOf(q.ParseFragmentWithOptions),
			"ParseOptionEnableScripting": reflect.ValueOf(q.ParseOptionEnableScripting),
			"ParseWithOptions":           reflect.ValueOf(q.ParseWithOptions),
			"Render":                     reflect.ValueOf(q.Render),
			"UnescapeString":             reflect.ValueOf(q.UnescapeString),
		},
		TypedConsts: map[string]igop.TypedConst{
			"CommentNode":         {reflect.TypeOf(q.CommentNode), constant.MakeInt64(int64(q.CommentNode))},
			"CommentToken":        {reflect.TypeOf(q.CommentToken), constant.MakeInt64(int64(q.CommentToken))},
			"DoctypeNode":         {reflect.TypeOf(q.DoctypeNode), constant.MakeInt64(int64(q.DoctypeNode))},
			"DoctypeToken":        {reflect.TypeOf(q.DoctypeToken), constant.MakeInt64(int64(q.DoctypeToken))},
			"DocumentNode":        {reflect.TypeOf(q.DocumentNode), constant.MakeInt64(int64(q.DocumentNode))},
			"ElementNode":         {reflect.TypeOf(q.ElementNode), constant.MakeInt64(int64(q.ElementNode))},
			"EndTagToken":         {reflect.TypeOf(q.EndTagToken), constant.MakeInt64(int64(q.EndTagToken))},
			"ErrorNode":           {reflect.TypeOf(q.ErrorNode), constant.MakeInt64(int64(q.ErrorNode))},
			"ErrorToken":          {reflect.TypeOf(q.ErrorToken), constant.MakeInt64(int64(q.ErrorToken))},
			"RawNode":             {reflect.TypeOf(q.RawNode), constant.MakeInt64(int64(q.RawNode))},
			"SelfClosingTagToken": {reflect.TypeOf(q.SelfClosingTagToken), constant.MakeInt64(int64(q.SelfClosingTagToken))},
			"StartTagToken":       {reflect.TypeOf(q.StartTagToken), constant.MakeInt64(int64(q.StartTagToken))},
			"TextNode":            {reflect.TypeOf(q.TextNode), constant.MakeInt64(int64(q.TextNode))},
			"TextToken":           {reflect.TypeOf(q.TextToken), constant.MakeInt64(int64(q.TextToken))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
