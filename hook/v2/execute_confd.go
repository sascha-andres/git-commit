package v2

import (
	"os"
	"path/filepath"

	"livingit.de/code/gitconfd"
)

// ExecuteConfDDirectories run scrips in .commit-msg.d/ directory
func (cfg *Configuration) ExecuteConfDDirectories() bool {
	if !cfg.ConfDDirectoriesEnabled {
		return true
	}

	path := filepath.Dir(os.Args[1])

	return gitconfd.Execute(path, os.Args[0])
}
