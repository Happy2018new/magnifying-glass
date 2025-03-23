package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Marks the end of a chunk batch.
//
// The vanilla client marks the time it receives
// this packet and calculates the elapsed duration
// since the beginning of the chunk batch.
//
// The server uses this duration and the batch size
// received in this packet to estimate the number of
// milliseconds elapsed per chunk received.
//
// This value is then used to calculate the desired
// number of chunks per tick through the formula
// 25 / millisPerChunk, which is reported to the server
// through Chunk Batch Received.
//
// This likely uses 25 instead of the normal tick duration
// of 50 so chunk processing will only use half of the
// client's and network's bandwidth.
//
// The vanilla client uses the samples from the latest 15
// batches to estimate the milliseconds per chunk number.
//
// For more information, see the following links.
//   - beginning of the chunk batch (https://minecraft.wiki/w/Java_Edition_protocol#Chunk_Batch_Start)
//   - Chunk Batch Received (https://minecraft.wiki/w/Java_Edition_protocol#Chunk_Batch_Received)
type ChunkBatchFinished struct {
	// Number of chunks.
	BatchSize int32
}

// ID ..
func (p *ChunkBatchFinished) ID() int32 {
	return IDClientBoundChunkBatchFinished
}

// Resource ..
func (p *ChunkBatchFinished) Resource() string {
	return "chunk_batch_finished"
}

// BoundType ..
func (p *ChunkBatchFinished) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ChunkBatchFinished) Marshal(io encoding.IO) {
	io.Varint32(&p.BatchSize)
}
