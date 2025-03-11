package packet_status

import (
	"magnifying-glass/minecraft/protocol/encoding"
	packet_interface "magnifying-glass/minecraft/protocol/packet/interface"
)

// StatusResponse ..
type StatusResponse struct {
	// See Server List Ping#Status Response (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Server_List_Ping#Status_Response);
	// as with all strings this is prefixed by its length as a VarInt.
	JsonResponse string
}

// ID ..
func (p *StatusResponse) ID() int32 {
	return IDClientBoundStatusResponse
}

// Resource ..
func (p *StatusResponse) Resource() string {
	return "status_response"
}

// BoundType ..
func (p *StatusResponse) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *StatusResponse) Marshal(io encoding.IO) {
	io.String(&p.JsonResponse)
}
