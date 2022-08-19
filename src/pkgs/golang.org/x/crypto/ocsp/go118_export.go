// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package ocsp

import (
	q "golang.org/x/crypto/ocsp"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "ocsp",
		Path: "golang.org/x/crypto/ocsp",
		Deps: map[string]string{
			"crypto":           "crypto",
			"crypto/ecdsa":     "ecdsa",
			"crypto/elliptic":  "elliptic",
			"crypto/rand":      "rand",
			"crypto/rsa":       "rsa",
			"crypto/sha1":      "sha1",
			"crypto/sha256":    "sha256",
			"crypto/sha512":    "sha512",
			"crypto/x509":      "x509",
			"crypto/x509/pkix": "pkix",
			"encoding/asn1":    "asn1",
			"errors":           "errors",
			"fmt":              "fmt",
			"math/big":         "big",
			"strconv":          "strconv",
			"time":             "time",
		},
		Interfaces: map[string]reflect.Type{},
		NamedTypes: map[string]reflect.Type{
			"ParseError":     reflect.TypeOf((*q.ParseError)(nil)).Elem(),
			"Request":        reflect.TypeOf((*q.Request)(nil)).Elem(),
			"RequestOptions": reflect.TypeOf((*q.RequestOptions)(nil)).Elem(),
			"Response":       reflect.TypeOf((*q.Response)(nil)).Elem(),
			"ResponseError":  reflect.TypeOf((*q.ResponseError)(nil)).Elem(),
			"ResponseStatus": reflect.TypeOf((*q.ResponseStatus)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"InternalErrorErrorResponse":    reflect.ValueOf(&q.InternalErrorErrorResponse),
			"MalformedRequestErrorResponse": reflect.ValueOf(&q.MalformedRequestErrorResponse),
			"SigRequredErrorResponse":       reflect.ValueOf(&q.SigRequredErrorResponse),
			"TryLaterErrorResponse":         reflect.ValueOf(&q.TryLaterErrorResponse),
			"UnauthorizedErrorResponse":     reflect.ValueOf(&q.UnauthorizedErrorResponse),
		},
		Funcs: map[string]reflect.Value{
			"CreateRequest":        reflect.ValueOf(q.CreateRequest),
			"CreateResponse":       reflect.ValueOf(q.CreateResponse),
			"ParseRequest":         reflect.ValueOf(q.ParseRequest),
			"ParseResponse":        reflect.ValueOf(q.ParseResponse),
			"ParseResponseForCert": reflect.ValueOf(q.ParseResponseForCert),
		},
		TypedConsts: map[string]igop.TypedConst{
			"InternalError":     {reflect.TypeOf(q.InternalError), constant.MakeInt64(int64(q.InternalError))},
			"Malformed":         {reflect.TypeOf(q.Malformed), constant.MakeInt64(int64(q.Malformed))},
			"SignatureRequired": {reflect.TypeOf(q.SignatureRequired), constant.MakeInt64(int64(q.SignatureRequired))},
			"Success":           {reflect.TypeOf(q.Success), constant.MakeInt64(int64(q.Success))},
			"TryLater":          {reflect.TypeOf(q.TryLater), constant.MakeInt64(int64(q.TryLater))},
			"Unauthorized":      {reflect.TypeOf(q.Unauthorized), constant.MakeInt64(int64(q.Unauthorized))},
		},
		UntypedConsts: map[string]igop.UntypedConst{
			"AACompromise":         {"untyped int", constant.MakeInt64(int64(q.AACompromise))},
			"AffiliationChanged":   {"untyped int", constant.MakeInt64(int64(q.AffiliationChanged))},
			"CACompromise":         {"untyped int", constant.MakeInt64(int64(q.CACompromise))},
			"CertificateHold":      {"untyped int", constant.MakeInt64(int64(q.CertificateHold))},
			"CessationOfOperation": {"untyped int", constant.MakeInt64(int64(q.CessationOfOperation))},
			"Good":                 {"untyped int", constant.MakeInt64(int64(q.Good))},
			"KeyCompromise":        {"untyped int", constant.MakeInt64(int64(q.KeyCompromise))},
			"PrivilegeWithdrawn":   {"untyped int", constant.MakeInt64(int64(q.PrivilegeWithdrawn))},
			"RemoveFromCRL":        {"untyped int", constant.MakeInt64(int64(q.RemoveFromCRL))},
			"Revoked":              {"untyped int", constant.MakeInt64(int64(q.Revoked))},
			"ServerFailed":         {"untyped int", constant.MakeInt64(int64(q.ServerFailed))},
			"Superseded":           {"untyped int", constant.MakeInt64(int64(q.Superseded))},
			"Unknown":              {"untyped int", constant.MakeInt64(int64(q.Unknown))},
			"Unspecified":          {"untyped int", constant.MakeInt64(int64(q.Unspecified))},
		},
	})
}
