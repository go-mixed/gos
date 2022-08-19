// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package dnsmessage

import (
	q "golang.org/x/net/dns/dnsmessage"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "dnsmessage",
		Path: "golang.org/x/net/dns/dnsmessage",
		Deps: map[string]string{
			"errors": "errors",
		},
		Interfaces: map[string]reflect.Type{
			"ResourceBody": reflect.TypeOf((*q.ResourceBody)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"AAAAResource":    reflect.TypeOf((*q.AAAAResource)(nil)).Elem(),
			"AResource":       reflect.TypeOf((*q.AResource)(nil)).Elem(),
			"Builder":         reflect.TypeOf((*q.Builder)(nil)).Elem(),
			"CNAMEResource":   reflect.TypeOf((*q.CNAMEResource)(nil)).Elem(),
			"Class":           reflect.TypeOf((*q.Class)(nil)).Elem(),
			"Header":          reflect.TypeOf((*q.Header)(nil)).Elem(),
			"MXResource":      reflect.TypeOf((*q.MXResource)(nil)).Elem(),
			"Message":         reflect.TypeOf((*q.Message)(nil)).Elem(),
			"NSResource":      reflect.TypeOf((*q.NSResource)(nil)).Elem(),
			"Name":            reflect.TypeOf((*q.Name)(nil)).Elem(),
			"OPTResource":     reflect.TypeOf((*q.OPTResource)(nil)).Elem(),
			"OpCode":          reflect.TypeOf((*q.OpCode)(nil)).Elem(),
			"Option":          reflect.TypeOf((*q.Option)(nil)).Elem(),
			"PTRResource":     reflect.TypeOf((*q.PTRResource)(nil)).Elem(),
			"Parser":          reflect.TypeOf((*q.Parser)(nil)).Elem(),
			"Question":        reflect.TypeOf((*q.Question)(nil)).Elem(),
			"RCode":           reflect.TypeOf((*q.RCode)(nil)).Elem(),
			"Resource":        reflect.TypeOf((*q.Resource)(nil)).Elem(),
			"ResourceHeader":  reflect.TypeOf((*q.ResourceHeader)(nil)).Elem(),
			"SOAResource":     reflect.TypeOf((*q.SOAResource)(nil)).Elem(),
			"SRVResource":     reflect.TypeOf((*q.SRVResource)(nil)).Elem(),
			"TXTResource":     reflect.TypeOf((*q.TXTResource)(nil)).Elem(),
			"Type":            reflect.TypeOf((*q.Type)(nil)).Elem(),
			"UnknownResource": reflect.TypeOf((*q.UnknownResource)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars: map[string]reflect.Value{
			"ErrNotStarted":  reflect.ValueOf(&q.ErrNotStarted),
			"ErrSectionDone": reflect.ValueOf(&q.ErrSectionDone),
		},
		Funcs: map[string]reflect.Value{
			"MustNewName": reflect.ValueOf(q.MustNewName),
			"NewBuilder":  reflect.ValueOf(q.NewBuilder),
			"NewName":     reflect.ValueOf(q.NewName),
		},
		TypedConsts: map[string]igop.TypedConst{
			"ClassANY":            {reflect.TypeOf(q.ClassANY), constant.MakeInt64(int64(q.ClassANY))},
			"ClassCHAOS":          {reflect.TypeOf(q.ClassCHAOS), constant.MakeInt64(int64(q.ClassCHAOS))},
			"ClassCSNET":          {reflect.TypeOf(q.ClassCSNET), constant.MakeInt64(int64(q.ClassCSNET))},
			"ClassHESIOD":         {reflect.TypeOf(q.ClassHESIOD), constant.MakeInt64(int64(q.ClassHESIOD))},
			"ClassINET":           {reflect.TypeOf(q.ClassINET), constant.MakeInt64(int64(q.ClassINET))},
			"RCodeFormatError":    {reflect.TypeOf(q.RCodeFormatError), constant.MakeInt64(int64(q.RCodeFormatError))},
			"RCodeNameError":      {reflect.TypeOf(q.RCodeNameError), constant.MakeInt64(int64(q.RCodeNameError))},
			"RCodeNotImplemented": {reflect.TypeOf(q.RCodeNotImplemented), constant.MakeInt64(int64(q.RCodeNotImplemented))},
			"RCodeRefused":        {reflect.TypeOf(q.RCodeRefused), constant.MakeInt64(int64(q.RCodeRefused))},
			"RCodeServerFailure":  {reflect.TypeOf(q.RCodeServerFailure), constant.MakeInt64(int64(q.RCodeServerFailure))},
			"RCodeSuccess":        {reflect.TypeOf(q.RCodeSuccess), constant.MakeInt64(int64(q.RCodeSuccess))},
			"TypeA":               {reflect.TypeOf(q.TypeA), constant.MakeInt64(int64(q.TypeA))},
			"TypeAAAA":            {reflect.TypeOf(q.TypeAAAA), constant.MakeInt64(int64(q.TypeAAAA))},
			"TypeALL":             {reflect.TypeOf(q.TypeALL), constant.MakeInt64(int64(q.TypeALL))},
			"TypeAXFR":            {reflect.TypeOf(q.TypeAXFR), constant.MakeInt64(int64(q.TypeAXFR))},
			"TypeCNAME":           {reflect.TypeOf(q.TypeCNAME), constant.MakeInt64(int64(q.TypeCNAME))},
			"TypeHINFO":           {reflect.TypeOf(q.TypeHINFO), constant.MakeInt64(int64(q.TypeHINFO))},
			"TypeMINFO":           {reflect.TypeOf(q.TypeMINFO), constant.MakeInt64(int64(q.TypeMINFO))},
			"TypeMX":              {reflect.TypeOf(q.TypeMX), constant.MakeInt64(int64(q.TypeMX))},
			"TypeNS":              {reflect.TypeOf(q.TypeNS), constant.MakeInt64(int64(q.TypeNS))},
			"TypeOPT":             {reflect.TypeOf(q.TypeOPT), constant.MakeInt64(int64(q.TypeOPT))},
			"TypePTR":             {reflect.TypeOf(q.TypePTR), constant.MakeInt64(int64(q.TypePTR))},
			"TypeSOA":             {reflect.TypeOf(q.TypeSOA), constant.MakeInt64(int64(q.TypeSOA))},
			"TypeSRV":             {reflect.TypeOf(q.TypeSRV), constant.MakeInt64(int64(q.TypeSRV))},
			"TypeTXT":             {reflect.TypeOf(q.TypeTXT), constant.MakeInt64(int64(q.TypeTXT))},
			"TypeWKS":             {reflect.TypeOf(q.TypeWKS), constant.MakeInt64(int64(q.TypeWKS))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
