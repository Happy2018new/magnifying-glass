package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// This packet is sent by the server when
// an entity moves more than 8 blocks.
//
// Note that:
//   - The Mojang-specified name of this packet was
//     changed in 1.21.2 from "teleport_entity" to
//     "entity_position_sync".
//   - There is a new "teleport_entity", which this
//     document more appropriately calls
//     Synchronize Vehicle Position (https://minecraft.wiki/w/Java_Edition_protocol#Synchronize_Vehicle_Position).
//   - That packet has a different function and will
//     lead to confusing results if used in place of this one.
type TeleportEntity struct {
	// EntityID ..
	EntityID int32
	// X ..
	X float64
	// Y ..
	Y float64
	// Z ..
	Z float64
	// VelocityX ..
	VelocityX float64
	// VelocityY ..
	VelocityY float64
	// VelocityZ ..
	VelocityZ float64
	// Rotation on the X axis, in degrees.
	Yaw float32
	// Rotation on the Y axis, in degrees.
	Pitch float32
	// OnGround ..
	OnGround bool
}

// ID ..
func (p *TeleportEntity) ID() int32 {
	return IDClientBoundTeleportEntity
}

// Resource ..
func (p *TeleportEntity) Resource() string {
	return "entity_position_sync"
}

// BoundType ..
func (p *TeleportEntity) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *TeleportEntity) Marshal(io encoding.IO) {
	io.Varint32(&p.EntityID)
	io.Float64(&p.X)
	io.Float64(&p.Y)
	io.Float64(&p.Z)
	io.Float64(&p.VelocityX)
	io.Float64(&p.VelocityY)
	io.Float64(&p.VelocityZ)
	io.Float32(&p.Yaw)
	io.Float32(&p.Pitch)
	io.Bool(&p.OnGround)
}
