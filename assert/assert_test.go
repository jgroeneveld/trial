package assert

import (
	"fmt"
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {
	mt := &mockT{}

	Equal(mt, 1, "1")

	actual := mt.errors[0]

	expected := "\r" + `assert/assert_test.go:12: Not Equal:
Expected: 1
  Actual: "1"
int != string`

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
