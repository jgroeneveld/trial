package assert

import (
	"fmt"
	"github.com/jgroeneveld/trial/th"
	"reflect"
)

// Asserter wraps testingT so that you do not have to pass it into each matcher
func Asserter(t testingT) *AsserterImpl {
	return &AsserterImpl{t}
}

// AsserterImpl wraps testingT so that you do not have to pass it into each matcher
type AsserterImpl struct {
	t testingT
}

// Equal checks if two values are equal.
// Reports if types are different, even though values "look" the same.
func (a *AsserterImpl) Equal(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(a.t, "Not equal", 1, expected, actual, msgf...)
	}
}

// MustBeEqual checks if two values are equal.
// Reports if types are different, even though values "look" the same.
// Will FailNow if expectation is not met.
func (a *AsserterImpl) MustBeEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(a.t, "Not equal", 1, expected, actual, msgf...)
		a.t.FailNow()
	}
}

// NotEqual checks if two values are not equal.
func (a *AsserterImpl) NotEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(a.t, 1, msg)
	}
}

// MustNotBeEqual checks if two values are equal.
// Will FailNow if expectation is not met.
func (a *AsserterImpl) MustNotBeEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(a.t, 1, msg)
		a.t.FailNow()
	}
}

// DeepEqual test if two values are deeply equal.
func (a *AsserterImpl) DeepEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		comparisonError(a.t, "Not deep equal", 1, expected, actual, msgf...)
	}
}

// MustBeDeepEqual test if two values are deeply equal.
// Will FailNow if expectation is not met.
func (a *AsserterImpl) MustBeDeepEqual(expected interface{}, actual interface{}, msgf ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		comparisonError(a.t, "Not deep equal", 1, expected, actual, msgf...)
		a.t.FailNow()
	}
}

// True tests if a value is true
func (a *AsserterImpl) True(expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(a.t, 1, titleOrMsgf("Not true", msgf))
	}
}

// MustBeTrue tests if a value is true
// Will FailNow if expectation is not met.
func (a *AsserterImpl) MustBeTrue(expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(a.t, 1, titleOrMsgf("Not true", msgf))
		a.t.FailNow()
	}
}

// False tests if a value is False
func (a *AsserterImpl) False(expression bool, msgf ...interface{}) {
	if expression {
		th.Error(a.t, 1, titleOrMsgf("Not false", msgf))
	}
}

// MustBeFalse tests if a value is False
// Will FailNow if expectation is not met.
func (a *AsserterImpl) MustBeFalse(expression bool, msgf ...interface{}) {
	if expression {
		th.Error(a.t, 1, titleOrMsgf("Not false", msgf))
		a.t.FailNow()
	}
}

// Nil tests if a value is Nil
func (a *AsserterImpl) Nil(expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(a.t, 1, msg)
	}
}

// MustBeNil tests if a value is Nil
// Will FailNow if expectation is not met.
func (a *AsserterImpl) MustBeNil(expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(a.t, 1, msg)
		a.t.FailNow()
	}
}

// NotNil tests if a value is Not Nil
func (a *AsserterImpl) NotNil(expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(a.t, 1, titleOrMsgf("Is nil", msgf))
	}
}

// MustNotBeNil tests if a value is Not Nil
// Will FailNow if expectation is not met.
func (a *AsserterImpl) MustNotBeNil(expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(a.t, 1, titleOrMsgf("Is nil", msgf))
		a.t.FailNow()
	}
}
