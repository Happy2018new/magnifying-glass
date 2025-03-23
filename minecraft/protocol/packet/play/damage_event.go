package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// DamageEvent ..
type DamageEvent struct {
	// The ID of the entity taking damage.
	EntityID int32
	// The type of damage in the minecraft:damage_type registry,
	// defined by the Registry Data (https://minecraft.wiki/w/Java_Edition_protocol#Registry_Data) packet.
	SourceTypeID int32
	// The ID + 1 of the entity responsible for the damage,
	// if present. If not present, the value is 0.
	SourceCauseID int32
	// The ID + 1 of the entity that directly dealt the damage,
	// if present. If not present, the value is 0.
	//
	// If this field is present:
	//		- and damage was dealt indirectly,
	// 		  such as by the use of a projectile,
	// 		  this field will contain the ID of such projectile;
	//		- and damage was dealt dirctly,
	// 		  such as by manually attacking,
	// 		  this field will contain the same
	// 		  value as Source Cause ID.
	SourceDirectID int32
	// The vanilla server sends the Source Position when the
	// damage was dealt by the "/damage" command and a position
	// was specified.
	SourcePosition encoding.Optional[encoding.EntityPos]
}

// ID ..
func (p *DamageEvent) ID() int32 {
	return IDClientBoundDamageEvent
}

// Resource ..
func (p *DamageEvent) Resource() string {
	return "damage_event"
}

// BoundType ..
func (p *DamageEvent) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *DamageEvent) Marshal(io encoding.IO) {
	io.Varint32(&p.EntityID)
	io.Varint32(&p.SourceTypeID)
	io.Varint32(&p.SourceCauseID)
	io.Varint32(&p.SourceDirectID)
	encoding.OptionalFunc(io, &p.SourcePosition, io.EntityPos)
}
