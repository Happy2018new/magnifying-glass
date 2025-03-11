package packet_login

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// See protocol encryption (https://minecraft.wiki/w/Protocol_encryption) for details.
type EncryptionResponse struct {
	// Shared Secret value, encrypted with the server's public key.
	SharedSecret []byte
	// Verify Token value, encrypted with the same public key as the shared secret.
	VerifyToken []byte
}

// ID ..
func (p *EncryptionResponse) ID() int32 {
	return IDServerBoundEncryptionResponse
}

// Resource ..
func (p *EncryptionResponse) Resource() string {
	return "key"
}

// BoundType ..
func (p *EncryptionResponse) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *EncryptionResponse) Marshal(io encoding.IO) {
	encoding.FuncSliceVarint32Length(io, &p.SharedSecret, io.Uint8)
	encoding.FuncSliceVarint32Length(io, &p.VerifyToken, io.Uint8)
}
