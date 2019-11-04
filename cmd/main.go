package main

import (
	"fmt"
	"os"

	"bufio"

	"github.com/pkg/errors"
	"livingit.de/code/git-commit/cmd/helper"
	"livingit.de/code/git-commit/cmd/methods"
	"livingit.de/code/git-commit/hook"
)

func main() {
	methods.PrintVersion()

	if prepare() {
		return
	}

	ok := validate()

	if !ok {
		os.Exit(1)
		return
	}

	os.Exit(0)
}

// prepare checks input and runs actions for install/uninstall when requested
func prepare() bool {
	if len(os.Args) < 2 {
		methods.Help()
		os.Exit(0)
		return true
	}

	var (
		fn  func() int
		run bool
	)

	if os.Args[1] == "install" {
		fn = methods.InstallHook
		run = true
	}

	if os.Args[1] == "uninstall" {
		fn = methods.UninstallHook
		run = true
	}

	if run {
		os.Exit(fn())
		return true
	}

	validationResult := validateInput()
	if 0 != validationResult {
		os.Exit(validationResult)
		return true
	}

	return false
}

// validate runs validation
func validate() bool {
	commitMessageFile := os.Args[1]
	config, err := hook.NewForVersion(commitMessageFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return false
	}

	commitFileContent, err := loadCommitMessageFile(commitMessageFile)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return false
	}

	ok, err := config.Validate(commitFileContent)
	if err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
		return false
	}

	return ok
}

// validateInput checks program oarameters when running
// as a hook
func validateInput() int {
	commitMessageFile := os.Args[1]
	if commitMessageFile == "" {
		_, _ = fmt.Fprintln(os.Stderr, errors.New("no commit message file passed as parameter 1"))
		return 1
	}

	if !helper.FileExists(commitMessageFile) {
		_, _ = fmt.Fprintln(os.Stderr, errors.New("passed commit message file not found"))
		return 1
	}
	return 0
}

// loadCommitMessageFile reads the commit message and returns it as a
// array containing the lines
func loadCommitMessageFile(commitMessageFile string) ([]string, error) {
	file, err := os.Open(commitMessageFile)
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file)
	commitFileContent := make([]string, 0)
	for scanner.Scan() {
		commitFileContent = append(commitFileContent, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return commitFileContent, nil
}
