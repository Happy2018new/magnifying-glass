package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Informs the client of which data packs are
// present on the server.
//
// The client is expected to respond with its own
// Serverbound Known Packs (https://minecraft.wiki/w/Java_Edition_protocol#Serverbound_Known_Packs)
// packet.
//
// The vanilla server does not continue with Configuration
// until it receives a response.
//
// The vanilla client requires the minecraft:core pack
// with version 1.21.4 for a normal login sequence.
// This packet must be sent before the Registry Data packets.
type ClientboundKnownPacks struct {
	// See ResourcePack for more information.
	KnownPacks []encoding.ResourcePack
}

// ID ..
func (p *ClientboundKnownPacks) ID() int32 {
	return IDClientBoundKnownPacks
}

// Resource ..
func (p *ClientboundKnownPacks) Resource() string {
	return "select_known_packs"
}

// BoundType ..
func (p *ClientboundKnownPacks) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ClientboundKnownPacks) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.KnownPacks)
}
