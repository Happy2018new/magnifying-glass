package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Replaces the contents of a container window.
// Sent by the server upon initialization of a
// container window or the player's inventory,
// and in response to state ID mismatches
// (See #Click Container).
//
// See inventory windows for further information
// about how slots are indexed.
// Use Open Screen to open the container on the client.
//
// Some helpful links:
//   - inventory windows (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Inventory#Windows)
//   - Open Screen (https://minecraft.wiki/w/Java_Edition_protocol#Open_Screen)
//   - #Click Container (https://minecraft.wiki/w/Java_Edition_protocol#Click_Container)
type SetContainerContent struct {
	// The ID of window which items are being sent for.
	// 0 for player inventory.
	// The client ignores any packets targeting a Window ID
	// other than the current one.
	// However, an exception is made for the player inventory,
	// which may be targeted at any time.
	// (The vanilla server does not appear to utilize this
	// special case.)
	WindowID int32
	// A server-managed sequence number
	// used to avoid desynchronization;
	// see #Click Container (https://minecraft.wiki/w/Java_Edition_protocol#Click_Container).
	StateID int32
	// SlotDate ..
	SlotDate []encoding.ItemStack
	// Item being dragged with the mouse.
	CarriedItem encoding.ItemStack
}

// ID ..
func (p *SetContainerContent) ID() int32 {
	return IDClientBoundSetContainerContent
}

// Resource ..
func (p *SetContainerContent) Resource() string {
	return "container_set_content"
}

// BoundType ..
func (p *SetContainerContent) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SetContainerContent) Marshal(io encoding.IO) {
	io.Varint32(&p.WindowID)
	io.Varint32(&p.StateID)
	encoding.FuncSliceVarint32Length(io, &p.SlotDate, io.ItemStack)
	io.ItemStack(&p.CarriedItem)
}
