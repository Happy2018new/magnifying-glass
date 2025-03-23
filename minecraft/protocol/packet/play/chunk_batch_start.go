package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Marks the start of a chunk batch.
// The vanilla client marks and stores
// the time it receives this packet.
type ChunkBatchStart struct{}

// ID ..
func (p *ChunkBatchStart) ID() int32 {
	return IDClientBoundChunkBatchStart
}

// Resource ..
func (p *ChunkBatchStart) Resource() string {
	return "chunk_batch_start"
}

// BoundType ..
func (p *ChunkBatchStart) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ChunkBatchStart) Marshal(io encoding.IO) {}
