package example

import (
	"github.com/jgroeneveld/go-test/assert"
	"testing"
)

func TestExample(t *testing.T) {

	assert.Equal(t, 1, 2)

	assert.Equal(t, 1, "1")

	assert.Equal(t, 42, int64(42))

	assert.True(t, true == false)

	assert.Equal(t, "Harald der Franz", "Peter der Keks")

	assert.MustEqual(t, "nicht", "gleich")

	assert.Equal(t, "never", "thesame - but should not be called")

}
