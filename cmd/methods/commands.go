package methods

import (
	"fmt"
	"os"

	"livingit.de/code/git-commit/cmd/helper"
	"livingit.de/code/git-commit/subcommands"
)

// UninstallHook removes a git hook
func UninstallHook() int {
	err := subcommands.Uninstall(".git")
	if err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		return 1
	}
	return 0
}

// InstallHook adds hook to git repository
func InstallHook() int {
	if helper.DirectoryExists(".git") {
		err := subcommands.Install(".git")
		if err != nil {
			fmt.Fprintln(os.Stderr, "error:", err)
			return 1
		}
	} else {
		fmt.Fprintln(os.Stderr, "error: no git repository")
		return 1
	}
	return 0
}
