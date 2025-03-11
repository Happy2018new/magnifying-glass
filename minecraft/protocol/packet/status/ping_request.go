package packet_status

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// PingRequest ..
type PingRequest struct {
	// May be any number, but vanilla clients use
	// will always use the timestamp in milliseconds.
	Timestamp int64
}

// ID ..
func (p *PingRequest) ID() int32 {
	return IDServerBoundPingRequest
}

// Resource ..
func (p *PingRequest) Resource() string {
	return "ping_request"
}

// BoundType ..
func (p *PingRequest) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *PingRequest) Marshal(io encoding.IO) {
	io.Int64(&p.Timestamp)
}
