package prey

import "testdoubles/positioner"

// NewPreyStub returns a new PreyStub instance
func NewPreyStub() (prey *PreyStub) {
	prey = &PreyStub{}
	return
}

// PreyStub is a struct that represents a stub prey
type PreyStub struct {
	// FuncGetSpeed is a function that returns the speed of the prey (in m/s)
	FuncGetSpeed func() (speed float64)
	// FuncGetPosition is a function that returns the position of the prey
	FuncGetPosition func() (position *positioner.Position)
}

// GetSpeed returns the speed of the prey (in m/s)
func (p *PreyStub) GetSpeed() (speed float64) {
	speed = p.FuncGetSpeed()
	return
}

// GetPosition returns the position of the prey
func (p *PreyStub) GetPosition() (position *positioner.Position) {
	position = p.FuncGetPosition()
	return
}
