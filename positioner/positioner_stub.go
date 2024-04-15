package positioner

// NewPositionerStub returns a new PositionerStub instance
func NewPositionerStub() (positioner *PositionerStub) {
	positioner = &PositionerStub{}
	return
}

// PositionerStub is a struct that represents a stub positioner
type PositionerStub struct {
	// FuncGetLinearDistance is a function that returns the linear distance between 2 positions (in meters)
	FuncGetLinearDistance func(from, to *Position) (linearDistance float64)
}

// GetLinearDistance returns the linear distance between 2 positions (in meters)
func (s *PositionerStub) GetLinearPosition(from, to *Position) (linearDistance float64) {
	linearDistance = s.FuncGetLinearDistance(from, to)
	return
}
