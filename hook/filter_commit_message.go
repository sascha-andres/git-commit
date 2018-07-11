package hook

// filterCommitMessage removes all lines from the message to check that match configured patterns.
// Usually you want to omit comments with something like ^#.*$
func (cfg *Configuration) filterCommitMessage(commitMessage []string) ([]string, error) {
	result := make([]string, 0)
	for _, line := range commitMessage {
		for _, r := range cfg.ignoreCompiled {
			if !r.Match([]byte(line)) {
				result = append(result, line)
			}
		}
	}
	return result, nil
}
