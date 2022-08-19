// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package cldr

import (
	q "golang.org/x/text/unicode/cldr"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "cldr",
		Path: "golang.org/x/text/unicode/cldr",
		Deps: map[string]string{
			"archive/zip":   "zip",
			"bufio":         "bufio",
			"bytes":         "bytes",
			"encoding/xml":  "xml",
			"errors":        "errors",
			"fmt":           "fmt",
			"io":            "io",
			"io/ioutil":     "ioutil",
			"log":           "log",
			"os":            "os",
			"path/filepath": "filepath",
			"reflect":       "reflect",
			"regexp":        "regexp",
			"sort":          "sort",
			"strconv":       "strconv",
			"strings":       "strings",
			"unicode":       "unicode",
			"unicode/utf8":  "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Elem":          reflect.TypeOf((*q.Elem)(nil)).Elem(),
			"Loader":        reflect.TypeOf((*q.Loader)(nil)).Elem(),
			"RuleProcessor": reflect.TypeOf((*q.RuleProcessor)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"CLDR":               reflect.TypeOf((*q.CLDR)(nil)).Elem(),
			"Calendar":           reflect.TypeOf((*q.Calendar)(nil)).Elem(),
			"Collation":          reflect.TypeOf((*q.Collation)(nil)).Elem(),
			"Common":             reflect.TypeOf((*q.Common)(nil)).Elem(),
			"Decoder":            reflect.TypeOf((*q.Decoder)(nil)).Elem(),
			"Draft":              reflect.TypeOf((*q.Draft)(nil)).Elem(),
			"LDML":               reflect.TypeOf((*q.LDML)(nil)).Elem(),
			"LDMLBCP47":          reflect.TypeOf((*q.LDMLBCP47)(nil)).Elem(),
			"LocaleDisplayNames": reflect.TypeOf((*q.LocaleDisplayNames)(nil)).Elem(),
			"Numbers":            reflect.TypeOf((*q.Numbers)(nil)).Elem(),
			"Slice":              reflect.TypeOf((*q.Slice)(nil)).Elem(),
			"SupplementalData":   reflect.TypeOf((*q.SupplementalData)(nil)).Elem(),
			"TimeZoneNames":      reflect.TypeOf((*q.TimeZoneNames)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Get":        reflect.ValueOf(q.Get),
			"Key":        reflect.ValueOf(q.Key),
			"MakeSlice":  reflect.ValueOf(q.MakeSlice),
			"ParseDraft": reflect.ValueOf(q.ParseDraft),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Approved":    {reflect.TypeOf(q.Approved), constant.MakeInt64(int64(q.Approved))},
			"Contributed": {reflect.TypeOf(q.Contributed), constant.MakeInt64(int64(q.Contributed))},
			"Provisional": {reflect.TypeOf(q.Provisional), constant.MakeInt64(int64(q.Provisional))},
			"Unconfirmed": {reflect.TypeOf(q.Unconfirmed), constant.MakeInt64(int64(q.Unconfirmed))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"Version": {"untyped string", constant.MakeString(string(q.Version))},
		},
	})
}
