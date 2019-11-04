package v2

import "testing"

func TestSubjectLineLength(t *testing.T) {
	subjectLines := map[string]struct {
		expected bool
		value    string
	}{
		"right-length":    {true, "this is a normal length line"},
		"too-long":        {false, "01234567890123456789012345678901234567890123456789012345678901234567890123456789"},
		"no-subject-line": {false, ""},
	}

	cfg := &Configuration{
		SubjectLineLength: 50,
	}

	for k, v := range subjectLines {
		t.Run(k, func(t *testing.T) {
			if result := cfg.validateSubjectLine(v.value); result != v.expected {
				t.Logf("test %s failed: it is %t but should be %t", k, result, v.expected)
			}
		})
	}
}

func TestOccurs(t *testing.T) {
	subjectLines := map[string]struct {
		expected bool
		value    string
	}{
		"found ticket":     {true, "TICKET-1: hello world"},
		"no ticket number": {false, "hello world without ticket"},
	}

	expr := make([]*ExpressionWithSeverity, 0)

	expr = append(expr, &ExpressionWithSeverity{
		Name:       "find ticket",
		Expression: "^TICKET-\\d{1,3}:.*",
	})

	cfg := &Configuration{
		SubjectExpressions: expr,
		SubjectLineLength:  50,
	}

	err := cfg.setupRegularExpressions()
	if err != nil {
		t.Fatalf("error with regex: %s", err)
	}

	for k, v := range subjectLines {
		t.Run(k, func(t *testing.T) {
			if result := cfg.validateSubjectLine(v.value); result != v.expected {
				t.Logf("test %s failed: it is %t but should be %t", k, result, v.expected)
			}
		})
	}
}
