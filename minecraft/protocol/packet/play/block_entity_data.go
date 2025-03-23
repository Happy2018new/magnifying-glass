package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/nbt"
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sets the block entity associated with the block at the given location.
type BlockEntityData struct {
	// Location ..
	Location encoding.BlockPos
	// The type of the block entity.
	Type int32
	// Data to set. May be a TAG_END (0),
	// in which case the block entity at
	// the given location is removed
	// (though this is not required since
	// the client will remove the block
	// entity automatically on chunk
	// unload or block removal).
	NBTData map[string]any
}

// ID ..
func (p *BlockEntityData) ID() int32 {
	return IDClientBoundBlockEntityData
}

// Resource ..
func (p *BlockEntityData) Resource() string {
	return "block_entity_data"
}

// BoundType ..
func (p *BlockEntityData) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *BlockEntityData) Marshal(io encoding.IO) {
	io.Position(&p.Location)
	io.Varint32(&p.Type)
	io.NBT(&p.NBTData, nbt.NetworkBigEndian)
}
