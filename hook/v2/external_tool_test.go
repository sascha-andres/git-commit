package v2

import "testing"

func TestExternalTool(t *testing.T) {
	tests := map[string]struct {
		command  string
		severity string
		result   bool
	}{
		"success": {"/bin/true", ErrorSeverity, true},
		"warning": {"/bin/false", "warning", false},
		"failure": {"/bin/false", ErrorSeverity, false},
	}

	for k, v := range tests {
		t.Run(k, func(t *testing.T) {
			tools := make([]*Tool, 0)
			commands := make([]string, 0)
			commands = append(commands, v.command)
			tools = append(tools, &Tool{
				Name:     v.command,
				Severity: v.severity,
				Command:  commands,
			})
			cfg := &Configuration{
				Externals: tools,
			}
			if result := cfg.runExternalTools(); result != v.result {
				t.Logf("for %s expected %t, received %t", k, v.result, result)
			}
		})
	}
}
