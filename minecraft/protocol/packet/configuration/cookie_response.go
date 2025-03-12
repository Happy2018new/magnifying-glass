package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Response to a
// Cookie Request (configuration) (https://minecraft.wiki/w/Java_Edition_protocol#Cookie_Request_(configuration))
// from the server.
//
// The vanilla server only accepts
// responses of up to 5 kiB in size.
type CookieResponse struct {
	// The identifier of the cookie.
	Key encoding.Identifier
	// The data of the cookie.
	Payload encoding.Optional[[]byte]
}

// ID ..
func (p *CookieResponse) ID() int32 {
	return IDServerBoundCookieResponse
}

// Resource ..
func (p *CookieResponse) Resource() string {
	return "cookie_response"
}

// BoundType ..
func (p *CookieResponse) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *CookieResponse) Marshal(io encoding.IO) {
	io.Identifier(&p.Key)
	encoding.OptionalSlice(io, &p.Payload, io.Uint8)
}
