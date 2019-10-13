package assert

import (
	"fmt"
	"strings"
	"testing"
)

type mockT struct {
	errors []string
	failedNow bool
}

func (t *mockT) Error(args ...interface{}) {
	t.errors = append(t.errors, toString(args))
}

func (t *mockT) FailNow() {
	t.failedNow = true
}

func toString(args []interface{}) string {
	var str []string

	for _, arg := range args {
		str = append(str, fmt.Sprintf("%s", arg))
	}

	return strings.Join(str, " ")
}

func testAssertionSuccess(t *testing.T, successCase func(mt *mockT)) {
	mt := &mockT{}
	successCase(mt)
	if len(mt.errors) > 0 {
		t.Fatalf("expected no errors but got %#v", mt.errors)
	}
}

func testAssertionFail(t *testing.T, failureCase func(mt *mockT)) {
	mt := &mockT{}
	failureCase(mt)
	if len(mt.errors) == 0 {
		t.Fatalf("expected errors but got none")
	}
}

func testAssertionFailNow(t *testing.T, failureCase func(mt *mockT)) {
	mt := &mockT{}
	failureCase(mt)
	if len(mt.errors) == 0 {
		t.Fatalf("expected errors but got none")
	}
	if !mt.failedNow {
		t.Fatalf("expected to fail now but did not")
	}
}
