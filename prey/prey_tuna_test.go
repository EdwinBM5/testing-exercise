package prey_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPreyTuna(t *testing.T) {
	t.Run("Tuna speed and position default", func(t *testing.T) {
		// arrange
		pr := prey.CreateTuna()
		ps := positioner.NewPositionerDefault()
		sm := simulator.NewCatchSimulatorDefault(20.0, ps)
		wh := hunter.NewWhiteShark(20, &positioner.Position{
			X: 2,
			Y: 2,
			Z: 2,
		}, sm)
		// act
		err := wh.Hunt(pr)

		// assert
		expectError := hunter.ErrCanNotHunt
		require.ErrorIs(t, err, expectError)
	})

	t.Run("Tuna speed and position custom", func(t *testing.T) {
		// arrange
		pr := prey.NewTuna(10, &positioner.Position{
			X: 5,
			Y: 5,
			Z: 5,
		})

		ps := positioner.NewPositionerDefault()
		sm := simulator.NewCatchSimulatorDefault(20.0, ps)
		wh := hunter.NewWhiteShark(30, &positioner.Position{
			X: 1,
			Y: 1,
			Z: 1,
		}, sm)
		// act
		err := wh.Hunt(pr)

		// assert
		expectError := error(nil)
		require.ErrorIs(t, err, expectError)
	})
}
