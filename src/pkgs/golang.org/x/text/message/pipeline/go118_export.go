// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package pipeline

import (
	q "golang.org/x/text/message/pipeline"

	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "pipeline",
		Path: "golang.org/x/text/message/pipeline",
		Deps: map[string]string{
			"bytes":                               "bytes",
			"encoding/json":                       "json",
			"errors":                              "errors",
			"fmt":                                 "fmt",
			"go/ast":                              "ast",
			"go/build":                            "build",
			"go/constant":                         "constant",
			"go/format":                           "format",
			"go/parser":                           "parser",
			"go/token":                            "token",
			"go/types":                            "types",
			"golang.org/x/text/collate":           "collate",
			"golang.org/x/text/feature/plural":    "plural",
			"golang.org/x/text/internal":          "internal",
			"golang.org/x/text/internal/catmsg":   "catmsg",
			"golang.org/x/text/internal/format":   "format",
			"golang.org/x/text/internal/gen":      "gen",
			"golang.org/x/text/language":          "language",
			"golang.org/x/text/runes":             "runes",
			"golang.org/x/tools/go/callgraph":     "callgraph",
			"golang.org/x/tools/go/callgraph/cha": "cha",
			"golang.org/x/tools/go/loader":        "loader",
			"golang.org/x/tools/go/ssa":           "ssa",
			"golang.org/x/tools/go/ssa/ssautil":   "ssautil",
			"io":                                  "io",
			"io/ioutil":                           "ioutil",
			"log":                                 "log",
			"os":                                  "os",
			"path/filepath":                       "filepath",
			"regexp":                              "regexp",
			"sort":                                "sort",
			"strings":                             "strings",
			"text/template":                       "template",
			"unicode":                             "unicode",
			"unicode/utf8":                        "utf8",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Config":      reflect.TypeOf((*q.Config)(nil)).Elem(),
			"Feature":     reflect.TypeOf((*q.Feature)(nil)).Elem(),
			"IDList":      reflect.TypeOf((*q.IDList)(nil)).Elem(),
			"Message":     reflect.TypeOf((*q.Message)(nil)).Elem(),
			"Messages":    reflect.TypeOf((*q.Messages)(nil)).Elem(),
			"Placeholder": reflect.TypeOf((*q.Placeholder)(nil)).Elem(),
			"Select":      reflect.TypeOf((*q.Select)(nil)).Elem(),
			"State":       reflect.TypeOf((*q.State)(nil)).Elem(),
			"Text":        reflect.TypeOf((*q.Text)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Extract":  reflect.ValueOf(q.Extract),
			"Generate": reflect.ValueOf(q.Generate),
			"Rewrite":  reflect.ValueOf(q.Rewrite),
		},
		TypedConsts:   map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
