package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// This packet is used for a number of actions and
// animations performed by blocks, usually non-persistent.
// The client ignores the provided block type and instead
// uses the block state in their world.
//
// Note that:
//   - This packet uses a block ID from the minecraft:block registry,
//     not a block state.
//
// See Block Actions for a list of values.
//   - Block Actions (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Block_Actions)
type BlockAction struct {
	// Block coordinates.
	Location encoding.BlockPos
	// Varies depending on block — see Block Actions.
	//		- Block Actions (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Block_Actions)
	ActionID byte
	// Varies depending on block — see Block Actions.
	// 		- Block Actions (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Block_Actions)
	ActionParameter byte
	// The block type ID for the block.
	// This value is unused by the vanilla client,
	// as it will infer the type of block based on
	// the given position.
	BlockType int32
}

// ID ..
func (p *BlockAction) ID() int32 {
	return IDClientBoundBlockAction
}

// Resource ..
func (p *BlockAction) Resource() string {
	return "block_event"
}

// BoundType ..
func (p *BlockAction) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *BlockAction) Marshal(io encoding.IO) {
	io.Position(&p.Location)
	io.Uint8(&p.ActionID)
	io.Uint8(&p.ActionParameter)
	io.Varint32(&p.BlockType)
}
