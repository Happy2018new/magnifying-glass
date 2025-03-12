package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// See Tag on the Minecraft Wiki for more information,
// including a list of vanilla tags.
//   - Tag (https://minecraft.wiki/w/Tag)
type UpdateTags struct {
	// See RegistryTag for more information
	ArrayOfTags []encoding.RegistryTag
}

// ID ..
func (p *UpdateTags) ID() int32 {
	return IDClientBoundUpdateTags
}

// Resource ..
func (p *UpdateTags) Resource() string {
	return "update_tags"
}

// BoundType ..
func (p *UpdateTags) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *UpdateTags) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.ArrayOfTags)
}
