package parsers

import "os"

var (
	_homePath, _  = os.UserHomeDir()
	_pathToConfig = _homePath + "/.config/mirage/nodes.toml"
)
