package config

import (
	"fmt"
	"io/ioutil"

	"github.com/mitchellh/go-homedir"
	"livingit.de/code/git-commit/cmd/helper"
)

const configFileName = ".commit-hook.yaml"

// LoadGlobalConfigFileContent reads data from the global configuration file
func LoadGlobalConfigFileContent() ([]byte, error) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}
	return loadConfigFileContent(fmt.Sprintf("%s/%s", home, configFileName))
}

// LoadProjectConfigFileContent reads data from the project configuration file
func LoadProjectConfigFileContent(commitMessageFile string) ([]byte, error) {
	return loadConfigFileContent(commitMessageFile)
}

// loadConfigFileContent returns the file content if the file exists
func loadConfigFileContent(file string) ([]byte, error) {
	if helper.FileExists(file) {
		return ioutil.ReadFile(file)
	}
	return nil, nil
}
