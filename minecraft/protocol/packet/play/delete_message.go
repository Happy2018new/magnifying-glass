package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

const DeleteMessageSignatureLength uint32 = 256

// Removes a message from the client's chat.
// This only works for messages with signatures,
// system messages cannot be deleted with this packet.
type DeleteMessage struct {
	// The message Id + 1, used for validating message signature.
	// The next field is present only when value of this field is
	// equal to 0.
	MessageID int32
	// The previous message's signature.
	// Always 256 bytes and not length-prefixed.
	Signature []byte
}

// ID ..
func (p *DeleteMessage) ID() int32 {
	return IDClientBoundDeleteMessage
}

// Resource ..
func (p *DeleteMessage) Resource() string {
	return "delete_chat"
}

// BoundType ..
func (p *DeleteMessage) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *DeleteMessage) Marshal(io encoding.IO) {
	io.Varint32(&p.MessageID)
	if p.MessageID == 0 {
		encoding.FuncSliceOfLen(io, DeleteMessageSignatureLength, &p.Signature, io.Uint8)
	}
}
