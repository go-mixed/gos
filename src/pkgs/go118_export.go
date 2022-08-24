package pkgs

/*
qexp -outdir . -addtags "//+build go1.18" -filename go118_export github.com/inconshreveable/mousetrap
qexp -outdir . -addtags "//+build go1.18" -filename go118_export github.com/spf13/pflag
qexp -outdir . -addtags "//+build go1.18" -filename go118_export github.com/spf13/cobra
qexp -outdir . -addtags "//+build go1.18" -filename go118_export go.uber.org/multierr
*/

import (
	//_ "igop/src/pkgs/github.com/inconshreveable/mousetrap"
	_ "igop/src/pkgs/github.com/spf13/cobra"
	_ "igop/src/pkgs/github.com/spf13/pflag"
	_ "igop/src/pkgs/go.uber.org/multierr"
)
