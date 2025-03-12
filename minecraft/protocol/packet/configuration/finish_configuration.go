package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sent by the server to notify the client that
// the configuration process has finished.
//
// The client answers with
// Acknowledge Finish Configuration (https://minecraft.wiki/w/Java_Edition_protocol#Acknowledge_Finish_Configuration)
// whenever it is ready to continue.
type FinishConfiguration struct{}

// ID ..
func (p *FinishConfiguration) ID() int32 {
	return IDClientBoundFinishConfiguration
}

// Resource ..
func (p *FinishConfiguration) Resource() string {
	return "finish_configuration"
}

// BoundType ..
func (p *FinishConfiguration) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *FinishConfiguration) Marshal(io encoding.IO) {}
