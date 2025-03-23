package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sent by the server before it disconnects a client.
// The client assumes that the server has already closed
// the connection by the time the packet arrives.
type Disconnect struct {
	// Displayed to the client when the connection terminates.
	Reason encoding.TextComponent
}

// ID ..
func (p *Disconnect) ID() int32 {
	return IDClientBoundDisconnect
}

// Resource ..
func (p *Disconnect) Resource() string {
	return "disconnect"
}

// BoundType ..
func (p *Disconnect) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *Disconnect) Marshal(io encoding.IO) {
	io.TextComponent(&p.Reason)
}
