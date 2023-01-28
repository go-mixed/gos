package pkgs

/*
qexp -outdir . -filename go_export github.com/inconshreveable/mousetrap github.com/spf13/pflag github.com/spf13/cobra go.uber.org/multierr gopkg.in/yaml.v3 github.com/pkg/errors

*/

import (
	_ "gopkg.in/go-mixed/igop.v1/pkgs/github.com/inconshreveable/mousetrap"
	_ "gopkg.in/go-mixed/igop.v1/pkgs/github.com/pkg/errors"
	_ "gopkg.in/go-mixed/igop.v1/pkgs/github.com/spf13/cobra"
	_ "gopkg.in/go-mixed/igop.v1/pkgs/github.com/spf13/pflag"
	_ "gopkg.in/go-mixed/igop.v1/pkgs/go.uber.org/multierr"
	_ "gopkg.in/go-mixed/igop.v1/pkgs/gopkg.in/yaml.v3"
)
