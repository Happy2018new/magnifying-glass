package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Clear the client's current title information,
// with the option to also reset it.
type ClearTitles struct {
	// Reset ..
	Reset bool
}

// ID ..
func (p *ClearTitles) ID() int32 {
	return IDClientBoundClearTitles
}

// Resource ..
func (p *ClearTitles) Resource() string {
	return "clear_titles"
}

// BoundType ..
func (p *ClearTitles) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ClearTitles) Marshal(io encoding.IO) {
	io.Bool(&p.Reset)
}
