package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

const ContainerWindowIDInventory int32 = iota

// This packet is sent from the server to the
// client when a window is forcibly closed,
// such as when a chest is destroyed while
// it's open.
// The vanilla client disregards the provided
// window ID and closes any active window.
type CloseContainer struct {
	// This is the ID of the window that was closed.
	// 0 for inventory.
	WindowID int32
}

// ID ..
func (p *CloseContainer) ID() int32 {
	return IDClientBoundCloseContainer
}

// Resource ..
func (p *CloseContainer) Resource() string {
	return "container_close"
}

// BoundType ..
func (p *CloseContainer) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *CloseContainer) Marshal(io encoding.IO) {
	io.Varint32(&p.WindowID)
}
