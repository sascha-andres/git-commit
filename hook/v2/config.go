package v2

import (
	"os"

	"github.com/imdario/mergo"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"livingit.de/code/git-commit/hook/config"
)

// LoadConfig handles global and local configuration
func LoadConfig() (*Configuration, error) {
	cfg, err := loadGlobalConfig()
	if err != nil {
		return nil, err
	}

	commitMessageFile := os.Args[1]
	prjConfig, err := loadProjectConfiguration(commitMessageFile, cfg)
	if err != nil {
		return nil, err
	}

	if cfg == nil {
		cfg = prjConfig
	}

	if nil == cfg {
		return nil, errors.New("no suitable configuration found")
	}
	return cfg, nil
}

// loadProjectConfiguration loads a project specific configuration
func loadProjectConfiguration(commitMessageFile string, globalConfig *Configuration) (*Configuration, error) {
	data, err := config.LoadProjectConfigFileContent(commitMessageFile)
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
		if err := mergo.Merge(globalConfig, cfg, mergo.WithAppendSlice); err != nil {
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
