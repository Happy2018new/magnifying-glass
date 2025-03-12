package encoding

import "github.com/Happy2018new/magnifying-glass/minecraft/nbt"

// RegistryEntry is used in Registry Data packet.
// For more information, see the following links for details.
//		- Registry Data (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Registry_Data)
type RegistryEntry struct {
	// EntryID ..
	EntryID Identifier
	// Data ..
	Data Optional[map[string]any]
}

func (r *RegistryEntry) Marshal(io IO) {
	io.Identifier(&r.EntryID)
	OptionalFunc(io, &r.Data, func(t *map[string]any) { io.NBT(t, nbt.NetworkBigEndian) })
}

// RegistryTag is used in Update Tags packet.
//
// See Tag on the Minecraft Wiki for more information,
// including a list of vanilla tags.
//		- Tag (https://minecraft.wiki/w/Tag)
type RegistryTag struct {
	// Registry identifier
	// (
	// 		Vanilla expects tags for the registries
	// 			minecraft:block,
	// 			minecraft:item,
	// 			minecraft:fluid,
	// 			minecraft:entity_type,
	// 			and minecraft:game_event
	//	)
	Registry Identifier
	// See RegistryArrayTag for more information.
	ArrayOfTag []RegistryArrayTag
}

func (r *RegistryTag) Marshal(io IO) {
	io.Identifier(&r.Registry)
	SliceVarint32Length(io, &r.ArrayOfTag)
}

// RegistryTag ..
type RegistryArrayTag struct {
	// TagName ..
	TagName Identifier
	// Numeric IDs of the given type (block, item, etc.).
	//
	// This list replaces the previous list of IDs for
	// the given tag.
	// If some preexisting tags are left unmentioned,
	// a warning is printed.
	Entries []int32
}

func (r *RegistryArrayTag) Marshal(io IO) {
	io.Identifier(&r.TagName)
	FuncSliceVarint32Length(io, &r.Entries, io.Varint32)
}
