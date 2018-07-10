package hook

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/pkg/errors"
)

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
	result = cfg.validateBodySeparation(filteredMessage) && result
	result = cfg.validateBody(filteredMessage) && result
	result = cfg.validateOccurs(filteredMessage) && result
	cfg.replacements(filteredMessage)
	return result, nil
}

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

func (cfg *Configuration) filterCommitMessage(commitMessage []string) ([]string, error) {
	result := make([]string, 0)
	for _, line := range commitMessage {
		for _, r := range cfg.ignoreCompiled {
			if !r.Match([]byte(line)) {
				result = append(result, line)
			}
		}
	}
	return result, nil
}

func (cfg *Configuration) validateBodySeparation(commitMessage []string) bool {
	result := true
	if cfg.BodyRequired {
		if len(commitMessage) == 1 {
			fmt.Println("error: body required but not provided")
			result = false
		} else {
			if len(commitMessage) == 2 {
				if cfg.SeparateBody {
					fmt.Println("error: body should be separated")
					result = false
				}
			}
			if cfg.SeparateBody {
				if commitMessage[1] != "" {
					fmt.Println("error: body should be separated")
					result = false
				}
			}
		}
	}
	if cfg.SeparateBody && len(commitMessage) > 1 {
		if len(commitMessage) == 2 {
			if cfg.SeparateBody {
				fmt.Println("error: body should be separated")
				result = false
			}
		}
		if cfg.SeparateBody {
			if commitMessage[1] != "" {
				fmt.Println("error: body should be separated")
				result = false
			}
		}
	}
	return result
}

func (cfg *Configuration) validateBody(commitMessage []string) bool {
	result := true
	for _, line := range commitMessage {
		if len(line) > cfg.BodyLineLength {
			if cfg.EnforceBodyLineLength {
				result = false
				fmt.Println(fmt.Sprintf("error: body line is longer than %d", cfg.BodyLineLength))
			} else {
				fmt.Println(fmt.Sprintf("warn: body line is longer than %d", cfg.BodyLineLength))
			}
		}
	}
	return result
}

func (cfg *Configuration) validateSubjectLine(subjectLine string) bool {
	result := true
	if "" == subjectLine {
		fmt.Println("error: no subject line provided")
		result = false
	}
	for _, r := range cfg.firstLineCompiled {
		if !r.Match([]byte(subjectLine)) {
			fmt.Println(fmt.Sprintf("error: [%s] does not match [%s]", subjectLine, r.String()))
			result = false
		}
	}
	return result
}

func (cfg *Configuration) validateOccurs(commitMessage []string) bool {
	result := true
	for _, r := range cfg.occursCompiled {
		localOK := false
		for _, line := range commitMessage {
			if r.Match([]byte(line)) {
				localOK = true
				break
			}
		}
		if !localOK {
			fmt.Println(fmt.Sprintf("error: nothing found that matches [%s]", r.String()))
			result = false
		}
	}
	return result
}

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
