package main

import (
	"regexp"
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestGetMessage(t *testing.T) {
	name := "SLIM"
	want := regexp.MustCompile(`\b` + name + `\b`)
	msg, err := GetMessage()
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`GetMessage() = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}
