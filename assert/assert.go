package assert

import (
	"fmt"
	"github.com/jgroeneveld/go-test/th"
)

const (
	msgContext  = " Context"
	msgActual   = "  Actual"
	msgExpected = "Expected"
	msgTypes    = "   Types"
)

func Equal(t TBInt, expected interface{}, actual interface{}, args ...interface{}) {
	if expected != actual {
		comparisonError(t, "Not Equal", 1, expected, actual, args...)
	}
}

func MustBeEqual(t TBInt, expected interface{}, actual interface{}, args ...interface{}) {
	if expected != actual {
		comparisonError(t, "Not Equal", 1, expected, actual, args...)
		t.FailNow()
	}
}

func True(t TBInt, expression bool, args ...interface{}) {
	if !expression {
		msg := "Not True"
		if len(args) > 0 {
			msg += fmt.Sprintf("\n%s: %s", msgContext, argsToString(args))
		}
		th.Error(t, msg, 1)
	}
}

func MustBeTrue(t TBInt, expression bool, args ...interface{}) {
	if !expression {
		msg := "Not True"
		if len(args) > 0 {
			msg += fmt.Sprintf("\n%s: %s", msgContext, argsToString(args))
		}
		th.Error(t, msg, 1)
		t.FailNow()
	}
}

func False(t TBInt, expression bool, args ...interface{}) {
	if !expression {
		msg := "Not False"
		if len(args) > 0 {
			msg += fmt.Sprintf("\n%s: %s", msgContext, argsToString(args))
		}
		th.Error(t, msg, 1)
	}
}

func MustBeFalse(t TBInt, expression bool, args ...interface{}) {
	if !expression {
		msg := "Not False"
		if len(args) > 0 {
			msg += fmt.Sprintf("\n%s: %s", msgContext, argsToString(args))
		}
		th.Error(t, msg, 1)
		t.FailNow()
	}
}

func Nil(t TBInt, expression interface{}, args ...interface{}) {
	if expression != nil {
		msg := "Not nil"
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		if len(args) > 0 {
			msg += fmt.Sprintf("\n%s: %s", msgContext, argsToString(args))
		}
		th.Error(t, msg, 1)
	}
}

func MustBeNil(t TBInt, expression interface{}, args ...interface{}) {
	if expression != nil {
		msg := "Not nil"
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		if len(args) > 0 {
			msg += fmt.Sprintf("\n%s: %s", msgContext, argsToString(args))
		}
		th.Error(t, msg, 1)
		t.FailNow()
	}
}

func comparisonError(t TBInt, prefix string, skip int, expected interface{}, actual interface{}, args ...interface{}) {
	msg := fmt.Sprintf("%s:\n%s: %#v\n%s: %#v", prefix, msgExpected, expected, msgActual, actual)

	expectedT := fmt.Sprintf("%T", expected)
	actualT := fmt.Sprintf("%T", actual)
	if expectedT != actualT {
		msg += fmt.Sprintf("\n%s: Expected:%s, Actual:%s", msgTypes, expectedT, actualT)
	}

	if len(args) > 0 {
		msg += fmt.Sprintf("\n%s: %s", msgContext, argsToString(args))
	}

	th.Error(t, msg, skip+1)
}

func argsToString(args []interface{}) string {
	if len(args) == 0 {
		return ""
	}

	if len(args) == 1 {
		return fmt.Sprintf("%s", args[0])
	}

	return fmt.Sprintf(args[0].(string), args[1:]...)
}

type TBInt interface {
	Error(args ...interface{})
	FailNow()
}
