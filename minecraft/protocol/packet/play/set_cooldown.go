package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Applies a cooldown period to all items with the given type.
// Used by the vanilla server with enderpearls.
//
// This packet should be sent when the cooldown starts and also
// when the cooldown ends (to compensate for lag),
// although the client will end the cooldown automatically.
//
// Can be applied to any item, note that interactions still get
// sent to the server with the item but the client does not play
// the animation nor attempt to predict results (i.e block placing).
type SetCooldown struct {
	// Numeric ID of the item to apply a cooldown to.
	ItemID int32
	// Number of ticks to apply a cooldown for,
	// or 0 to clear the cooldown.
	CooldownTicks int32
}

// ID ..
func (p *SetCooldown) ID() int32 {
	return IDClientBoundSetCooldown
}

// Resource ..
func (p *SetCooldown) Resource() string {
	return "cooldown"
}

// BoundType ..
func (p *SetCooldown) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SetCooldown) Marshal(io encoding.IO) {
	io.Varint32(&p.ItemID)
	io.Varint32(&p.CooldownTicks)
}
