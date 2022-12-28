package pkgs

/*
qexp -outdir . -filename go_export github.com/inconshreveable/mousetrap github.com/spf13/pflag github.com/spf13/cobra go.uber.org/multierr gopkg.in/yaml.v3 github.com/pkg/errors

*/

import (
	_ "github.com/fly-studio/igop/pkgs/github.com/inconshreveable/mousetrap"
	_ "github.com/fly-studio/igop/pkgs/github.com/pkg/errors"
	_ "github.com/fly-studio/igop/pkgs/github.com/spf13/cobra"
	_ "github.com/fly-studio/igop/pkgs/github.com/spf13/pflag"
	_ "github.com/fly-studio/igop/pkgs/go.uber.org/multierr"
	_ "github.com/fly-studio/igop/pkgs/gopkg.in/yaml.v3"
)
