package packet_login

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Acknowledgement to the
// Login Success (https://minecraft.wiki/w/Java_Edition_protocol#Login_Success)
// packet sent by the server.
type LoginAcknowledged struct{}

// ID ..
func (p *LoginAcknowledged) ID() int32 {
	return IDServerBoundLoginAcknowledged
}

// Resource ..
func (p *LoginAcknowledged) Resource() string {
	return "login_acknowledged"
}

// BoundType ..
func (p *LoginAcknowledged) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *LoginAcknowledged) Marshal(io encoding.IO) {}
