// export by github.com/goplus/xgo/cmd/qexp

package yaml

import (
	"github.com/goplus/ixgo"
	q "gopkg.in/yaml.v3"

	"go/constant"
	"reflect"
)

func init() {
	ixgo.RegisterPackage(&ixgo.Package{
		Name: "yaml",
		Path: "gopkg.in/yaml.v3",
		Deps: map[string]string{
			"bytes":           "bytes",
			"encoding":        "encoding",
			"encoding/base64": "base64",
			"errors":          "errors",
			"fmt":             "fmt",
			"io":              "io",
			"math":            "math",
			"reflect":         "reflect",
			"regexp":          "regexp",
			"sort":            "sort",
			"strconv":         "strconv",
			"strings":         "strings",
			"sync":            "sync",
			"time":            "time",
			"unicode":         "unicode",
			"unicode/utf8":    "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"IsZeroer":    reflect.TypeOf((*q.IsZeroer)(nil)).Elem(),
			"Marshaler":   reflect.TypeOf((*q.Marshaler)(nil)).Elem(),
			"Unmarshaler": reflect.TypeOf((*q.Unmarshaler)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Decoder":   reflect.TypeOf((*q.Decoder)(nil)).Elem(),
			"Encoder":   reflect.TypeOf((*q.Encoder)(nil)).Elem(),
			"Kind":      reflect.TypeOf((*q.Kind)(nil)).Elem(),
			"Node":      reflect.TypeOf((*q.Node)(nil)).Elem(),
			"Style":     reflect.TypeOf((*q.Style)(nil)).Elem(),
			"TypeError": reflect.TypeOf((*q.TypeError)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Marshal":    reflect.ValueOf(q.Marshal),
			"NewDecoder": reflect.ValueOf(q.NewDecoder),
			"NewEncoder": reflect.ValueOf(q.NewEncoder),
			"Unmarshal":  reflect.ValueOf(q.Unmarshal),
		},
		TypedConsts: map[string]ixgo.TypedConst{
			"AliasNode":         {reflect.TypeOf(q.AliasNode), constant.MakeInt64(int64(q.AliasNode))},
			"DocumentNode":      {reflect.TypeOf(q.DocumentNode), constant.MakeInt64(int64(q.DocumentNode))},
			"DoubleQuotedStyle": {reflect.TypeOf(q.DoubleQuotedStyle), constant.MakeInt64(int64(q.DoubleQuotedStyle))},
			"FlowStyle":         {reflect.TypeOf(q.FlowStyle), constant.MakeInt64(int64(q.FlowStyle))},
			"FoldedStyle":       {reflect.TypeOf(q.FoldedStyle), constant.MakeInt64(int64(q.FoldedStyle))},
			"LiteralStyle":      {reflect.TypeOf(q.LiteralStyle), constant.MakeInt64(int64(q.LiteralStyle))},
			"MappingNode":       {reflect.TypeOf(q.MappingNode), constant.MakeInt64(int64(q.MappingNode))},
			"ScalarNode":        {reflect.TypeOf(q.ScalarNode), constant.MakeInt64(int64(q.ScalarNode))},
			"SequenceNode":      {reflect.TypeOf(q.SequenceNode), constant.MakeInt64(int64(q.SequenceNode))},
			"SingleQuotedStyle": {reflect.TypeOf(q.SingleQuotedStyle), constant.MakeInt64(int64(q.SingleQuotedStyle))},
			"TaggedStyle":       {reflect.TypeOf(q.TaggedStyle), constant.MakeInt64(int64(q.TaggedStyle))},
		},
		UntypedConsts: map[string]ixgo.UntypedConst{},
	})
}
