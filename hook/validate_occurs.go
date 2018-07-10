package hook

import "fmt"

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
