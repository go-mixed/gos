//go:build !windows
// +build !windows

package mod

import (
	"github.com/pkg/errors"
	"plugin"
)

func (ctx *Context) LoadPlugins(plugins []string) error {
	for _, path := range plugins {
		p, err := plugin.Open(path)
		if err != nil {
			return err
		}
		loadFunc, err := p.Lookup("Load")
		if err != nil {
			return errors.Wrapf(err, "func \"Load\" not found in plugin \"%s\"", path)
		}
		loadFunc.(func())()
	}
	return nil
}
