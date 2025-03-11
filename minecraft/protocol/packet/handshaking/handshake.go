package packet_handshaking

import (
	"magnifying-glass/minecraft/protocol/encoding"
	packet_interface "magnifying-glass/minecraft/protocol/packet/interface"
)

const (
	HandShakeNextStateStatus int32 = iota + 1
	HandShakeNextStateLogin
	HandShakeNextStateTransfer
)

// This packet causes the server to switch into the target state,
// it should be sent right after opening the TCP connection to
// avoid the server from disconnecting.
type HandShake struct {
	// See protocol version numbers
	// (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol_version_numbers).
	ProtocolVersion int32
	// Hostname or IP, e.g. localhost or 127.0.0.1,
	// that was used to connect.
	//
	// The vanilla server does not use this information.
	// Note that SRV records are a simple redirect,
	//
	// e.g.
	// 		- if _minecraft._tcp.example.com points to mc.example.org,
	// 		  users connecting to example.com will provide example.org
	// 		  as server address in addition to connecting to it.
	ServerAddress string
	// Default is 25565.
	// The vanilla server does not use this information.
	ServerPort uint16
	// 1 for Status (https://minecraft.wiki/w/Java_Edition_protocol#Status),
	// 2 for Login (https://minecraft.wiki/w/Java_Edition_protocol#Login),
	// 3 for Transfer (https://minecraft.wiki/w/Java_Edition_protocol#Login).
	NextState int32
}

// ID ..
func (p *HandShake) ID() int32 {
	return IDServerBoundHandshaking
}

// Resource ..
func (p *HandShake) Resource() string {
	return "intention"
}

// BoundType ..
func (p *HandShake) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *HandShake) Marshal(io encoding.IO) {
	io.Varint32(&p.ProtocolVersion)
	io.String(&p.ServerAddress)
	io.Uint16(&p.ServerPort)
	io.Varint32(&p.NextState)
}
