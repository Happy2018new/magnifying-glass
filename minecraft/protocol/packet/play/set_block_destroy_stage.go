package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// 0–9 are the displayable destroy stages and each other
// number means that there is no animation on this coordinate.
//
// Block break animations can still be applied on air;
// the animation will remain visible although there is
// no block being broken. However, if this is applied to
// a transparent block, odd graphical effects may happen,
// including water losing its transparency.
// (An effect similar to this can be seen in normal gameplay
// when breaking ice blocks)
//
// If you need to display several break animations at the same
// time you have to give each of them a unique Entity ID.
// The entity ID does not need to correspond to an actual entity
// on the client.
// It is valid to use a randomly generated number.
type SetBlockDestroyStage struct {
	// The ID of the entity breaking the block.
	EntityID int32
	// Block Position.
	Location encoding.BlockPos
	// 0–9 to set it, any other value to remove it.
	DestroyStage byte
}

// ID ..
func (p *SetBlockDestroyStage) ID() int32 {
	return IDClientBoundSetBlockDestroyStage
}

// Resource ..
func (p *SetBlockDestroyStage) Resource() string {
	return "block_destruction"
}

// BoundType ..
func (p *SetBlockDestroyStage) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SetBlockDestroyStage) Marshal(io encoding.IO) {
	io.Varint32(&p.EntityID)
	io.Position(&p.Location)
	io.Uint8(&p.DestroyStage)
}
