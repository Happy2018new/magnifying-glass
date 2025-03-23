package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

const (
	PlayerAnimationSwingMainArm uint8 = iota
	PlayerAnimationLeaveBed
	PlayerAnimationCriticalEffect
	PlayerAnimationMagicCriticalEffect
)

// Sent whenever an entity should change animation.
type EntityAnimation struct {
	// Player ID.
	EntityID int32
	// Animation ID (see constant enum above).
	Animation byte
}

// ID ..
func (p *EntityAnimation) ID() int32 {
	return IDClientBoundEntityAnimation
}

// Resource ..
func (p *EntityAnimation) Resource() string {
	return "animate"
}

// BoundType ..
func (p *EntityAnimation) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *EntityAnimation) Marshal(io encoding.IO) {
	io.Varint32(&p.EntityID)
	io.Uint8(&p.Animation)
}
