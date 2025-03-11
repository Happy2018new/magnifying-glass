package packet_login

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// LoginPluginResponse ..
type LoginPluginResponse struct {
	// Should match ID from server.
	MessageID int32
	// Any data, depending on the channel.
	// The length of this array must be inferred from the packet length.
	// Only present if the client understood the request.
	Data encoding.Optional[[]byte]
}

// ID ..
func (p *LoginPluginResponse) ID() int32 {
	return IDServerBoundLoginPluginResponse
}

// Resource ..
func (p *LoginPluginResponse) Resource() string {
	return "custom_query_answer"
}

// BoundType ..
func (p *LoginPluginResponse) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *LoginPluginResponse) Marshal(io encoding.IO) {
	io.Varint32(&p.MessageID)
	encoding.OptionalFunc(io, &p.Data, io.Bytes)
}
