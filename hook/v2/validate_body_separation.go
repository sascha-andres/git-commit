package v2

import "fmt"

// validateBodySeparation runs body related rules
func (cfg *Configuration) validateBodySeparation(commitMessage []string) bool {
	result := true
	if cfg.BodyRequired {
		if len(commitMessage) == 1 {
			fmt.Println("error: body required but not provided")
			result = false
		}
	}
	result = cfg.validateBodySeparationOnSeparateBody(commitMessage) && result
	return result
}

// validateBodySeparationOnSeparateBody runs rules if separate-body is set
func (cfg *Configuration) validateBodySeparationOnSeparateBody(commitMessage []string) bool {
	result := true
	if cfg.SeparateBody && len(commitMessage) > 1 {
		if len(commitMessage) == 2 {
			if cfg.SeparateBody {
				fmt.Println("error: body should be separated")
				result = false
			}
		}
		if cfg.SeparateBody {
			if commitMessage[1] != "" {
				fmt.Println("error: body should be separated")
				result = false
			}
		}
	}
	return result
}
