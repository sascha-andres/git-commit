package subcommands

import (
	"os"

	"flag"

	"github.com/pkg/errors"
)

// Install activates hook for project
func Install(gitFolderPath string) error {
	var forceOverwrite bool
	flagSet := flag.NewFlagSet("git-commit-hook install", flag.ContinueOnError)
	flagSet.BoolVar(&forceOverwrite, "f", false, `force file creation by overwriting`)
	err := flagSet.Parse(os.Args[2:])
	if err != nil {
		return errors.New("unable to parse arguments")
	}

	exeFile, err := os.Executable()
	if err != nil {
		return err
	}

	commitHookFilePath := createCommitHookFilePath(gitFolderPath)

	if checkFileExists(commitHookFilePath) {
		if !forceOverwrite {
			return errors.New("file already exists, use -f to force overwriting")
		}

		err = os.Remove(commitHookFilePath)
		if err != nil {
			return err
		}
	}

	return os.Symlink(exeFile, commitHookFilePath)
}

func checkFileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	return err == nil
}
