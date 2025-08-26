package tuktuk

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	t.Run("should return 0 when input is 0", func(t *testing.T) {
		assert.Equal(t, 0.0, distance(0.0))
	})

	t.Run("should return 0.5 when input is 0.1", func(t *testing.T) {
		assert.Equal(t, 0.5, distance(0.1))
	})

	t.Run("should return 0.5 when input is 0.5", func(t *testing.T) {
		assert.Equal(t, 0.5, distance(0.5))
	})

	t.Run("should return 1 when input is 0.6", func(t *testing.T) {
		assert.Equal(t, 1.0, distance(0.6))
	})

	t.Run("should return 1 when input is 1.0", func(t *testing.T) {
		assert.Equal(t, 1.0, distance(1.0))
	})

	t.Run("should return 1.5 when input is 1.2", func(t *testing.T) {
		assert.Equal(t, 1.5, distance(1.2))
	})

	t.Run("should return 20.0 when input is 19.9", func(t *testing.T) {
		assert.Equal(t, 20.0, distance(19.9))
	})
}

func TestWaitTime(t *testing.T) {
	t.Run("should return 0 when input is 0", func(t *testing.T) {
		assert.Equal(t, 0.0, waitTime(0.0))
	})

	t.Run("should return 1 when input is 0.1", func(t *testing.T) {
		assert.Equal(t, 1.0, waitTime(0.1))
	})

	t.Run("should return 1 when input is 0.5", func(t *testing.T) {
		assert.Equal(t, 1.0, waitTime(0.5))
	})

	t.Run("should return 1 when input is 0.6", func(t *testing.T) {
		assert.Equal(t, 1.0, waitTime(0.6))
	})

	t.Run("should return 1.0 when input is 1.0", func(t *testing.T) {
		assert.Equal(t, 1.0, waitTime(1.0))
	})

	t.Run("should return 2 when input is 1.2", func(t *testing.T) {
		assert.Equal(t, 2.0, waitTime(1.2))
	})
}
