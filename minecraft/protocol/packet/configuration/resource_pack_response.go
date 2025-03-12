package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
	"github.com/google/uuid"
)

const (
	ResourcePackResponseResultSuccessfullyDownloaded int32 = iota
	ResourcePackResponseResultDeclined
	ResourcePackResponseResultFailedToDownload
	ResourcePackResponseResultAccepted
	ResourcePackResponseResultDownloaded
	ResourcePackResponseResultInvalidURL
	ResourcePackResponseResultFailedToReload
	ResourcePackResponseResultDiscarded
)

// ResourcePackResponse ..
type ResourcePackResponse struct {
	// The unique identifier of the
	// resource pack received in the
	// Add Resource Pack (configuration) (https://minecraft.wiki/w/Java_Edition_protocol#Add_Resource_Pack_(configuration))
	// request.
	UUID uuid.UUID
	// Result ID, see const enum above.
	Result int32
}

// ID ..
func (p *ResourcePackResponse) ID() int32 {
	return IDServerBoundResourcePackResponse
}

// Resource ..
func (p *ResourcePackResponse) Resource() string {
	return "resource_pack"
}

// BoundType ..
func (p *ResourcePackResponse) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *ResourcePackResponse) Marshal(io encoding.IO) {
	io.UUID(&p.UUID)
	io.Varint32(&p.Result)
}
