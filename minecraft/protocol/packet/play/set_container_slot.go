package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sent by the server when an item in a slot
// (in a window) is added/removed.
//
// If Window ID is 0, the hotbar and offhand slots
// (slots 36 through 45) may be updated even when a
// different container window is open (The vanilla server
// does not appear to utilize this special case).
// Updates are also restricted to those slots when the
// player is looking at a creative inventory tab other
// than the survival inventory (The vanilla server does
// not handle this restriction in any way, leading to
// MC-242392).
//
// If Window ID is -1, the item being dragged with the
// mouse is set.
// In this case, State ID and Slot are ignored.
//
// If Window ID is -2, any slot in the player's inventory can
// be updated irrespective of the current container window.
// In this case, State ID is ignored, and the vanilla server
// uses a bogus value of 0. Used by the vanilla server to
// implement the #Pick Item functionality.
//
// When a container window is open, the server never sends updates targeting
// Window ID 0—all of the window types include slots for the player inventory.
// The client must automatically apply changes targeting the inventory
// portion of a container window to the main inventory; the server does
// not resend them for ID 0 when the window is closed.
// However, since the armor and offhand slots are only present on ID 0,
// updates to those slots occurring while a window is open must be deferred
// by the server until the window's closure.
//
// Helpful links:
//   - #Pick Item (https://minecraft.wiki/w/Java_Edition_protocol#Pick_Item)
type SetContainerSlot struct {
	// The window which is being updated.
	// 0 for player inventory.
	// The client ignores any packets targeting
	// a Window ID other than the current one;
	// see the introduction of SetContainerSlot for exceptions.
	WindowID int32
	// A server-managed sequence number
	// used to avoid desynchronization;
	// see #Click Container (https://minecraft.wiki/w/Java_Edition_protocol#Click_Container).
	StateID int32
	// The slot that should be updated.
	Slot int16
	// SlotData ..
	SlotData encoding.ItemStack
}

// ID ..
func (p *SetContainerSlot) ID() int32 {
	return IDClientBoundSetContainerSlot
}

// Resource ..
func (p *SetContainerSlot) Resource() string {
	return "container_set_slot"
}

// BoundType ..
func (p *SetContainerSlot) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SetContainerSlot) Marshal(io encoding.IO) {
	io.Varint32(&p.WindowID)
	io.Varint32(&p.StateID)
	io.Int16(&p.Slot)
	io.ItemStack(&p.SlotData)
}
