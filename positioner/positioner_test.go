package positioner

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultPositioner_GetLinearDistance(t *testing.T) {
	t.Run("Given a negative coords", func(t *testing.T) {
		// arrange
		positioner := NewPositionerDefault()
		expectedLinearDistance := 2.8284271247461903
		// act
		linearDistance := positioner.GetLinearDistance(&Position{X: -1, Y: -2, Z: -3}, &Position{X: -3, Y: -2, Z: -1})

		// assert
		require.Equal(t, expectedLinearDistance, linearDistance)
	})
	t.Run("Given a positive coords", func(t *testing.T) {
		// arrange
		positioner := NewPositionerDefault()
		expectedLinearDistance := 2.8284271247461903
		// act
		linearDistance := positioner.GetLinearDistance(&Position{X: 3, Y: 2, Z: 1}, &Position{X: 1, Y: 2, Z: 3})

		// assert
		require.Equal(t, expectedLinearDistance, linearDistance)
	})
	t.Run("Coords returns linear distance without decimals", func(t *testing.T) {
		// arrange
		positioner := NewPositionerDefault()
		expectedLinearDistance := 6.0
		// act
		linearDistance := positioner.GetLinearDistance(&Position{X: 3, Y: 6, Z: 7}, &Position{X: 1, Y: 2, Z: 3})

		// assert
		require.Equal(t, expectedLinearDistance, linearDistance)
	})
}
