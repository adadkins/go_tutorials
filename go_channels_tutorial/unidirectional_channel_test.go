package main_test

import (
	main "go_tutorials/go_channels_tutorial"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnidirectionalChannel(t *testing.T) {
	t.Run("Unidirectional Channel test", func(t *testing.T) {
		result := main.UnidirectionalChannel(10)
		assert.Equal(t, 1100, result)
	})

	t.Run("5 squared + 5 cubed is 1100", func(t *testing.T) {
		result := main.UnidirectionalChannel(5)
		assert.Equal(t, 150, result)
	})
}
