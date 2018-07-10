package hook

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
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
	fmt.Println("info: setup replacements regular expressions")
	if cfg.ReplaceExpressions != nil && len(cfg.ReplaceExpressions) > 0 {
		cfg.replaceCompiled = make([]*regexp.Regexp, 0)
		cfg.replaceDestination = make([]string, 0)
		for _, expr := range cfg.ReplaceExpressions {
			splitExpression := strings.Split(expr, "-->")
			if len(splitExpression) != 2 {
				return errors.New(fmt.Sprintf("error: replacement config invalid: %s", expr))
			}
			r, err := regexp.Compile(splitExpression[0])
			if err != nil {
				return err
			}
			cfg.replaceCompiled = append(cfg.replaceCompiled, r)
			cfg.replaceDestination = append(cfg.replaceDestination, splitExpression[1])
		}
	}
	return nil
}
