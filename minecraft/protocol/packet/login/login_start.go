package packet_login

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"

	"github.com/google/uuid"
)

// LoginStart ..
type LoginStart struct {
	// Player's Username.
	Name string
	// The UUID of the player logging in.
	// Unused by the vanilla server.
	PlayerUUID uuid.UUID
}

// ID ..
func (p *LoginStart) ID() int32 {
	return IDServerBoundLoginStart
}

// Resource ..
func (p *LoginStart) Resource() string {
	return "hello"
}

// BoundType ..
func (p *LoginStart) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *LoginStart) Marshal(io encoding.IO) {
	io.String(&p.Name)
	io.UUID(&p.PlayerUUID)
}
