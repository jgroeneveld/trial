package assert

import (
	"github.com/jgroeneveld/schema"
	"strings"
	"testing"
)

var jsonString = `{ "foo": "bar" }`
var matchingSchema = schema.Map{
	"foo": schema.IsPresent,
}
var notMatchingSchema = schema.Map{
	"foo":            schema.IsPresent,
	"something_else": schema.IsPresent,
}

func TestJSONSchema(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		JSONSchema(mt, strings.NewReader(jsonString), matchingSchema)
	})

	testAssertionFail(t, func(mt *mockT) {
		JSONSchema(mt, strings.NewReader(jsonString), notMatchingSchema)
	})
}

func TestMustMatchJSONSchema(t *testing.T) {
	testAssertionSuccess(t, func(mt *mockT) {
		MustMatchJSONSchema(mt, strings.NewReader(jsonString), matchingSchema)
	})

	testAssertionFailNow(t, func(mt *mockT) {
		MustMatchJSONSchema(mt, strings.NewReader(jsonString), notMatchingSchema)
	})
}

func TestJSONSchema_Output(t *testing.T) {
	mt := &mockT{}

	JSONSchema(mt, strings.NewReader(jsonString), notMatchingSchema)

	if len(mt.errors) != 1 {
		t.Fatalf("wrong number of errors %d", len(mt.errors))
	}

	actual := mt.errors[0]
	expected := "\r" + `schema_assertions_test.go:41: JSON Schema invalid
Missing keys: "something_else"`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}
}
