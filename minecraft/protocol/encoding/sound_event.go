package encoding

// Describes a sound that can be played.
type SoundEvent struct {
	// SoundName ..
	SoundName Identifier
	// If this sound has a fixed range,
	// then it refer to the maximum range of the sound.
	// If this field set to not exist,
	// this is a variable volume based on distance.
	FixedRange Optional[float32]
}
