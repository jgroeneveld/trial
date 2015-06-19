package assert

import (
	"fmt"
	"github.com/jgroeneveld/go-test/th"
)

const (
	msgActual   = "  Actual"
	msgExpected = "Expected"
	msgTypes    = "   Types"
)

func Equal(t TBInt, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(t, "Not equal", 1, expected, actual, msgf...)
	}
}

func MustBeEqual(t TBInt, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(t, "Not equal", 1, expected, actual, msgf...)
		t.FailNow()
	}
}

func NotEqual(t TBInt, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(t, 1, msg)
	}
}

func MustNotBeEqual(t TBInt, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(t, 1, msg)
		t.FailNow()
	}
}

func True(t TBInt, expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(t, 1, titleOrMsgf("Not true", msgf))
	}
}

func MustBeTrue(t TBInt, expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(t, 1, titleOrMsgf("Not true", msgf))
		t.FailNow()
	}
}

func False(t TBInt, expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(t, 1, titleOrMsgf("Not false", msgf))
	}
}

func MustBeFalse(t TBInt, expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(t, 1, titleOrMsgf("Not false", msgf))
		t.FailNow()
	}
}

func Nil(t TBInt, expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(t, 1, msg)
	}
}

func MustBeNil(t TBInt, expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(t, 1, msg)
		t.FailNow()
	}
}

func NotNil(t TBInt, expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(t, 1, titleOrMsgf("Is nil", msgf))
	}
}

func MustNotBeNil(t TBInt, expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(t, 1, titleOrMsgf("Is nil", msgf))
		t.FailNow()
	}
}

func comparisonError(t TBInt, title string, skip int, expected interface{}, actual interface{}, msgf ...interface{}) {
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

type TBInt interface {
	Error(args ...interface{})
	FailNow()
}
