package cmd

type CmdOptions struct {
	Debug       bool
	Script      string
	ScriptIsSet bool

	Path        string
	VendorPath  string
	ImportPaths map[string]string
	PluginPaths []string

	realPath  string
	isDir     bool
	isArchive bool
}
