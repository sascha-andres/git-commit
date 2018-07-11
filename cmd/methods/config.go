package methods

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"path"

	"github.com/imdario/mergo"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/pkg/errors"
	yaml "gopkg.in/yaml.v2"
	"livingit.de/code/git-commit/cmd/helper"
	"livingit.de/code/git-commit/hook"
)

const configFileName = ".commit-hook.yaml"

// LoadConfig handles global and local configuration
func LoadConfig() (*hook.Configuration, error) {
	config, err := loadGlobalConfig()
	if err != nil {
		return nil, err
	}

	commitMessageFile := os.Args[1]
	prjConfig, err := loadProjectConfiguration(commitMessageFile, config)
	if err != nil {
		return nil, err
	}

	if config == nil {
		config = prjConfig
	}

	if nil == config {
		return nil, errors.New("no suitable configuration found")
	}
	return config, nil
}

// loadProjectConfiguration loads a project specific configuration
func loadProjectConfiguration(commitMessageFile string, config *hook.Configuration) (*hook.Configuration, error) {
	projectPath, err := filepath.Abs(path.Join(filepath.Dir(commitMessageFile), ".."))
	if err != nil {
		return nil, err
	}

	localConfig := fmt.Sprintf("%s/%s", projectPath, configFileName)
	if helper.FileExists(localConfig) {
		data, err := ioutil.ReadFile(localConfig)
		if err != nil {
			return nil, err
		}
		var cfg hook.Configuration
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
		if nil == config {
			return &cfg, nil
		}
		if err := mergo.Merge(config, cfg); err != nil {
			return nil, err
		}
		return config, nil
	}
	return nil, nil
}

// loadGlobalConfig loads the global configuration if present
func loadGlobalConfig() (*hook.Configuration, error) {
	// Find home directory.
	home, err := homedir.Dir()
	if err != nil {
		return nil, err
	}

	globalConfig := fmt.Sprintf("%s/%s", home, configFileName)
	if helper.FileExists(globalConfig) {
		data, err := ioutil.ReadFile(globalConfig)
		if err != nil {
			return nil, err
		}
		var cfg hook.Configuration
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
		return &cfg, nil
	}

	return nil, nil
}
