package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Marks the start of a chunk batch.
// The vanilla client marks and stores
// the time it receives this packet.
type ChunkBiomes struct {
	// See ChunkBiomeData for more information.
	ChunkBiomeData []encoding.ChunkBiomeData
}

// ID ..
func (p *ChunkBiomes) ID() int32 {
	return IDClientBoundChunkBiomes
}

// Resource ..
func (p *ChunkBiomes) Resource() string {
	return "chunks_biomes"
}

// BoundType ..
func (p *ChunkBiomes) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ChunkBiomes) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.ChunkBiomeData)
}
