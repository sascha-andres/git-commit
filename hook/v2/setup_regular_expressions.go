package v2

import (
	"regexp"
)

// setupRegularExpressions compiles all configured regular expressions
func (cfg *Configuration) setupRegularExpressions() error {
	if err := cfg.regexIgnore(); err != nil {
		return err
	}
	if err := cfg.regexSubject(); err != nil {
		return err
	}
	return cfg.regexOccur()
}

// regexOccur compiles occur regular expressions
func (cfg *Configuration) regexOccur() error {
	return loadExpressionsWithSeverity(cfg.FindOccurrenceExpressions)
}

// regexIgnore compiles ignore regular expressions
func (cfg *Configuration) regexIgnore() error {
	c, err := loadExpressions(cfg.IgnoreExpressions)
	if err != nil {
		return err
	}
	cfg.ignoreCompiled = c
	return nil
}

// regexSubject compiles regular expressions for the subject line
func (cfg *Configuration) regexSubject() error {
	return loadExpressionsWithSeverity(cfg.SubjectExpressions)
}

// loadExpressions compiles a list of regular expressions to a list of Regexp
func loadExpressions(expressions []string) ([]*regexp.Regexp, error) {
	if len(expressions) > 0 {
		compiled := make([]*regexp.Regexp, 0)
		for _, expr := range expressions {
			r, err := regexp.Compile(expr)
			if err != nil {
				return nil, err
			}
			compiled = append(compiled, r)
		}
		return compiled, nil
	}
	return nil, nil
}

// loadExpressionsWithSeverity compiles a list of regular expressions with severity
func loadExpressionsWithSeverity(expressions []ExpressionWithSeverity) error {
	if len(expressions) > 0 {
		for _, expr := range expressions {
			err := expr.setup()
			if err != nil {
				return err
			}
		}
	}
	return nil
}
