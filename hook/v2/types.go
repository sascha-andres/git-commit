package v2

import "regexp"

// ErrorSeverity set to issue an error and terminate commit
const ErrorSeverity = "error"

type (

	// Configuration is used to load global/per-project configuration
	// codebeat:disable[TOO_MANY_IVARS]
	Configuration struct {
		Version                   string                    `yaml:"version"`                  // Version denotes the configuration file version
		IgnoreExpressions         []string                  `yaml:"ignore"`                   // IgnoreExpressions is a list of regular expressions that determine whether a line should be checked or not
		SubjectExpressions        []*ExpressionWithSeverity `yaml:"subject"`                  // SubjectExpressions is a list of regular expressions to check the first line
		FindOccurrenceExpressions []*ExpressionWithSeverity `yaml:"occurs"`                   // FindOccurrenceExpressions is a list of expressions that have to match at least once
		SubjectLineLength         int                       `yaml:"subject-line-length"`      // SubjectLineLength provides the ability to limit the subject line's length
		BodyRequired              bool                      `yaml:"body-required"`            // BodyRequired forces a body if set
		SeparateBody              bool                      `yaml:"separate-body"`            // SeparateBody forces a blank line between subject and body
		BodyLineLength            int                       `yaml:"body-line-length"`         // BodyLineLength provides the ability to limit the body lines' length
		EnforceBodyLineLength     bool                      `yaml:"enforce-body-line-length"` // EnforceBodyLineLength determines whether to print a warning when body line length it too long or to error
		ExternalChecks            []string                  `yaml:"calls"`                    // ExternalChecks contains a list of commands to execute

		ignoreCompiled []*regexp.Regexp
	}
	// codebeat:enable[TOO_MANY_IVARS]

	// ExpressionWithSeverity is a container for expressions that may be errors or warnings
	ExpressionWithSeverity struct {
		Expression string `yaml:"expression"` // Regular expression
		Severity   string `yaml:"severity"`   // Severity name
		Name       string `yaml:"name"`       // Name of expression

		compiled *regexp.Regexp
	}
)
