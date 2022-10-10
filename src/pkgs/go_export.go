package pkgs

/*
qexp -outdir . -addtags "//+build go1.19" -filename go119_export github.com/inconshreveable/mousetrap github.com/spf13/pflag github.com/spf13/cobra go.uber.org/multierr
*/

import (
	_ "igop/src/pkgs/github.com/inconshreveable/mousetrap"
	_ "igop/src/pkgs/github.com/spf13/cobra"
	_ "igop/src/pkgs/github.com/spf13/pflag"
	_ "igop/src/pkgs/go.uber.org/multierr"
)
