package hook

import (
	"regexp"
)

func (cfg *Configuration) setupRegularExpressions() error {
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
	if cfg.FirstLineExpressions != nil && len(cfg.FirstLineExpressions) > 0 {
		cfg.firstLineCompiled = make([]*regexp.Regexp, 0)
		for _, expr := range cfg.FirstLineExpressions {
			r, err := regexp.Compile(expr)
			if err != nil {
				return err
			}
			cfg.firstLineCompiled = append(cfg.firstLineCompiled, r)
		}
	}
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
