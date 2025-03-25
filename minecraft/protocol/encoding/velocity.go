package encoding

// EntityVelocity is the velocity of a entity.
// It is composed of one double and two float,
// and is written as 1 float64 and 2 float32s.
type EntityVelocity struct {
	// X velocity of this entity.
	VelocityX float64
	// Y velocity of this entity.
	VelocityY float32
	// Z velocity of this entity.
	VelocityZ float32
}

// X returns the X velocity of this entity.
// It is equivalent to e.VelocityX.
func (e EntityVelocity) X() float64 {
	return e.VelocityX
}

// Y returns the Y velocity of this entity.
// It is equivalent to e.VelocityY.
func (e EntityVelocity) Y() float32 {
	return e.VelocityY
}

// Z returns the Z velocity of this entity.
// It is equivalent to e.VelocityZ.
func (e EntityVelocity) Z() float32 {
	return e.VelocityZ
}

func (e *EntityVelocity) Marshal(io IO) {
	io.Float64(&e.VelocityX)
	io.Float32(&e.VelocityY)
	io.Float32(&e.VelocityZ)
}
