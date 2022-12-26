package mod

import "fmt"

func (ctx *Context) LoadPlugins(plugins []string) error {
	if len(plugins) > 0 {
		fmt.Println("WARNING: Loading plugin in Windows is not supported.")
	}
	return nil
}
