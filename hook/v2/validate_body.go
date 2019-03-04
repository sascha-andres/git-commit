package v2

import (
	"fmt"
	"strings"
)

// validateBody runs all rules for a commit message body
func (cfg *Configuration) validateBody(commitMessage []string) bool {
	result := true
	for _, line := range commitMessage {
		if strings.HasPrefix(line, "Co-authored-by") {
			fmt.Println("Co-authored-by found, not checking line length")
			continue
		}
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
