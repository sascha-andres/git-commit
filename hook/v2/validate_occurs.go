package v2

import "fmt"

// validateOccurs checks whether required patterns are found in the commit message independently
// of the place ( subject or body )
func (cfg *Configuration) validateOccurs(commitMessage []string) bool {
	result := true
	for _, r := range cfg.FindOccurrenceExpressions {
		localOK := false
		for _, line := range commitMessage {
			if r.match([]byte(line)) {
				localOK = true
				break
			}
		}
		if !localOK {
			fmt.Println(fmt.Sprintf("%s: nothing found that matches [%s]", r.Severity, r.Expression))
			result = result && r.Severity != ErrorSeverity
		}
	}
	return result
}
