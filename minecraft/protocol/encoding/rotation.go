package encoding

// Rotation ..
type Rotation [3]float32

// X returns the rotation on X.
func (e Rotation) X() float32 {
	return e[0]
}

// Z returns the rotation on Y.
func (e Rotation) Y() float32 {
	return e[1]
}

// Z returns the rotation on Z.
func (e Rotation) Z() float32 {
	return e[2]
}
