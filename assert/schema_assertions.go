package assert

import (
	"fmt"
	"github.com/jgroeneveld/schema"
	"github.com/jgroeneveld/trial/th"
	"io"
)

func JSONSchema(t testingT, r io.Reader, matcher schema.Matcher, msgf ...interface{}) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		msg := titleOrMsgf("JSON Schema invalid", msgf)
		msg += fmt.Sprintf("\n%s", err.Error())
		th.Error(t, 1, msg)
	}
}

func MustMatchJSONSchema(t testingT, r io.Reader, matcher schema.Matcher, msgf ...interface{}) {
	err := schema.MatchJSON(matcher, r)
	if err != nil {
		msg := titleOrMsgf("JSON Schema invalid", msgf)
		msg += fmt.Sprintf("\n%s", err.Error())
		th.Error(t, 1, msg)
		t.FailNow()
	}
}
