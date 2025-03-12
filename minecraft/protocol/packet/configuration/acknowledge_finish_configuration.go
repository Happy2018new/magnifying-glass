package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sent by the client to notify the server
// that the configuration process has finished.
// It is sent in response to the server's
// Finish Configuration (https://minecraft.wiki/w/Java_Edition_protocol#Finish_Configuration).
type AcknowledgeFinishConfiguration struct{}

// ID ..
func (p *AcknowledgeFinishConfiguration) ID() int32 {
	return IDServerBoundAcknowledgeFinishConfiguration
}

// Resource ..
func (p *AcknowledgeFinishConfiguration) Resource() string {
	return "finish_configuration"
}

// BoundType ..
func (p *AcknowledgeFinishConfiguration) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *AcknowledgeFinishConfiguration) Marshal(io encoding.IO) {}
