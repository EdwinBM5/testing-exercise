package hunt

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		// arrange
		shark := NewWhiteShark(true, false, 10.0)
		tuna := NewTuna("Tuna", 5.0)

		// act
		err := shark.Hunt(tuna)

		// assert
		require.NoError(t, err)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		// arrange
		shark := NewWhiteShark(false, true, 10.0)
		tuna := NewTuna("Tuna", 5.0)
		expectedErr := ErrSharkIsNotHungry

		// act
		err := shark.Hunt(tuna)

		// assert
		require.ErrorIs(t, expectedErr, err)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		// arrange
		shark := NewWhiteShark(true, true, 10.0)
		tuna := NewTuna("Tuna", 5.0)
		expectedErr := ErrSharkIsTired

		// act
		err := shark.Hunt(tuna)

		// assert
		require.ErrorIs(t, expectedErr, err)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		// arrange
		shark := NewWhiteShark(true, false, 10.0)
		tuna := NewTuna("Tuna", 15.0)
		expectedErr := ErrSharkIsSlower

		// act
		err := shark.Hunt(tuna)

		// assert
		require.ErrorIs(t, expectedErr, err)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		// arrange
		shark := NewWhiteShark(true, false, 10.0)
		var tuna *Tuna
		expectedErr := ErrPreyIsNil

		// act
		err := shark.Hunt(tuna)

		// assert
		require.ErrorIs(t, expectedErr, err)
	})
}
