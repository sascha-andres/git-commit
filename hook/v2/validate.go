package v2

// Validate validates the message against the rules
func (cfg *Configuration) Validate(commitMessage []string) (bool, error) {
	err := cfg.setupRegularExpressions()
	if err != nil {
		return false, err
	}
	filteredMessage, err := cfg.filterCommitMessage(commitMessage)
	if err != nil {
		return false, err
	}
	result := cfg.validateSubjectLine(filteredMessage[0])
	result = cfg.validateBody(filteredMessage) && result
	result = cfg.validateOccurs(filteredMessage) && result
	result = cfg.runExternalTools() && result
	result = cfg.validateGitLabCI() && result

	return result, nil
}
