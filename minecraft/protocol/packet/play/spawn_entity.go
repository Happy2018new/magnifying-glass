package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
	"github.com/google/uuid"
)

// Sent by the server when an entity
// (aside from Experience Orb) is created.
//
// The points listed below should be considered
// when this packet is used to spawn a player entity.
//   - When in online mode (https://minecraft.wiki/w/Server.properties#online-mode),
//     the UUIDs must be valid and have valid skin blobs.
//     In offline mode, the vanilla server uses UUID v3 and
//     chooses the player's UUID by using the String
//     "OfflinePlayer:<player name>", encoding it in UTF-8
//     (and case-sensitive), then processes it with
//     "UUID.nameUUIDFromBytes".
//   - For NPCs UUID v2 should be used. Note:
//     "<+Grum> i will never confirm this as a feature you know that :)"
//   - In an example UUID, "xxxxxxxx-xxxx-Yxxx-xxxx-xxxxxxxxxxxx",
//     the UUID version is specified by "Y".
//     So, for UUID v3, "Y" will always be 3, and for UUID v2, "Y" will always be 2.
//
// See the following links for more information.
//   - Experience Orb (https://minecraft.wiki/w/Java_Edition_protocol#Spawn_Experience_Orb)
type SpawnEntity struct {
	// A unique integer ID mostly used in
	// the protocol to identify the entity.
	EntityID int32
	// A unique identifier that is mostly used
	// in persistence and places where the
	// uniqueness matters more.
	EntityUUID uuid.UUID
	// ID in the minecraft:entity_type registry
	// (see "type" field in Entity metadata#Entities).
	//		- Entity metadata#Entities (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Entity_metadata#Entities)
	Type int32
	// X ..
	X float64
	// Y ..
	Y float64
	// Z ..
	Z float64
	// Pitch ..
	Pitch float32
	// Yaw ..
	Yaw float32
	// HeadYaw ..
	HeadYaw float32
	// Meaning dependent on the value of the Type field,
	// see Object Data (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Object_Data)
	// for details.
	Data int32
	// Same units as
	// Set Entity Velocity (https://minecraft.wiki/w/Java_Edition_protocol#Set_Entity_Velocity).
	VelocityX int16
	// Same units as
	// Set Entity Velocity (https://minecraft.wiki/w/Java_Edition_protocol#Set_Entity_Velocity).
	VelocityY int16
	// Same units as
	// Set Entity Velocity (https://minecraft.wiki/w/Java_Edition_protocol#Set_Entity_Velocity).
	VelocityZ int16
}

// ID ..
func (p *SpawnEntity) ID() int32 {
	return IDClientBoundSpawnEntity
}

// Resource ..
func (p *SpawnEntity) Resource() string {
	return "add_entity"
}

// BoundType ..
func (p *SpawnEntity) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SpawnEntity) Marshal(io encoding.IO) {
	io.Varint32(&p.EntityID)
	io.UUID(&p.EntityUUID)
	io.Varint32(&p.Type)
	io.Float64(&p.X)
	io.Float64(&p.Y)
	io.Float64(&p.Z)
	io.Angle(&p.Pitch)
	io.Angle(&p.Yaw)
	io.Angle(&p.HeadYaw)
	io.Varint32(&p.Data)
	io.Int16(&p.VelocityX)
	io.Int16(&p.VelocityY)
	io.Int16(&p.VelocityZ)
}
