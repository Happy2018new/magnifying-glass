package packet_status

import (
	"magnifying-glass/minecraft/protocol/encoding"
	packet_interface "magnifying-glass/minecraft/protocol/packet/interface"
)

// The status can only be requested once immediately
// after the handshake, before any ping.
// The server won't respond otherwise.
type StatusRequest struct{}

// ID ..
func (p *StatusRequest) ID() int32 {
	return IDServerBoundStatusRequest
}

// Resource ..
func (p *StatusRequest) Resource() string {
	return "status_request"
}

// BoundType ..
func (p *StatusRequest) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *StatusRequest) Marshal(io encoding.IO) {}
