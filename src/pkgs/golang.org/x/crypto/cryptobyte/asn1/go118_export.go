// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package asn1

import (
	q "golang.org/x/crypto/cryptobyte/asn1"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name:       "asn1",
		Path:       "golang.org/x/crypto/cryptobyte/asn1",
		Deps:       map[string]string{},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"Tag": reflect.TypeOf((*q.Tag)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs:      map[string]reflect.Value{},
		TypedConsts: map[string]igop.TypedConst{
			"BIT_STRING":        {reflect.TypeOf(q.BIT_STRING), constant.MakeInt64(int64(q.BIT_STRING))},
			"BOOLEAN":           {reflect.TypeOf(q.BOOLEAN), constant.MakeInt64(int64(q.BOOLEAN))},
			"ENUM":              {reflect.TypeOf(q.ENUM), constant.MakeInt64(int64(q.ENUM))},
			"GeneralString":     {reflect.TypeOf(q.GeneralString), constant.MakeInt64(int64(q.GeneralString))},
			"GeneralizedTime":   {reflect.TypeOf(q.GeneralizedTime), constant.MakeInt64(int64(q.GeneralizedTime))},
			"IA5String":         {reflect.TypeOf(q.IA5String), constant.MakeInt64(int64(q.IA5String))},
			"INTEGER":           {reflect.TypeOf(q.INTEGER), constant.MakeInt64(int64(q.INTEGER))},
			"NULL":              {reflect.TypeOf(q.NULL), constant.MakeInt64(int64(q.NULL))},
			"OBJECT_IDENTIFIER": {reflect.TypeOf(q.OBJECT_IDENTIFIER), constant.MakeInt64(int64(q.OBJECT_IDENTIFIER))},
			"OCTET_STRING":      {reflect.TypeOf(q.OCTET_STRING), constant.MakeInt64(int64(q.OCTET_STRING))},
			"PrintableString":   {reflect.TypeOf(q.PrintableString), constant.MakeInt64(int64(q.PrintableString))},
			"SEQUENCE":          {reflect.TypeOf(q.SEQUENCE), constant.MakeInt64(int64(q.SEQUENCE))},
			"SET":               {reflect.TypeOf(q.SET), constant.MakeInt64(int64(q.SET))},
			"T61String":         {reflect.TypeOf(q.T61String), constant.MakeInt64(int64(q.T61String))},
			"UTCTime":           {reflect.TypeOf(q.UTCTime), constant.MakeInt64(int64(q.UTCTime))},
			"UTF8String":        {reflect.TypeOf(q.UTF8String), constant.MakeInt64(int64(q.UTF8String))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
