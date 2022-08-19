// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package xsrftoken

import (
	q "golang.org/x/net/xsrftoken"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "xsrftoken",
		Path: "golang.org/x/net/xsrftoken",
		Deps: map[string]string{
			"crypto/hmac":     "hmac",
			"crypto/sha1":     "sha1",
			"crypto/subtle":   "subtle",
			"encoding/base64": "base64",
			"fmt":             "fmt",
			"strconv":         "strconv",
			"strings":         "strings",
			"time":            "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"Generate": reflect.ValueOf(q.Generate),
			"Valid":    reflect.ValueOf(q.Valid),
			"ValidFor": reflect.ValueOf(q.ValidFor),
		},
		TypedConsts: map[string]igop.TypedConst{
			"Timeout": {reflect.TypeOf(q.Timeout), constant.MakeInt64(int64(q.Timeout))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
