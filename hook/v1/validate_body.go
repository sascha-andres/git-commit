package v1

import "fmt"

// validateBody runs all rules for a commit message body
func (cfg *Configuration) validateBody(commitMessage []string) bool {
	result := true
	for _, line := range commitMessage {
		if len(line) > cfg.BodyLineLength {
			if cfg.EnforceBodyLineLength {
				result = false
				fmt.Println(fmt.Sprintf("error: body line is longer than %d", cfg.BodyLineLength))
			} else {
				fmt.Println(fmt.Sprintf("warn: body line is longer than %d", cfg.BodyLineLength))
			}
		}
	}
	result = cfg.validateBodySeparation(commitMessage) && result
	return result
}
