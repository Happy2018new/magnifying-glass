package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Acknowledges a user-initiated block change.
// After receiving this packet, the client will
// display the block state sent by the server
// instead of the one predicted by the client.
type AcknowledgeBlockChange struct {
	// Represents the sequence to acknowledge,
	// this is used for properly syncing block
	// changes to the client after interactions.
	SequenceID int32
}

// ID ..
func (p *AcknowledgeBlockChange) ID() int32 {
	return IDClientBoundAcknowledgeBlockChange
}

// Resource ..
func (p *AcknowledgeBlockChange) Resource() string {
	return "block_changed_ack"
}

// BoundType ..
func (p *AcknowledgeBlockChange) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *AcknowledgeBlockChange) Marshal(io encoding.IO) {
	io.Varint32(&p.SequenceID)
}
