package main_test

import (
	main "go_tutorials/go_channels_tutorial"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquareCubeAddition(t *testing.T) {
	t.Run("10 squared + 10 cubed is 1100", func(t *testing.T) {
		result := main.SquareCubeAddition(10)
		assert.Equal(t, 1100, result)
	})

	t.Run("5 squared + 5 cubed is 1100", func(t *testing.T) {
		result := main.SquareCubeAddition(5)
		assert.Equal(t, 150, result)
	})
}
