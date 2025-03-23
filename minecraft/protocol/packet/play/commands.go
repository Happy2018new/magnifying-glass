package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Lists all of the commands on the server,
// and how they are parsed.
// This is a directed graph, with one root node.
// Each redirect or child node must refer only to
// nodes that have already been declared.
//
// For more information on this packet,
// see the Command Data article (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Command_Data).
type Commands struct {
	// An array of nodes.
	Nodes []encoding.CommandNode
	// Index of the root node in the previous array.
	RootIndex int32
}

// ID ..
func (p *Commands) ID() int32 {
	return IDClientBoundCommands
}

// Resource ..
func (p *Commands) Resource() string {
	return "commands"
}

// BoundType ..
func (p *Commands) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *Commands) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.Nodes)
	io.Varint32(&p.RootIndex)
}
