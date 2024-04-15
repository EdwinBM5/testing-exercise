package simulator

// NewCatchSimulatorMock returns a new CatchSimulatorMock instance
func NewCatchSimulatorMock() (catchSimulator *CatchSimulatorMock) {
	catchSimulator = &CatchSimulatorMock{}
	return
}

// CatchSimulatorMock is a struct that represents a mock catch simulator
type CatchSimulatorMock struct {
	// FuncCanCatch is a function that returns whether a hunter can catch a prey
	FuncCanCatch func(hunter, prey *Subject) (canCatch bool)
	// FuncGetCalls is a function that returns the number of calls to the CanCatch function
	Calls struct {
		CanCatch int
	}
}

// CanCatch returns whether a hunter can catch a prey
func (s *CatchSimulatorMock) CanCatch(hunter, prey *Subject) (canCatch bool) {
	s.Calls.CanCatch++
	canCatch = s.FuncCanCatch(hunter, prey)
	return
}
