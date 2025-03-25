package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Tells the client to unload a chunk column.
//
// Note that:
//
//   - Note: The order is inverted,
//     because the client reads this
//     packet as one big-endian Long,
//     with Z being the upper 32 bits.
//
//   - It is legal to send this packet
//     even if the given chunk is not
//     currently loaded.
type UnloadChunk struct {
	// Block coordinate divided by 16, rounded down.
	ChunkZ int32
	// Block coordinate divided by 16, rounded down.
	ChunkX int32
}

// ID ..
func (p *UnloadChunk) ID() int32 {
	return IDClientBoundUnloadChunk
}

// Resource ..
func (p *UnloadChunk) Resource() string {
	return "forget_level_chunk"
}

// BoundType ..
func (p *UnloadChunk) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *UnloadChunk) Marshal(io encoding.IO) {
	io.Int32(&p.ChunkZ)
	io.Int32(&p.ChunkX)
}
