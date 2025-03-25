package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sent when an explosion occurs
// (creepers, TNT, and ghast fireballs).
type Explosion struct {
	// X ..
	X float64
	// Y ..
	Y float64
	// Z ..
	Z float64
	// X, Y and Z velocity of the player
	// being pushed by the explosion.
	PlayerVelocity encoding.Optional[encoding.EntityVelocity]
	// Particle ..
	Particle encoding.Particle
	// ID in the minecraft:sound_event
	// registry, or an inline definition.
	ExplosionSound encoding.IDOrX[encoding.SoundEvent]
}

// ID ..
func (p *Explosion) ID() int32 {
	return IDClientBoundExplosion
}

// Resource ..
func (p *Explosion) Resource() string {
	return "explode"
}

// BoundType ..
func (p *Explosion) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *Explosion) Marshal(io encoding.IO) {
	io.Float64(&p.X)
	io.Float64(&p.Y)
	io.Float64(&p.Z)
	encoding.OptionalMarshaler(io, &p.PlayerVelocity)
	io.Particle(&p.Particle)
	encoding.IDOrXFunc(io, &p.ExplosionSound, io.SoundEvent)
}
