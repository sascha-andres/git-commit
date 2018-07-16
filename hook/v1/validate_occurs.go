package v1

import "fmt"

// validateOccurs checks whether required patterns are found in the commit message independently
// of the place ( subject or body )
func (cfg *Configuration) validateOccurs(commitMessage []string) bool {
	result := true
	for _, r := range cfg.occursCompiled {
		localOK := false
		for _, line := range commitMessage {
			if r.Match([]byte(line)) {
				localOK = true
				break
			}
		}
		if !localOK {
			fmt.Println(fmt.Sprintf("error: nothing found that matches [%s]", r.String()))
			result = false
		}
	}
	return result
}
