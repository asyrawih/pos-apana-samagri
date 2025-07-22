package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSample(t *testing.T) {
	// This is a sample test
	t.Run("Sample test", func(t *testing.T) {
		result := 1 + 1
		expected := 2
		assert.Equal(t, expected, result, "1+1 should be 2")
	})
}
