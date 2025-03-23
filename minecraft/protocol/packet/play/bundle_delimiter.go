package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// The delimiter for a bundle of packets.
// When received, the client should store
// every subsequent packet it receives,
// and wait until another delimiter is received.
// Once that happens, the client is guaranteed
// to process every packet in the bundle on the
// same tick, and the client should stop storing
// packets.
//
// As of 1.20.6, the vanilla server only uses this
// to ensure Spawn Entity (https://minecraft.wiki/w/Java_Edition_protocol#Spawn_Entity)
// and associated packets used to configure the entity
// happen on the same tick.
// Each entity gets a separate bundle.
//
// The vanilla client doesn't allow more than 4096
// packets in the same bundle.
type BundleDelimiter struct{}

// ID ..
func (p *BundleDelimiter) ID() int32 {
	return IDClientBoundBundleDelimiter
}

// Resource ..
func (p *BundleDelimiter) Resource() string {
	return "bundle_delimiter"
}

// BoundType ..
func (p *BundleDelimiter) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *BundleDelimiter) Marshal(io encoding.IO) {}
