package demos

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestGcb(t *testing.T) {
	assert.Equal(t, 3, Gcb(6, 9))
}
