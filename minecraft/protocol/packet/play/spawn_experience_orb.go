package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Spawns one or more experience orbs.
type SpawnExperienceOrb struct {
	// EntityID ..
	EntityID int32
	// X ..
	X float64
	// Y ..
	Y float64
	// Z ..
	Z float64
	// The amount of experience this
	// orb will reward once collected.
	Count int16
}

// ID ..
func (p *SpawnExperienceOrb) ID() int32 {
	return IDClientBoundSpawnExperienceOrb
}

// Resource ..
func (p *SpawnExperienceOrb) Resource() string {
	return "add_experience_orb"
}

// BoundType ..
func (p *SpawnExperienceOrb) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SpawnExperienceOrb) Marshal(io encoding.IO) {
	io.Float64(&p.X)
	io.Float64(&p.Y)
	io.Float64(&p.Z)
	io.Int16(&p.Count)
}
