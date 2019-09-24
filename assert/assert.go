// Package assert (trial/assert) is a simple and lightweight assertion library.
//
// Example:
//
//  import "github.com/jgroeneveld/trial/assert"
//
//  assert.Equal(t, 1, 2)
//
//  Output:
//  unit_test.go:42: Not equal:
// 		Expected: 1
// 		  Actual: 2
//
// Additional arguments to overwrite the message
//  assert.Equal(t, 1, 2, "numbers dont match for %q", "my param")
//
//  Output:
//  unit_test.go:42: numbers dont match for "my param":
// 		Expected: 1
// 		  Actual: 2
//
// See https://github.com/jgroeneveld/trial for more examples.
//
// Also see https://github.com/jgroeneveld/schema for easier JSON Schema testing.
package assert

import (
	"fmt"
	"github.com/jgroeneveld/trial/th"
	"reflect"
)

// Equal checks if two values are equal.
// Reports if types are different, even though values "look" the same.
func Equal(t testingT, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(t, "Not equal", 1, expected, actual, msgf...)
	}
}

// MustBeEqual checks if two values are equal.
// Reports if types are different, even though values "look" the same.
// Will FailNow if expectation is not met.
func MustBeEqual(t testingT, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected != actual {
		comparisonError(t, "Not equal", 1, expected, actual, msgf...)
		t.FailNow()
	}
}

// NotEqual checks if two values are not equal.
func NotEqual(t testingT, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(t, 1, msg)
	}
}

// MustNotBeEqual checks if two values are equal.
// Will FailNow if expectation is not met.
func MustNotBeEqual(t testingT, expected interface{}, actual interface{}, msgf ...interface{}) {
	if expected == actual {
		msg := titleOrMsgf("Is equal", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, actual)
		th.Error(t, 1, msg)
		t.FailNow()
	}
}

// DeepEqual test if two values are deeply equal.
func DeepEqual(t testingT, expected interface{}, actual interface{}, msgf ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		comparisonError(t, "Not deep equal", 1, expected, actual, msgf...)
	}
}

// MustBeDeepEqual test if two values are deeply equal.
// Will FailNow if expectation is not met.
func MustBeDeepEqual(t testingT, expected interface{}, actual interface{}, msgf ...interface{}) {
	if !reflect.DeepEqual(expected, actual) {
		comparisonError(t, "Not deep equal", 1, expected, actual, msgf...)
		t.FailNow()
	}
}

// True tests if a value is true
func True(t testingT, expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(t, 1, titleOrMsgf("Not true", msgf))
	}
}

// MustBeTrue tests if a value is true
// Will FailNow if expectation is not met.
func MustBeTrue(t testingT, expression bool, msgf ...interface{}) {
	if !expression {
		th.Error(t, 1, titleOrMsgf("Not true", msgf))
		t.FailNow()
	}
}

// False tests if a value is False
func False(t testingT, expression bool, msgf ...interface{}) {
	if expression {
		th.Error(t, 1, titleOrMsgf("Not false", msgf))
	}
}

// MustBeFalse tests if a value is False
// Will FailNow if expectation is not met.
func MustBeFalse(t testingT, expression bool, msgf ...interface{}) {
	if expression {
		th.Error(t, 1, titleOrMsgf("Not false", msgf))
		t.FailNow()
	}
}

// Nil tests if a value is Nil
func Nil(t testingT, expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(t, 1, msg)
	}
}

// MustBeNil tests if a value is Nil
// Will FailNow if expectation is not met.
func MustBeNil(t testingT, expression interface{}, msgf ...interface{}) {
	if expression != nil {
		msg := titleOrMsgf("Not nil", msgf)
		msg += fmt.Sprintf("\n%s: %#v", msgActual, expression)
		th.Error(t, 1, msg)
		t.FailNow()
	}
}

// NotNil tests if a value is Not Nil
func NotNil(t testingT, expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(t, 1, titleOrMsgf("Is nil", msgf))
	}
}

// MustNotBeNil tests if a value is Not Nil
// Will FailNow if expectation is not met.
func MustNotBeNil(t testingT, expression interface{}, msgf ...interface{}) {
	if expression == nil {
		th.Error(t, 1, titleOrMsgf("Is nil", msgf))
		t.FailNow()
	}
}
