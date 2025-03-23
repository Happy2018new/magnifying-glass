package packet_login

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Disconnect (login)
type Disconnect struct {
	// The reason why the player was disconnected.
	Reason encoding.JsonTextComponent
}

// ID ..
func (p *Disconnect) ID() int32 {
	return IDClientBoundDisconnect
}

// Resource ..
func (p *Disconnect) Resource() string {
	return "login_disconnect"
}

// BoundType ..
func (p *Disconnect) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *Disconnect) Marshal(io encoding.IO) {
	io.JsonTextComponent(&p.Reason)
}
