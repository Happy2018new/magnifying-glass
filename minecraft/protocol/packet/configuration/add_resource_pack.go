package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
	"github.com/google/uuid"
)

// AddResourcePack ..
type AddResourcePack struct {
	// The unique identifier of the resource pack.
	UUID uuid.UUID
	// uuid.UUID
	URL string
	// A 40 character hexadecimal,
	// case-insensitive SHA-1 hash
	// of the resource pack file.
	//
	// If it's not a 40 character
	// hexadecimal string,
	// the client will not use it
	// for hash verification and
	// likely waste bandwidth.
	Hash string
	// The vanilla client will be forced
	// to use the resource pack from the server.
	// If they decline they will be kicked from the server.
	Forced bool
	// This is shown in the prompt making
	// the client accept or decline the
	// resource pack (only if present).
	PromptMessage encoding.TextComponentOptional
}

// ID ..
func (p *AddResourcePack) ID() int32 {
	return IDClientBoundAddResourcePack
}

// Resource ..
func (p *AddResourcePack) Resource() string {
	return "resource_pack_push"
}

// BoundType ..
func (p *AddResourcePack) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *AddResourcePack) Marshal(io encoding.IO) {
	io.UUID(&p.UUID)
	io.String(&p.URL)
	io.String(&p.Hash)
	io.Bool(&p.Forced)
	encoding.Single(io, &p.PromptMessage)
}
