// export by github.com/goplus/igop/cmd/qexp

//go:build go1.18
// +build go1.18

package font

import (
	q "golang.org/x/image/font"

	"go/constant"
	"reflect"

	"github.com/goplus/igop"
)

func init() {
	igop.RegisterPackage(&igop.Package{
		Name: "font",
		Path: "golang.org/x/image/font",
		Deps: map[string]string{
			"golang.org/x/image/math/fixed": "fixed",
			"image":                         "image",
			"image/draw":                    "draw",
			"io":                            "io",
			"unicode/utf8":                  "utf8",
		},
		Interfaces: map[string]reflect.Type{
			"Face": reflect.TypeOf((*q.Face)(nil)).Elem(),
		},
		NamedTypes: map[string]reflect.Type{
			"Drawer":  reflect.TypeOf((*q.Drawer)(nil)).Elem(),
			"Hinting": reflect.TypeOf((*q.Hinting)(nil)).Elem(),
			"Metrics": reflect.TypeOf((*q.Metrics)(nil)).Elem(),
			"Stretch": reflect.TypeOf((*q.Stretch)(nil)).Elem(),
			"Style":   reflect.TypeOf((*q.Style)(nil)).Elem(),
			"Weight":  reflect.TypeOf((*q.Weight)(nil)).Elem(),
		},
		AliasTypes: map[string]reflect.Type{},
		Vars:       map[string]reflect.Value{},
		Funcs: map[string]reflect.Value{
			"BoundBytes":    reflect.ValueOf(q.BoundBytes),
			"BoundString":   reflect.ValueOf(q.BoundString),
			"MeasureBytes":  reflect.ValueOf(q.MeasureBytes),
			"MeasureString": reflect.ValueOf(q.MeasureString),
		},
		TypedConsts: map[string]igop.TypedConst{
			"HintingFull":           {reflect.TypeOf(q.HintingFull), constant.MakeInt64(int64(q.HintingFull))},
			"HintingNone":           {reflect.TypeOf(q.HintingNone), constant.MakeInt64(int64(q.HintingNone))},
			"HintingVertical":       {reflect.TypeOf(q.HintingVertical), constant.MakeInt64(int64(q.HintingVertical))},
			"StretchCondensed":      {reflect.TypeOf(q.StretchCondensed), constant.MakeInt64(int64(q.StretchCondensed))},
			"StretchExpanded":       {reflect.TypeOf(q.StretchExpanded), constant.MakeInt64(int64(q.StretchExpanded))},
			"StretchExtraCondensed": {reflect.TypeOf(q.StretchExtraCondensed), constant.MakeInt64(int64(q.StretchExtraCondensed))},
			"StretchExtraExpanded":  {reflect.TypeOf(q.StretchExtraExpanded), constant.MakeInt64(int64(q.StretchExtraExpanded))},
			"StretchNormal":         {reflect.TypeOf(q.StretchNormal), constant.MakeInt64(int64(q.StretchNormal))},
			"StretchSemiCondensed":  {reflect.TypeOf(q.StretchSemiCondensed), constant.MakeInt64(int64(q.StretchSemiCondensed))},
			"StretchSemiExpanded":   {reflect.TypeOf(q.StretchSemiExpanded), constant.MakeInt64(int64(q.StretchSemiExpanded))},
			"StretchUltraCondensed": {reflect.TypeOf(q.StretchUltraCondensed), constant.MakeInt64(int64(q.StretchUltraCondensed))},
			"StretchUltraExpanded":  {reflect.TypeOf(q.StretchUltraExpanded), constant.MakeInt64(int64(q.StretchUltraExpanded))},
			"StyleItalic":           {reflect.TypeOf(q.StyleItalic), constant.MakeInt64(int64(q.StyleItalic))},
			"StyleNormal":           {reflect.TypeOf(q.StyleNormal), constant.MakeInt64(int64(q.StyleNormal))},
			"StyleOblique":          {reflect.TypeOf(q.StyleOblique), constant.MakeInt64(int64(q.StyleOblique))},
			"WeightBlack":           {reflect.TypeOf(q.WeightBlack), constant.MakeInt64(int64(q.WeightBlack))},
			"WeightBold":            {reflect.TypeOf(q.WeightBold), constant.MakeInt64(int64(q.WeightBold))},
			"WeightExtraBold":       {reflect.TypeOf(q.WeightExtraBold), constant.MakeInt64(int64(q.WeightExtraBold))},
			"WeightExtraLight":      {reflect.TypeOf(q.WeightExtraLight), constant.MakeInt64(int64(q.WeightExtraLight))},
			"WeightLight":           {reflect.TypeOf(q.WeightLight), constant.MakeInt64(int64(q.WeightLight))},
			"WeightMedium":          {reflect.TypeOf(q.WeightMedium), constant.MakeInt64(int64(q.WeightMedium))},
			"WeightNormal":          {reflect.TypeOf(q.WeightNormal), constant.MakeInt64(int64(q.WeightNormal))},
			"WeightSemiBold":        {reflect.TypeOf(q.WeightSemiBold), constant.MakeInt64(int64(q.WeightSemiBold))},
			"WeightThin":            {reflect.TypeOf(q.WeightThin), constant.MakeInt64(int64(q.WeightThin))},
		},
		UntypedConsts: map[string]igop.UntypedConst{},
	})
}
