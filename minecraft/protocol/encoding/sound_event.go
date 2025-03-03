package encoding

// Describes a sound that can be played.
type SoundEvent struct {
	// SoundName ..
	SoundName Identifier
	// Whether this sound has a fixed range,
	// as opposed to a variable volume based on distance.
	HasFixedRange bool
	// The maximum range of the sound.
	// Only present if Has Fixed Range is true.
	FixedRange float32
}

func (s *SoundEvent) Marshal(io IO) {
	Single(io, &s.SoundName)
	io.Bool(&s.HasFixedRange)
	if s.HasFixedRange {
		io.Float32(&s.FixedRange)
	}
}
