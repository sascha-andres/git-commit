package v2

import (
	"strings"
	"testing"
)

func TestTicket(t *testing.T) {
	commitMessages := map[string]struct {
		expected bool
		value    string
	}{
		"in-subject": {true, `this is a normal length line TICKET-1

Hello`},
		"in-body": {true, `this is a normal length line

Hello TICKET-1`},
		"not found": {false, "01234567890123456789012345678901234567890123456789012345678901234567890123456789"},
	}

	expr := make([]*ExpressionWithSeverity, 0)

	expr = append(expr, &ExpressionWithSeverity{
		Name:       "find ticket",
		Expression: "^TICKET-\\d{1,3}:.*",
	})

	cfg := &Configuration{
		FindOccurrenceExpressions: expr,
	}

	err := cfg.setupRegularExpressions()
	if err != nil {
		t.Fatalf("error with regex: %s", err)
	}

	for k, v := range commitMessages {
		msg := strings.Split(v.value, "\n")
		t.Run(k, func(t *testing.T) {
			if result := cfg.validateOccurs(msg); result != v.expected {
				t.Logf("test %s failed: it is %t but should be %t", k, result, v.expected)
			}
		})
	}

}
