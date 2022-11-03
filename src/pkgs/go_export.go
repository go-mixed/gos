package pkgs

/*
qexp -outdir . -filename go_export github.com/inconshreveable/mousetrap github.com/spf13/pflag github.com/spf13/cobra go.uber.org/multierr gopkg.in/yaml.v3

*/

import (
	_ "igop/src/pkgs/github.com/inconshreveable/mousetrap"
	_ "igop/src/pkgs/github.com/spf13/cobra"
	_ "igop/src/pkgs/github.com/spf13/pflag"
	_ "igop/src/pkgs/go.uber.org/multierr"
	_ "igop/src/pkgs/gopkg.in/yaml.v3"
)
