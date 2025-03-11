package packet_interface

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
)

const (
	BoundTypeServer uint8 = iota
	BoundTypeClient
)

// Packet represents a packet that may be sent over a Minecraft network connection.
// The packet hold some methods to allows you to do the following things.
//   - Get the ID of this packet.
//   - Get the resource string of this packet.
//   - Understand this packet is used to send to server or send to client.
//   - Encode this packet to binary and decode this packet from binary.
type Packet interface {
	// ID returns the ID of the packet.
	// All of these identifiers of packets may be found in id.go.
	// Note that there a multiple id.go.
	ID() int32
	// Resource returns the resource of the packet.
	// However, I don't know how it works currently.
	Resource() string
	// BoundType returns the bound type of the packet.
	// If return 0 (BoundTypeServer), it means this packet is send from client to server.
	// Otherwise (1; BoundTypeClient) this packet is send from server to client.
	BoundType() uint8
	// Marshal encodes or decodes a Packet, depending on the encoding.IO
	// implementation passed. When passing a protocol.Writer, Marshal will
	// encode the Packet into its binary representation and write it to the
	// protocol.Writer. On the other hand, when passing a protocol.Reader,
	// Marshal will decode the bytes from the reader into the Packet.
	Marshal(io encoding.IO)
}

// Pool is a map holding packets indexed by a packet ID.
type Pool map[int32]Packet
