package v2

import "testing"

func TestLineLengthLastLine(t *testing.T) {
	body := []string{
		"this is a normal length line",
		"01234567890123456789012345678901234567890123456789012345678901234567890123456789",
	}
	cfg := &Configuration{
		BodyLineLength:        72,
		EnforceBodyLineLength: true,
	}
	if cfg.validateBody(body) {
		t.Log("line is too long, validation did not get it")
		t.Fail()
	}
}

func TestLineLengthFirstLine(t *testing.T) {
	body := []string{
		"01234567890123456789012345678901234567890123456789012345678901234567890123456789",
		"this is a normal length line",
	}
	cfg := &Configuration{
		BodyLineLength:        72,
		EnforceBodyLineLength: true,
	}
	if cfg.validateBody(body) {
		t.Log("line is too long, validation did not get it")
		t.Fail()
	}
}

func TestLineLengthCoAuthored(t *testing.T) {
	body := []string{
		"this is a normal length line",
		"Co-authored-by: This is the name <this-is-a-long@email-address-to-get-over71.de>",
	}
	cfg := &Configuration{
		BodyLineLength:        72,
		EnforceBodyLineLength: true,
	}
	if !cfg.validateBody(body) {
		t.Log("co-authored lines should not be enforced to be max characters long")
		t.Fail()
	}
}
