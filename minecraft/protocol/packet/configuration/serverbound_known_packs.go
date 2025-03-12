package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Informs the server of which data packs
// are present on the client.
//
// The client sends this in response to
// Clientbound Known Packs (https://minecraft.wiki/w/Java_Edition_protocol#Clientbound_Known_Packs).
//
// If the client specifies a pack in this packet,
// the server should omit its contained data from
// the Registry Data packet (https://minecraft.wiki/w/Java_Edition_protocol#Registry_Data).
type ServerboundKnownPacks struct {
	// See ResourcePack for more information.
	KnownPacks []encoding.ResourcePack
}

// ID ..
func (p *ServerboundKnownPacks) ID() int32 {
	return IDServerBoundKnownPacks
}

// Resource ..
func (p *ServerboundKnownPacks) Resource() string {
	return "select_known_packs"
}

// BoundType ..
func (p *ServerboundKnownPacks) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *ServerboundKnownPacks) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.KnownPacks)
}
