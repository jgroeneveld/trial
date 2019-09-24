package assert

import (
	"fmt"
	"github.com/jgroeneveld/trial/th"
)

const (
	msgActual   = "  Actual"
	msgExpected = "Expected"
	msgTypes    = "   Types"
)

func comparisonError(t testingT, title string, skip int, expected interface{}, actual interface{}, msgf ...interface{}) {
	msg := fmt.Sprintf("%s:\n%s: %#v\n%s: %#v", titleOrMsgf(title, msgf), msgExpected, expected, msgActual, actual)

	expectedT := fmt.Sprintf("%T", expected)
	actualT := fmt.Sprintf("%T", actual)
	if expectedT != actualT {
		msg += fmt.Sprintf("\n%s: Expected:%s, Actual:%s", msgTypes, expectedT, actualT)
	}

	th.Error(t, skip+1, msg)
}

func titleOrMsgf(title string, msgf []interface{}) string {
	if len(msgf) > 0 {
		return msgfToString(msgf)
	}
	return title
}

func msgfToString(args []interface{}) string {
	if len(args) == 1 {
		return fmt.Sprintf("%s", args[0])
	}

	return fmt.Sprintf(args[0].(string), args[1:]...)
}

type testingT interface {
	Error(args ...interface{})
	FailNow()
}
