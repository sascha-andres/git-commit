package subcommands

import (
	"os"
)

// Uninstall removes a hook
func Uninstall(gitFolderPath string) error {
	commitHookFilePath := createCommitHookFilePath(gitFolderPath)

	return os.Remove(commitHookFilePath)
}
