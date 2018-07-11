package hook

import "regexp"

type (
	// Configuration is used to load global/per-project configuration
	Configuration struct {
		IgnoreExpressions         []string `yaml:"ignore"`                   // IgnoreExpressions is a list of regular expressions that determine whether a line should be checked or not
		SubjectExpressions        []string `yaml:"subject"`                  // SubjectExpressions is a list of regular expressions to check the first line
		FindOccurrenceExpressions []string `yaml:"occurs"`                   // FindOccurrenceExpressions is a list of expressions that have to match at least once
		SubjectLineLength         int      `yaml:"subject-line-length"`      // SubjectLineLength provides the ability to limit the subject line's length
		BodyRequired              bool     `yaml:"body-required"`            // BodyRequired forces a body if set
		SeparateBody              bool     `yaml:"separate-body"`            // SeparateBody forces a blank line between subject and body
		BodyLineLength            int      `yaml:"body-line-length"`         // BodyLineLength provides the ability to limit the body lines' length
		EnforceBodyLineLength     bool     `yaml:"enforce-body-line-length"` // EnforceBodyLineLength determines whether to print a warning when body line length it too long or to error

		ignoreCompiled  []*regexp.Regexp
		subjectCompiled []*regexp.Regexp
		occursCompiled  []*regexp.Regexp
	}
)
