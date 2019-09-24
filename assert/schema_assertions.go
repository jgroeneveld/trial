package assert

import (
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/th"
	"io"
)

func JSONSchema(t testingT, r io.Reader, matcher schema.Matcher, msgf ...interface{}) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, titleOrMsgf("JSON Schema invalid", msgf))
	}
}

func MustMatchJSONSchema(t testingT, r io.Reader, matcher schema.Matcher, msgf ...interface{}) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		th.Error(t, 1, titleOrMsgf("JSON Schema invalid", msgf))
		t.FailNow()
	}
}
