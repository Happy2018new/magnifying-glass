package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Represents certain registries that are sent
// from the server and are applied on the client.
//
// See Registry Data for details.
// Registry Data (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Registry_Data).
type RegistryData struct {
	// RegistryID ..
	RegistryID encoding.Identifier
	// Entries ..
	Entries []encoding.RegistryEntry
}

// ID ..
func (p *RegistryData) ID() int32 {
	return IDClientBoundRegistryData
}

// Resource ..
func (p *RegistryData) Resource() string {
	return "registry_data"
}

// BoundType ..
func (p *RegistryData) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *RegistryData) Marshal(io encoding.IO) {
	io.Identifier(&p.RegistryID)
	encoding.SliceVarint32Length(io, &p.Entries)
}
