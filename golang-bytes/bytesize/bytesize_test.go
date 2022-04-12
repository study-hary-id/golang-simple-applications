package bytesize

import (
	"fmt"
	"testing"
)

// TestStringBytes prints ByteSize to string with value less than 1MB.
func TestStringBytes(t *testing.T) {
	var (
		want = "2.00KB"
		str  = fmt.Sprint(ByteSize(2048))
	)

	if str != want {
		t.Fatalf(`ByteSize(2048) = %q, want match for %q, nil`, str, want)
	}
}

// TestStringMegaBytes prints ByteSize to string with value 3.14MB.
func TestStringMegaBytes(t *testing.T) {
	var (
		want = "3.14MB"
		str  = fmt.Sprint(ByteSize(3292528.64))
	)

	if str != want {
		t.Fatalf(`ByteSize(3292528.64) = %q, want match for %q, nil`, str, want)
	}
}

// TestLessThanKiloBytes prints ByteSize to string with value 512B.
func TestLessThanKiloBytes(t *testing.T) {
	var (
		want = "512.0B"
		str  = fmt.Sprint(ByteSize(512))
	)

	if str != want {
		t.Fatalf(`ByteSize(512) = %q, want match for %q, nil`, str, want)
	}
}
