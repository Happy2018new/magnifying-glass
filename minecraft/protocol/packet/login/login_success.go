package packet_login

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"

	"github.com/google/uuid"
)

// LoginSuccess ..
type LoginSuccess struct {
	// UUID ..
	UUID uuid.UUID
	// Username ..
	Username string
	// See PlayerProfileProperty for more information.
	Property []encoding.PlayerProfileProperty
}

// ID ..
func (p *LoginSuccess) ID() int32 {
	return IDClientBoundLoginSuccess
}

// Resource ..
func (p *LoginSuccess) Resource() string {
	return "login_finished"
}

// BoundType ..
func (p *LoginSuccess) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *LoginSuccess) Marshal(io encoding.IO) {
	io.UUID(&p.UUID)
	io.String(&p.Username)
	encoding.SliceVarint32Length(io, &p.Property)
}
