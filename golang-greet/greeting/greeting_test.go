package greeting

import (
	"regexp"
	"testing"
)

// TestHelloName calls greeting.Hello with a name,
// checking for a valid return value.
func TestHelloName(t *testing.T) {
	var (
		name         = "Gladys"
		want         = regexp.MustCompile(`\b` + name + `\b`)
		message, err = Hello("Gladys")
	)

	if !want.MatchString(message) || err != nil {
		t.Fatalf(`Hello("Gladys") = %q, %v, want match for %#q, nil`, message, err, want)
	}
}

func TestHelloEmpty(t *testing.T) {
	message, err := Hello("")
	if message != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, message, err)
	}
}
