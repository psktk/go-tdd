package tuktuk

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDistance(t *testing.T) {
	cases := []struct {
		input    float64
		expected float64
	}{
		{input: 0.0, expected: 0.0},
		{input: 0.1, expected: 0.5},
		{input: 0.5, expected: 0.5},
		{input: 0.6, expected: 1.0},
		{input: 1.0, expected: 1.0},
		{input: 1.2, expected: 1.5},
		{input: 19.9, expected: 20.0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("should return %.1f when input is %.1f", c.expected, c.input), func(t *testing.T) {
			assert.Equal(t, c.expected, distance(c.input))
		})
	}
}

func TestWaitTime(t *testing.T) {
	cases := []struct {
		input    float64
		expected float64
	}{
		{input: 0.0, expected: 0.0},
		{input: 0.1, expected: 1.0},
		{input: 0.5, expected: 1.0},
		{input: 0.6, expected: 1.0},
		{input: 1.0, expected: 1.0},
		{input: 1.2, expected: 2.0},
	}

	for _, c := range cases {
		t.Run(fmt.Sprintf("should return %.1f when input is %.1f", c.expected, c.input), func(t *testing.T) {
			assert.Equal(t, c.expected, waitTime(c.input))
		})
	}
}
