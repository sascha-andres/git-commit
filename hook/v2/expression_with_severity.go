package v2

import (
	"regexp"
)

func (e *ExpressionWithSeverity) setup() error {
	if err := e.compile(); err != nil {
		return err
	}
	if e.Severity == "" {
		e.Severity = ErrorSeverity
	}
	return nil
}

// compile compiles the given expression to a regex
func (e *ExpressionWithSeverity) compile() error {
	r, err := regexp.Compile(e.Expression)
	if err != nil {
		return err
	}
	e.compiled = r
	return nil
}

// match checks if given value matches
func (e *ExpressionWithSeverity) match(value []byte) bool {
	return e.compiled.Match(value)
}
