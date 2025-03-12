package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// ResetChat ..
type ResetChat struct{}

// ID ..
func (p *ResetChat) ID() int32 {
	return IDClientBoundResetChat
}

// Resource ..
func (p *ResetChat) Resource() string {
	return "reset_chat"
}

// BoundType ..
func (p *ResetChat) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ResetChat) Marshal(io encoding.IO) {}
