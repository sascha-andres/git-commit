package hook

import (
	"errors"
	"fmt"

	"livingit.de/code/git-commit/hook/v1"
)

type (
	// Validator defines which methods must be runnable
	Validator interface {
		Validate(commitMessage []string) (bool, error)
	}
)

// NewForVersion returns a hook implementation
func NewForVersion(version string) (Validator, error) {
	switch version {
	case "1":
		return v1.LoadConfig()
	case "":
		return v1.LoadConfig()
	}
	return nil, errors.New(fmt.Sprintf("no version %s known", version))
}
