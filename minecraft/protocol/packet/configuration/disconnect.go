package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Disconnect ..
type Disconnect struct {
	// The reason why the player was disconnected.
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
