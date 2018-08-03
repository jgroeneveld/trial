package assert

import (
	"fmt"
	"github.com/jgroeneveld/trial/th"
	"reflect"
)

func Asserter(t testingT) *asserter {
	return &asserter{t}
}

type asserter struct {
	t testingT
}

func (a *asserter) Equal(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(a.t, "Not equal", 1, expected, actual, msgf...)
	}
}

func (a *asserter) MustBeEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(a.t, "Not equal", 1, expected, actual, msgf...)
		a.t.FailNow()
	}
}

func (a *asserter) NotEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(a.t, 1, msg)
	}
}

func (a *asserter) MustNotBeEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(a.t, 1, msg)
		a.t.FailNow()
	}
}

func (a *asserter) DeepEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		comparisonError(a.t, "Not deep equal", 1, expected, actual, msgf...)
	}
}

func (a *asserter) MustBeDeepEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		comparisonError(a.t, "Not deep equal", 1, expected, actual, msgf...)
		a.t.FailNow()
	}
}

func (a *asserter) True(expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(a.t, 1, titleOrMsgf("Not true", msgf))
	}
}

func (a *asserter) MustBeTrue(expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(a.t, 1, titleOrMsgf("Not true", msgf))
		a.t.FailNow()
	}
}

func (a *asserter) False(expression bool, msgf ...interface{}) {
	if expression {
		th.Error(a.t, 1, titleOrMsgf("Not false", msgf))
	}
}

func (a *asserter) MustBeFalse(expression bool, msgf ...interface{}) {
	if expression {
		th.Error(a.t, 1, titleOrMsgf("Not false", msgf))
		a.t.FailNow()
	}
}

func (a *asserter) Nil(expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(a.t, 1, msg)
	}
}

func (a *asserter) MustBeNil(expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(a.t, 1, msg)
		a.t.FailNow()
	}
}

func (a *asserter) NotNil(expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(a.t, 1, titleOrMsgf("Is nil", msgf))
	}
}

func (a *asserter) MustNotBeNil(expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(a.t, 1, titleOrMsgf("Is nil", msgf))
		a.t.FailNow()
	}
}
