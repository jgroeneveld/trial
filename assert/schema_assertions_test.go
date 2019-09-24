package assert

import (
	"github.com/jgroeneveld/schema"
	"strings"
	"testing"
)

func TestJSONSchema(t *testing.T) {
	mt := &mockT{}
	jsonString := `{ "foo": "bar" }`
	matchingSchema := schema.Map{
		"foo": schema.IsPresent,
	}
	notMatchingSchema := schema.Map{
		"foo": schema.IsPresent,
		"something_else": schema.IsPresent,
	}

	JSONSchema(mt, strings.NewReader(jsonString), matchingSchema)
	JSONSchema(mt, strings.NewReader(jsonString), notMatchingSchema)

	if len(mt.errors) != 1 {
		t.Fatalf("wrong number of errors %d", len(mt.errors))
	}

	actual := mt.errors[0]
	expected := "\r" + `schema_assertions_test.go:21: JSON Schema invalid
Missing keys: "something_else"`

	if expected != actual {
		t.Errorf("wrong output:\nexpected:\n%q\nactual:\n%q", expected, actual)
	}
}