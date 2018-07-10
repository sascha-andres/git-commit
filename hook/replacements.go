package hook

func (cfg *Configuration) replacements(commitMessage []string) {
	for i, r := range cfg.replaceCompiled {
		for lineIndex, line := range commitMessage {
			if r.Match([]byte(line)) {
				commitMessage[lineIndex] = r.ReplaceAllString(line, cfg.replaceDestination[i])
				break
			}
		}
	}
}
