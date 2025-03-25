package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Entity statuses generally trigger
// an animation for an entity.
// The available statuses vary by the
// entity's type (and are available to
// subclasses of that type as well).
type EntityEvent struct {
	// EntityID ..
	EntityID int32
	// See Entity statuses for a list of
	// which statuses are valid for each
	// type of entity.
	//		- Entity statuses (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Entity_statuses)
	EntityStatus byte
}

// ID ..
func (p *EntityEvent) ID() int32 {
	return IDClientBoundEntityEvent
}

// Resource ..
func (p *EntityEvent) Resource() string {
	return "entity_event"
}

// BoundType ..
func (p *EntityEvent) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *EntityEvent) Marshal(io encoding.IO) {
	io.Int32(&p.EntityID)
	io.Uint8(&p.EntityStatus)
}
