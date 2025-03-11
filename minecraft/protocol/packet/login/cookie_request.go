package packet_login

import (
	"magnifying-glass/minecraft/protocol/encoding"
	packet_interface "magnifying-glass/minecraft/protocol/packet/interface"
)

// Requests a cookie that was previously stored.
type CookieRequest struct {
	// The identifier of the cookie.
	Key encoding.Identifier
}

// ID ..
func (p *CookieRequest) ID() int32 {
	return IDClientBoundCookieRequest
}

// Resource ..
func (p *CookieRequest) Resource() string {
	return "cookie_request"
}

// BoundType ..
func (p *CookieRequest) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *CookieRequest) Marshal(io encoding.IO) {
	io.Identifier(&p.Key)
}
