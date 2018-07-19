package hook

import (
	"errors"
	"fmt"

	"livingit.de/code/git-commit/hook/config"
	"livingit.de/code/git-commit/hook/v2"
	"livingit.de/code/versioned"
)

type (
	// Validator defines which methods must be runnable
	Validator interface {
		Validate(commitMessage []string) (bool, error)
	}
)

// NewForVersion returns a hook implementation
func NewForVersion(commitMessageFile string) (Validator, error) {
	version, err := getVersion(commitMessageFile)
	if err != nil {
		return nil, err
	}
	switch version {
	case "1":
		return nil, errors.New("v1 is removed from utility")
	case "2":
		return v2.LoadConfig()
	case "":
		return v2.LoadConfig()
	}
	return nil, fmt.Errorf("no version %s known", version)
}

// getVersion reads the version of global and local configuration
// and returns the version after some validations
func getVersion(commitMessageFile string) (string, error) {
	global, globalVersion, err := getGlobalVersion()

	if err != nil {
		return "", err
	}

	local, localVersion, err := getLocalVersion(commitMessageFile)
	if err != nil {
		return "", err
	}

	if nil != global {
		if nil != local {
			if localVersion == globalVersion && globalVersion == "" {
				return "", errors.New("you have to provide versions for global and local config")
			}
			if localVersion != globalVersion {
				return "", errors.New("version mismatch for global and project version")
			}
		}
		return globalVersion, nil
	}

	if nil != local {
		return localVersion, nil
	}

	return "", errors.New("no suitable versioned configuration found")
}

// getLocalVersion returns version data from project configuration
// file
func getLocalVersion(commitMessageFile string) ([]byte, string, error) {
	v := versioned.NewVersionReader()
	local, err := config.LoadProjectConfigFileContent(commitMessageFile)
	if err != nil {
		return nil, "", err
	}
	if nil != local {
		localVersion, err := v.YAML.GetVersion(local)
		if err != nil {
			return local, "", err
		}
		return local, localVersion, nil
	}
	return nil, "", nil
}

// getGlobalVersion returns version data from global configuration
// file
func getGlobalVersion() ([]byte, string, error) {
	v := versioned.NewVersionReader()
	global, err := config.LoadGlobalConfigFileContent()
	if err != nil {
		return nil, "", err
	}
	if nil != global {
		version, err := v.YAML.GetVersion(global)
		return global, version, err
	}
	return nil, "", nil
}
