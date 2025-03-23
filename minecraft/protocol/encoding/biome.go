package encoding

// ChunkBiomeData is used on Chunk Biomes packet.
// Note that the order of X and Z is inverted,
// because the client reads them as one big-endian Long,
// with Z being the upper 32 bits.
type ChunkBiomeData struct {
	// Chunk coordinate
	// (block coordinate divided by 16, rounded down)
	ChunkZ int32
	// Chunk coordinate
	// (block coordinate divided by 16, rounded down)
	ChunkX int32
	// Chunk data structure (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Chunk_Format#Data_structure),
	// with sections (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Chunk_Format#Chunk_Section)
	// containing only the Biomes field.
	Data []byte
}

func (c *ChunkBiomeData) Marshal(io IO) {
	io.Int32(&c.ChunkZ)
	io.Int32(&c.ChunkX)
	FuncSliceVarint32Length(io, &c.Data, io.Uint8)
}
