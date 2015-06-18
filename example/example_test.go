package example

import (
	"fmt"
	"github.com/jgroeneveld/go-test/assert"
	"github.com/jgroeneveld/go-test/th"
	"testing"
)

func TestExample(t *testing.T) {

	if true {
		// generic t.Error replacement that allows skipping
		th.Error(t, "This will happen", 0)
	}

	// simple equals
	assert.Equal(t, 1, 2)

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
