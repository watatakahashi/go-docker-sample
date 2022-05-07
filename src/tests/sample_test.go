package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func sample(a int) int {
	return a + 1
}
func TestSample(t *testing.T) {
	t.Run("テストサンプル", func(t *testing.T) {
		assert.Equal(t, 2, sample(1))
	})
}
