package assert

import (
	"testing"
)

func TestAsserterEqual(t *testing.T) {
	mt := &mockT{}

	asserter := Asserter(mt)

	asserter.Equal("a", "a")
	asserter.Equal("a", "b")
	asserter.Equal(1, int64(1), "Test (%s)", "value")

	if len(mt.errors) != 2 {
		t.Fatalf("wrong number of errors %d", len(mt.errors))
	}

	actual := mt.errors[0]
	expected := "\r" + `asserter_test.go:13: Not equal:
Expected: "a"
  Actual: "b"`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}

	actual = mt.errors[1]
	expected = "\r" + `asserter_test.go:14: Test (value):
Expected: 1
  Actual: 1
   Types: Expected:int, Actual:int64`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}
}
