package hook

import (
	"fmt"
	"regexp"
)

func (cfg *Configuration) setupRegularExpressions() error {
	fmt.Println("info: setup regular expressions")
	fmt.Println("info: setup ignore regular expressions")
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
	fmt.Println("info: setup subject regular expressions")
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
	fmt.Println("info: setup occurs regular expressions")
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
