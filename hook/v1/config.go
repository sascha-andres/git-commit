package v1

import (
	"fmt"
	"os"
	"path/filepath"

	"path"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"livingit.de/code/git-commit/hook/config"
)

const configFileName = ".commit-hook.yaml"

// LoadConfig handles global and local configuration
func LoadConfig() (*Configuration, error) {
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
func loadProjectConfiguration(commitMessageFile string, globalConfig *Configuration) (*Configuration, error) {
	projectPath, err := filepath.Abs(path.Join(filepath.Dir(commitMessageFile), ".."))
	if err != nil {
		return nil, err
	}

	localConfig := fmt.Sprintf("%s/%s", projectPath, configFileName)
	data, err := config.LoadProjectConfigFileContent(localConfig)
	if err != nil {
		return nil, err
	}
	if data != nil {
		var cfg Configuration
		err = yaml.Unmarshal(data, &cfg)
		if err != nil {
			return nil, err
		}
		if nil == globalConfig {
			return &cfg, nil
		}
		if err := mergo.Merge(globalConfig, cfg); err != nil {
			return nil, err
		}
	}
	return globalConfig, nil
}

// loadGlobalConfig loads the global configuration if present
func loadGlobalConfig() (*Configuration, error) {
	data, err := config.LoadGlobalConfigFileContent()
	if err != nil {
		return nil, err
	}
	var cfg Configuration
	err = yaml.Unmarshal(data, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
