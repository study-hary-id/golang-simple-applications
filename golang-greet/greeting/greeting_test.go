package greeting

import (
	"regexp"
	"testing"
)

// TestHelloName calls greeting.Helo with a name,
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	var (
		name           = "Gladys"
		want           = regexp.MustCompile(`\b` + name + `\b`)
		message, error = Hello("Gladys")
	)

	if !want.MatchString(message) || error != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, error, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, error := Hello("")
	if message != "" || error == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, error)
	}
}
