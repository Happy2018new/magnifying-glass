package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
	"github.com/google/uuid"
)

// RemoveResourcePack ..
type RemoveResourcePack struct {
	// The UUID of the resource pack to be removed.
	// If not present every resource pack will be removed.
	UUID encoding.Optional[uuid.UUID]
}

// ID ..
func (p *RemoveResourcePack) ID() int32 {
	return IDClientBoundRemoveResourcePack
}

// Resource ..
func (p *RemoveResourcePack) Resource() string {
	return "resource_pack_pop"
}

// BoundType ..
func (p *RemoveResourcePack) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *RemoveResourcePack) Marshal(io encoding.IO) {
	encoding.OptionalFunc(io, &p.UUID, io.UUID)
}
