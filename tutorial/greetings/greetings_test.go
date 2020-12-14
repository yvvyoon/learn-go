package greetings

import (
	"testing"
	"regexp"
)

// testing.T의 포인터를 인자로 사용
func TestHelloName(t *testing.T) {
	name := "Yoon"
	want := regexp.MustCompile(`\b`+name+`\b`)
	msg, err := Hello("Yoon")

	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Yoon") = %q, %v, want match for %#q, nil`, msg, err, want)
	}
}

// testing.T의 포인터를 인자로 사용
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("")

	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
	}
}


