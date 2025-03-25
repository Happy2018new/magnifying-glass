package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// The server will frequently send out a keep-alive
// [see Clientbound Keep Alive (https://minecraft.wiki/w/Java_Edition_protocol#Clientbound_Keep_Alive_(configuration)) ],
// each containing a random ID.
// The client must respond with the same packet.
type ServerboundKeepAlive struct {
	// KeepAliveID ..
	KeepAliveID int64
}

// ID ..
func (p *ServerboundKeepAlive) ID() int32 {
	return IDServerBoundKeepAlive
}

// Resource ..
func (p *ServerboundKeepAlive) Resource() string {
	return "keep_alive"
}

// BoundType ..
func (p *ServerboundKeepAlive) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *ServerboundKeepAlive) Marshal(io encoding.IO) {
	io.Int64(&p.KeepAliveID)
}
