package ladokmocks

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMockSuperFeed(t *testing.T) {
	assert.Equal(t, 10, MockSuperFeed(10).ID)
}
