package encoding

import "github.com/Happy2018new/magnifying-glass/minecraft/nbt"

// ChunkBlockEntity refer to the NBT data
// of a Minecraft Java block entity that
// record in ChunkData.
type ChunkBlockEntity struct {
	// The packed section coordinates are
	// relative to the chunk they are in.
	// Values 0-15 are valid.
	//
	// See the following code for details.
	//		```
	// 		packed_xz = ((blockX & 15) << 4) | (blockZ & 15) // encode
	// 		x = packed_xz >> 4, z = packed_xz & 15 // decode
	//		```
	PackedXZ byte
	// The height relative to the world.
	Y int16
	// The type of block entity.
	Type int32
	// The block entity's data,
	// without the X, Y, and Z values.
	Data map[string]any
}

// GetRelativePos return the real relative pos
// of this block entity that relative to the
// chunk they are in.
func (c ChunkBlockEntity) GetRelativePos() BlockPos {
	return [3]int32{
		int32(c.PackedXZ >> 4),
		int32(c.Y),
		int32(c.PackedXZ & 15),
	}
}

// SetRelativePos set the relative pos
// of this block entity.
// This pos is a relative pos that relative
// to the chunk they are in.
func (c *ChunkBlockEntity) SetRelativePos(blockPos BlockPos) {
	c.PackedXZ = byte(
		((blockPos[0] & 15) << 4) | (blockPos[2] & 15),
	)
	c.Y = int16(blockPos[1])
}

func (c *ChunkBlockEntity) Marshal(io IO) {
	io.Uint8(&c.PackedXZ)
	io.Int16(&c.Y)
	io.Varint32(&c.Type)
	io.NBT(&c.Data, nbt.NetworkBigEndian)
}

// ChunkData represents the chunk data
// passing in Minecraft Java protocol,
// which used to do world generate.
type ChunkData struct {
	// See (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Chunk_Format#Heightmaps_structure).
	Heightmaps map[string]any
	// See (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Chunk_Format#Data_structure).
	Data []byte
	// See ChunkBlockEntity for more information.
	BlockEntities []ChunkBlockEntity
}
