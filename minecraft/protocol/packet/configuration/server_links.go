package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// This packet contains a list of links that the vanilla
// client will display in the menu available from the pause menu.
// Link labels can be built-in or custom (i.e., any text).
type ServerLinks struct {
	// See ServerLink for more information.
	Links []encoding.ServerLink
}

// ID ..
func (p *ServerLinks) ID() int32 {
	return IDClientBoundServerLinks
}

// Resource ..
func (p *ServerLinks) Resource() string {
	return "server_links"
}

// BoundType ..
func (p *ServerLinks) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ServerLinks) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.Links)
}
