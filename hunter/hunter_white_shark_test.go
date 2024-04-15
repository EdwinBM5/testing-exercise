package hunter

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testdoubles/simulator"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestWhiteSharkHunt(t *testing.T) {
	t.Run("white shark hunts a prey - prey short distance, low speed", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.FuncGetSpeed = func() (speed float64) {
			speed = 10
			return
		}
		pr.FuncGetPosition = func() (position *positioner.Position) {
			position = &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}
			return
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.FuncCanCatch = func(hunter, prey *simulator.Subject) (canCatch bool) {
			canCatch = true
			return
		}

		wh := NewWhiteShark(20, &positioner.Position{
			X: 2,
			Y: 2,
			Z: 2,
		}, sm)

		// act
		err := wh.Hunt(pr)

		// assert
		expectError := error(nil)
		expectedMockCallCanCatch := 1
		require.ErrorIs(t, err, expectError)
		require.Equal(t, expectedMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunts a prey - shark is slow", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.FuncGetSpeed = func() (speed float64) {
			speed = 20
			return
		}
		pr.FuncGetPosition = func() (position *positioner.Position) {
			position = &positioner.Position{
				X: 0,
				Y: 0,
				Z: 0,
			}
			return
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.FuncCanCatch = func(hunter, prey *simulator.Subject) (canCatch bool) {
			canCatch = false
			return
		}

		wh := NewWhiteShark(5, &positioner.Position{
			X: 2,
			Y: 2,
			Z: 2,
		}, sm)

		// act
		err := wh.Hunt(pr)

		// assert
		expectError := ErrCanNotHunt
		expectedMockCallCanCatch := 1
		require.ErrorIs(t, err, expectError)
		require.Equal(t, expectedMockCallCanCatch, sm.Calls.CanCatch)
	})

	t.Run("white shark can not hunts a prey - prey is too far", func(t *testing.T) {
		// arrange
		pr := prey.NewPreyStub()
		pr.FuncGetSpeed = func() (speed float64) {
			speed = 20
			return
		}
		pr.FuncGetPosition = func() (position *positioner.Position) {
			position = &positioner.Position{
				X: 70,
				Y: 70,
				Z: 70,
			}
			return
		}

		sm := simulator.NewCatchSimulatorMock()
		sm.FuncCanCatch = func(hunter, prey *simulator.Subject) (canCatch bool) {
			canCatch = false
			return
		}

		wh := NewWhiteShark(10, &positioner.Position{
			X: 2,
			Y: 2,
			Z: 2,
		}, sm)

		// act
		err := wh.Hunt(pr)

		// assert
		expectError := ErrCanNotHunt
		expectedMockCallCanCatch := 1
		require.ErrorIs(t, err, expectError)
		require.Equal(t, expectedMockCallCanCatch, sm.Calls.CanCatch)
	})
}
