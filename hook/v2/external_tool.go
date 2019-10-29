package v2

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"syscall"
)

// runExternalTools iterates over external tools to validate commit
func (cfg *Configuration) runExternalTools() bool {
	result := true
	for _, val := range cfg.Externals {
		result = val.execute() && result
	}
	return result
}

// execute runs command
func (val *Tool) execute() bool {
	cmd := exec.Command(val.Command[0], val.Command[1:]...)

	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Start(); err != nil {
		fmt.Println(fmt.Sprintf("error: unable to start [%s]: %s", val.Name, err.Error()))
		return false
	}
	if err := cmd.Wait(); err != nil {
		return val.handleExecuteError(stdout, stderr, err)
	}
	return true
}

// handleExecuteError does the lifting for an execute error
func (val *Tool) handleExecuteError(stdout, stderr bytes.Buffer, err error) bool {
	var (
		exitError *exec.ExitError
		ok        bool
	)

	if exitError, ok = err.(*exec.ExitError); !ok {
		fmt.Println(fmt.Sprintf("error: unable get exit code [%s]: %s", val.Name, err.Error()))
	}
	printOutput(stdout, stderr)

	if _, ok := exitError.Sys().(syscall.WaitStatus); ok {
		fmt.Println(fmt.Sprintf("%s: execution of [%s] failed", val.Severity, val.Name))
		if val.Severity != ErrorSeverity {
			return true
		}
	}
	return false
}

// printOutput dumps command output
func printOutput(stdout, stderr bytes.Buffer) {
	fmt.Println("output:")
	fmt.Println()
	fmt.Println(stdout.String())

	fmt.Println("error:")
	fmt.Println()
	_, _ = fmt.Fprintln(os.Stderr, stderr.String())
}
