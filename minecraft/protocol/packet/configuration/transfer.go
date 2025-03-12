package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Notifies the client that it should
// transfer to the given server.
// Cookies previously stored are preserved
// between server transfers.
type Transfer struct {
	// The hostname or IP of the server.
	Host string
	// The port of the server.
	Port int32
}

// ID ..
func (p *Transfer) ID() int32 {
	return IDClientBoundTransfer
}

// Resource ..
func (p *Transfer) Resource() string {
	return "transfer"
}

// BoundType ..
func (p *Transfer) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *Transfer) Marshal(io encoding.IO) {
	io.String(&p.Host)
	io.Varint32(&p.Port)
}
