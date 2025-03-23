package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Fired whenever a block is changed within the render distance.
//
// Note that:
//   - Changing a block in a chunk that is not loaded is not a
//     stable action.
//     The vanilla client currently uses a shared empty chunk
//     which is modified for all block changes in unloaded chunks;
//     while in 1.9 this chunk never renders in older versions the
//     changed block will appear in all copies of the empty chunk.
//     Servers should avoid sending block changes in unloaded chunks
//     and clients should ignore such packets.
type BlockUpdate struct {
	// Block coordinates.
	Location encoding.BlockPos
	// The new block state ID for the block
	// as given in the block state registry.
	// 		- block state registry (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Chunk_Format#Block_state_registry)
	BlockID int32
}

// ID ..
func (p *BlockUpdate) ID() int32 {
	return IDClientBoundBlockUpdate
}

// Resource ..
func (p *BlockUpdate) Resource() string {
	return "block_update"
}

// BoundType ..
func (p *BlockUpdate) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *BlockUpdate) Marshal(io encoding.IO) {
	io.Position(&p.Location)
	io.Varint32(&p.BlockID)
}
