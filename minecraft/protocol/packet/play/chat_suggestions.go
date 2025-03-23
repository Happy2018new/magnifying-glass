package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

const (
	ChatSuggestAdd int32 = iota
	ChatSuggestRemove
	ChatSuggestSet
)

// Unused by the vanilla server.
// Likely provided for custom servers to
// send chat message completions to clients.
type ChatSuggestions struct {
	// See constant enum above.
	Action int32
	// Entries ..
	Entries []string
}

// ID ..
func (p *ChatSuggestions) ID() int32 {
	return IDClientBoundChatSuggestions
}

// Resource ..
func (p *ChatSuggestions) Resource() string {
	return "custom_chat_completions"
}

// BoundType ..
func (p *ChatSuggestions) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ChatSuggestions) Marshal(io encoding.IO) {
	io.Varint32(&p.Action)
	encoding.FuncSliceVarint32Length(io, &p.Entries, io.String)
}
