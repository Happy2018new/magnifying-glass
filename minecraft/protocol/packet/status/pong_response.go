package packet_status

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// PongResponse ..
type PongResponse struct {
	// Should match the one sent by the client.
	Timestamp int64
}

// ID ..
func (p *PongResponse) ID() int32 {
	return IDClientBoundPongResponse
}

// Resource ..
func (p *PongResponse) Resource() string {
	return "pong_response"
}

// BoundType ..
func (p *PongResponse) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *PongResponse) Marshal(io encoding.IO) {
	io.Int64(&p.Timestamp)
}
