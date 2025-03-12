package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Packet is not used by the vanilla server.
// When sent to the client, client responds
// with a Pong packet with the same id.
//
// See the following links for more details.
//   - Pong (https://minecraft.wiki/w/Java_Edition_protocol#Pong_(configuration))
type Ping struct {
	// PingID ..
	PingID int32
}

// ID ..
func (p *Ping) ID() int32 {
	return IDClientBoundPing
}

// Resource ..
func (p *Ping) Resource() string {
	return "ping"
}

// BoundType ..
func (p *Ping) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *Ping) Marshal(io encoding.IO) {
	io.Int32(&p.PingID)
}
