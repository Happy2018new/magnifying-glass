package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Stores some arbitrary data on the client,
// which persists between server transfers.
// The vanilla client only accepts cookies
// of up to 5 kiB in size.
type StoreCookie struct {
	// The identifier of the cookie.
	Key encoding.Identifier
	// The data of the cookie.
	Payload []byte
}

// ID ..
func (p *StoreCookie) ID() int32 {
	return IDClientBoundStoreCookie
}

// Resource ..
func (p *StoreCookie) Resource() string {
	return "store_cookie"
}

// BoundType ..
func (p *StoreCookie) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *StoreCookie) Marshal(io encoding.IO) {
	io.Identifier(&p.Key)
	encoding.FuncSliceVarint32Length(io, &p.Payload, io.Uint8)
}
