package packet_handshaking

import (
	"magnifying-glass/minecraft/protocol/encoding"
	packet_interface "magnifying-glass/minecraft/protocol/packet/interface"
)

// While not technically part of the current protocol,
// legacy clients may send this packet to initiate Server List Ping,
// and modern servers should handle it correctly.
//
// The format of this packet is a remnant of the pre-Netty age,
// before the switch to Netty in 1.7 brought the standard format
// that is recognized now.
// This packet merely exists to inform legacy clients that they
// can't join our modern server.
//
// For Server List Ping, see the following links for details.
// (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Server_List_Ping).
//
// See Server List Ping#1.6 (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Server_List_Ping#1.6)
// for the details of the protocol that follows this packet.
type LegacyServerListPing struct {
	// always 1 (0x01).
	Payload byte
}

// ID ..
func (p *LegacyServerListPing) ID() int32 {
	return IDServerBoundLegacyServerListPing
}

// Resource ..
func (p *LegacyServerListPing) Resource() string {
	return "not-implement"
}

// BoundType ..
func (p *LegacyServerListPing) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *LegacyServerListPing) Marshal(io encoding.IO) {
	io.Uint8(&p.Payload)
}
