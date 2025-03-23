package encoding

// BlockPos is the position of a block. It is composed of three integers, and is typically written as either
// 3 varint32s or a varint32, varuint32 and varint32.
type BlockPos [3]int32

// X returns the X coordinate of the block position.
// It is equivalent to BlockPos[0].
func (pos BlockPos) X() int32 {
	return pos[0]
}

// Y returns the Y coordinate of the block position.
// It is equivalent to BlockPos[1].
func (pos BlockPos) Y() int32 {
	return pos[1]
}

// Z returns the Z coordinate of the block position.
// It is equivalent to BlockPos[2].
func (pos BlockPos) Z() int32 {
	return pos[2]
}

// EntityPos is the position of a entity.
// It is composed of three double, and is
// written as 3 float64s.
type EntityPos [3]float64

// X returns the X coordinate of the entity position.
// It is equivalent to EntityPos[0].
func (pos EntityPos) X() float64 {
	return pos[0]
}

// Y returns the Y coordinate of the entity position.
// It is equivalent to EntityPos[1].
func (pos EntityPos) Y() float64 {
	return pos[1]
}

// Z returns the Z coordinate of the entity position.
// It is equivalent to EntityPos[2].
func (pos EntityPos) Z() float64 {
	return pos[2]
}

// GlobalBlockPos record a block pos
// with the dimension id.
type GlobalBlockPos struct {
	// Dimension ..
	Dimension Identifier
	// Position ..
	Position BlockPos
}

func (g *GlobalBlockPos) Marshal(io IO) {
	io.Identifier(&g.Dimension)
	io.Position(&g.Position)
}
