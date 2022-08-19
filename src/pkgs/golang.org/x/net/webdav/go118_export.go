// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package webdav

import (
	q "golang.org/x/net/webdav"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "webdav",
		Path: "golang.org/x/net/webdav",
		Deps: map[string]string{
			"bytes":                                "bytes",
			"container/heap":                       "heap",
			"context":                              "context",
			"encoding/xml":                         "xml",
			"errors":                               "errors",
			"fmt":                                  "fmt",
			"golang.org/x/net/webdav/internal/xml": "xml",
			"io":                                   "io",
			"mime":                                 "mime",
			"net/http":                             "http",
			"net/url":                              "url",
			"os":                                   "os",
			"path":                                 "path",
			"path/filepath":                        "filepath",
			"runtime":                              "runtime",
			"strconv":                              "strconv",
			"strings":                              "strings",
			"sync":                                 "sync",
			"time":                                 "time",
		},
		Interfaces: map[string]reflect.Type{
			"ContentTyper":    reflect.TypeOf((*q.ContentTyper)(nil)).Elem(),
			"DeadPropsHolder": reflect.TypeOf((*q.DeadPropsHolder)(nil)).Elem(),
			"ETager":          reflect.TypeOf((*q.ETager)(nil)).Elem(),
			"File":            reflect.TypeOf((*q.File)(nil)).Elem(),
			"FileSystem":      reflect.TypeOf((*q.FileSystem)(nil)).Elem(),
			"LockSystem":      reflect.TypeOf((*q.LockSystem)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Condition":   reflect.TypeOf((*q.Condition)(nil)).Elem(),
			"Dir":         reflect.TypeOf((*q.Dir)(nil)).Elem(),
			"Handler":     reflect.TypeOf((*q.Handler)(nil)).Elem(),
			"LockDetails": reflect.TypeOf((*q.LockDetails)(nil)).Elem(),
			"Property":    reflect.TypeOf((*q.Property)(nil)).Elem(),
			"Proppatch":   reflect.TypeOf((*q.Proppatch)(nil)).Elem(),
			"Propstat":    reflect.TypeOf((*q.Propstat)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrConfirmationFailed": reflect.ValueOf(&q.ErrConfirmationFailed),
			"ErrForbidden":          reflect.ValueOf(&q.ErrForbidden),
			"ErrLocked":             reflect.ValueOf(&q.ErrLocked),
			"ErrNoSuchLock":         reflect.ValueOf(&q.ErrNoSuchLock),
			"ErrNotImplemented":     reflect.ValueOf(&q.ErrNotImplemented),
		},
		Funcs: map[string]reflect.Value{
			"NewMemFS":   reflect.ValueOf(q.NewMemFS),
			"NewMemLS":   reflect.ValueOf(q.NewMemLS),
			"StatusText": reflect.ValueOf(q.StatusText),
		},
		TypedConsts: map[string]igop.TypedConst{},
		UntypedConsts: map[string]igop.UntypedConst{
			"StatusFailedDependency":    {"untyped int", constant.MakeInt64(int64(q.StatusFailedDependency))},
			"StatusInsufficientStorage": {"untyped int", constant.MakeInt64(int64(q.StatusInsufficientStorage))},
			"StatusLocked":              {"untyped int", constant.MakeInt64(int64(q.StatusLocked))},
			"StatusMulti":               {"untyped int", constant.MakeInt64(int64(q.StatusMulti))},
			"StatusUnprocessableEntity": {"untyped int", constant.MakeInt64(int64(q.StatusUnprocessableEntity))},
		},
	})
}
