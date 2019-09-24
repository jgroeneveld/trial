package assert

import (
	"fmt"
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/th"
	"io"
)

// JSONSchema uses schema to assert that json matches a given structure
func JSONSchema(t testingT, r io.Reader, matcher schema.Matcher, msgf ...interface{}) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		msg := titleOrMsgf("JSON Schema invalid", msgf)
		msg += fmt.Sprintf("\n%s", err.Error())
		th.Error(t, 1, msg)
	}
}

// MustMatchJSONSchema uses schema to assert that json matches a given structure
// Will FailNow if expectation is not met.
func MustMatchJSONSchema(t testingT, r io.Reader, matcher schema.Matcher, msgf ...interface{}) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		msg := titleOrMsgf("JSON Schema invalid", msgf)
		msg += fmt.Sprintf("\n%s", err.Error())
		th.Error(t, 1, msg)
		t.FailNow()
	}
}
