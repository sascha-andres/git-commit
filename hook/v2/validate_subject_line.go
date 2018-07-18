package v2

import "fmt"

// validateSubjectLine checks the subject line for configured patterns
func (cfg *Configuration) validateSubjectLine(subjectLine string) bool {
	result := true
	if "" == subjectLine {
		fmt.Println("error: no subject line provided")
		result = false
	}
	if len(subjectLine) > cfg.SubjectLineLength {
		fmt.Println(fmt.Sprintf("error: subject line is longer than [%d]", cfg.SubjectLineLength))
		result = false
	}
	for _, r := range cfg.SubjectExpressions {
		if !r.match([]byte(subjectLine)) {
			fmt.Println(fmt.Sprintf("%s: [%s] does not match [%s]", r.Severity, subjectLine, r.Expression))
			result = result && r.Severity != ErrorSeverity
		}
	}
	return result
}
