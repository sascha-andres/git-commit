package v2

import (
	"fmt"
	"os"
	"os/exec"
)

// validateGitLabCI runs the gitlab ci linter if configured
func (cfg *Configuration) validateGitLabCI() (result bool) {
	if !cfg.LintGitLabCI {
		return true
	}

	var (
		cmd *exec.Cmd
	)

	fmt.Println("running GitLab CI linter")

	pathToLinter, err := exec.LookPath("gitlab-ci-linter")
	if err != nil {
		return
	}

	if cfg.GitLabCIFile != "" {
		cmd = exec.Command(pathToLinter, "--ci-file", cfg.GitLabCIFile)
	} else {
		cmd = exec.Command(pathToLinter)
	}
	cmd.Stdout = os.Stdin
	cmd.Stderr = os.Stderr
	err = cmd.Start()
	if err != nil {
		return
	}
	err = cmd.Wait()
	if err != nil {
		return
	}
	exitCode := cmd.ProcessState.ExitCode()
	if exitCode != 0 {
		return
	}

	return true
}
