package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// The server will frequently send out a keep-alive,
// each containing a random ID.
//
// The client must respond with the same payload
// [see Serverbound Keep Alive (https://minecraft.wiki/w/Java_Edition_protocol#Serverbound_Keep_Alive_(configuration))].
//
// If the client does not respond to a Keep Alive
// packet within 15 seconds after it was sent,
// the server kicks the client.
//
// Vice versa, if the server does not send any
// keep-alives for 20 seconds, the client will
// disconnect and yields a "Timed out" exception.
type ClientBoundKeepAlive struct {
	// KeepAliveID ..
	KeepAliveID int64
}

// ID ..
func (p *ClientBoundKeepAlive) ID() int32 {
	return IDClientBoundPluginMessage
}

// Resource ..
func (p *ClientBoundKeepAlive) Resource() string {
	return "keep_alive"
}

// BoundType ..
func (p *ClientBoundKeepAlive) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ClientBoundKeepAlive) Marshal(io encoding.IO) {
	io.Int64(&p.KeepAliveID)
}
