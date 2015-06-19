package assert

import (
	"fmt"
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {
	mt := &mockT{}

	Equal(mt, "a", "a")
	Equal(mt, "a", "b")
	Equal(mt, 1, int64(1), "Test (%s)", "value")

	if len(mt.errors) != 2 {
		t.Fatalf("wrong number of errors %d", len(mt.errors))
	}

	actual := mt.errors[0]
	expected := "\r" + `assert/assert_test.go:13: Not equal:
Expected: "a"
  Actual: "b"`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}

	actual = mt.errors[1]
	expected = "\r" + `assert/assert_test.go:14: Test (value):
Expected: 1
  Actual: 1
   Types: Expected:int, Actual:int64`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}
}

type mockT struct {
	errors []string
}

func (t *mockT) Error(args ...interface{}) {
	t.errors = append(t.errors, toString(args))
}

func (t *mockT) FailNow() {

}

func toString(args []interface{}) string {
	var str []string

	for _, arg := range args {
		str = append(str, fmt.Sprintf("%s", arg))
	}

	return strings.Join(str, " ")
}
