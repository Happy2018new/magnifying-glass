package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// The server responds with a list of
// auto-completions of the last word sent to it.
// In the case of regular chat, this is a player username.
// Command names and parameters are also supported.
// The client sorts these alphabetically before listing them.
type CommandSuggestionsResponse struct {
	// Transaction ID.
	TransactionID int32
	// Start of the text to replace.
	Start int32
	// Length of the text to replace.
	Length int32
	// See CommandSuggestMatch for more information.
	Matches []encoding.CommandSuggestMatch
}

// ID ..
func (p *CommandSuggestionsResponse) ID() int32 {
	return IDClientBoundCommandSuggestionsResponse
}

// Resource ..
func (p *CommandSuggestionsResponse) Resource() string {
	return "command_suggestions"
}

// BoundType ..
func (p *CommandSuggestionsResponse) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *CommandSuggestionsResponse) Marshal(io encoding.IO) {
	io.Varint32(&p.TransactionID)
	io.Varint32(&p.Start)
	io.Varint32(&p.Length)
	encoding.SliceVarint32Length(io, &p.Matches)
}
