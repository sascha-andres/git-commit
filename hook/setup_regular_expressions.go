package hook

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
	if cfg.FindOccurrenceExpressions != nil && len(cfg.FindOccurrenceExpressions) > 0 {
		cfg.occursCompiled = make([]*regexp.Regexp, 0)
		for _, expr := range cfg.FindOccurrenceExpressions {
			r, err := regexp.Compile(expr)
			if err != nil {
				return err
			}
			cfg.occursCompiled = append(cfg.occursCompiled, r)
		}
	}
	return nil
}

// regexIgnore compiles ignore regular expressions
func (cfg *Configuration) regexIgnore() error {
	if cfg.IgnoreExpressions != nil && len(cfg.IgnoreExpressions) > 0 {
		cfg.ignoreCompiled = make([]*regexp.Regexp, 0)
		for _, expr := range cfg.IgnoreExpressions {
			r, err := regexp.Compile(expr)
			if err != nil {
				return err
			}
			cfg.ignoreCompiled = append(cfg.ignoreCompiled, r)
		}
	}
	return nil
}

// regexSubject compiles regular expressions for the subject line
func (cfg *Configuration) regexSubject() error {
	if cfg.SubjectExpressions != nil && len(cfg.SubjectExpressions) > 0 {
		cfg.subjectCompiled = make([]*regexp.Regexp, 0)
		for _, expr := range cfg.SubjectExpressions {
			r, err := regexp.Compile(expr)
			if err != nil {
				return err
			}
			cfg.subjectCompiled = append(cfg.subjectCompiled, r)
		}
	}
	return nil
}
