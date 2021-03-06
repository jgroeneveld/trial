package example

import (
	"fmt"
	"github.com/jgroeneveld/trial/assert"
	"github.com/jgroeneveld/trial/th"
	"testing"
)

func TestExample(t *testing.T) {

	if true {
		// generic t.Error replacement that allows skipping
		th.Error(t, 0, "This will happen")
	}

	// simple equals
	assert.Equal(t, 1, 2, "numbers dont match")

	assert.NotEqual(t, 1, 1, "numbers match")

	// if the types dont match, it will be printed
	assert.Equal(t, 1, "1")

	assert.True(t, true == false)

	assert.False(t, true == true)

	// assert nil to have easy error handling in tests
	err := someError()
	assert.Nil(t, err)

	err = noErr()
	assert.Nil(t, err)

	// Must* functions will call FailNow. (Fatal equivalent)
	assert.MustBeEqual(t, "nicht", "gleich")

	assert.Equal(t, "never", "thesame - but should not be called")
}

func someError() error {
	return fmt.Errorf("Hallo Welt")
}

func noErr() error {
	return nil
}
