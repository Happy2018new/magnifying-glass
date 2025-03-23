package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sends the client a chat message,
// but without any message signing
// information.
// The vanilla server uses this packet
// when the console is communicating with
// players through commands, such as /say,
// /tell, /me, among others.
type DisguisedChatMessage struct {
	// This is used as the content parameter
	// when formatting the message on the client.
	Message encoding.TextComponent
	// The type of chat in the minecraft:chat_type registry,
	// defined by the Registry Data packet (https://minecraft.wiki/w/Java_Edition_protocol#Registry_Data).
	ChatType int32
	// The name of the one sending the message,
	// usually the sender's display name.
	SenderName encoding.TextComponent
	// The name of the one receiving the message,
	// usually the receiver's display name.
	// This is used as the target parameter when
	// formatting the message on the client.
	TargetName encoding.TextComponentOptional
}

// ID ..
func (p *DisguisedChatMessage) ID() int32 {
	return IDClientBoundDisguisedChatMessage
}

// Resource ..
func (p *DisguisedChatMessage) Resource() string {
	return "disguised_chat"
}

// BoundType ..
func (p *DisguisedChatMessage) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *DisguisedChatMessage) Marshal(io encoding.IO) {
	io.TextComponent(&p.Message)
	io.Varint32(&p.ChatType)
	io.TextComponent(&p.SenderName)
	encoding.Single(io, &p.TargetName)
}
