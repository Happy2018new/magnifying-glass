package packet_login

import (
	"magnifying-glass/minecraft/protocol/encoding"
	packet_interface "magnifying-glass/minecraft/protocol/packet/interface"
)

// Enables compression.
//
// If compression is enabled, all following packets are
// encoded in the compressed packet format.
//
// Negative values will disable compression, meaning the packet
// format should remain in the uncompressed packet format.
//
// However, this packet is entirely optional, and if not sent,
// compression will also not be enabled
// (the vanilla server does not send the packet when compression
// is disabled).
//
// For more information, see the following links.
//   - Compressed packet format (https://minecraft.wiki/w/Java_Edition_protocol#With_compression)
//   - Uncompressed packet format (https://minecraft.wiki/w/Java_Edition_protocol#Without_compression)
type SetCompression struct {
	// Maximum size of a packet before it is compressed.
	Threshold int32
}

// ID ..
func (p *SetCompression) ID() int32 {
	return IDClientBoundSetCompression
}

// Resource ..
func (p *SetCompression) Resource() string {
	return "login_compression"
}

// BoundType ..
func (p *SetCompression) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SetCompression) Marshal(io encoding.IO) {
	io.Varint32(&p.Threshold)
}
