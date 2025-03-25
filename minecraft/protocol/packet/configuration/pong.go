package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Response to the clientbound packet
// [Ping (https://minecraft.wiki/w/Java_Edition_protocol#Ping_(configuration))]
// with the same id.
type Pong struct {
	// ID ..
	PingID int32
}

// ID ..
func (p *Pong) ID() int32 {
	return IDServerBoundPong
}

// Resource ..
func (p *Pong) Resource() string {
	return "pong"
}

// BoundType ..
func (p *Pong) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *Pong) Marshal(io encoding.IO) {
	io.Int32(&p.PingID)
}
