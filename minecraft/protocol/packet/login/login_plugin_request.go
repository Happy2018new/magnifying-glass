package packet_login

import (
	"magnifying-glass/minecraft/protocol/encoding"
	packet_interface "magnifying-glass/minecraft/protocol/packet/interface"
)

// Used to implement a custom handshaking flow together with
// Login Plugin Response (https://minecraft.wiki/w/Java_Edition_protocol#Login_Plugin_Response).
//
// Unlike plugin messages in "play" mode, these messages follow
// a lock-step request/response scheme, where the client is
// expected to respond to a request indicating whether it
// understood.
//
// The vanilla client always responds that it hasn't understood,
// and sends an empty payload.
type LoginPluginRequest struct {
	// Generated by the server - should be unique to the connection.
	MessageID int32
	// Name of the plugin channel used to send the data.
	// plugin channel (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Plugin_channels).
	Channel encoding.Identifier
	// Any data, depending on the channel.
	// The length of this array must be inferred from the
	// packet length.
	Data []byte
}

// ID ..
func (p *LoginPluginRequest) ID() int32 {
	return IDClientBoundLoginPluginRequest
}

// Resource ..
func (p *LoginPluginRequest) Resource() string {
	return "custom_query"
}

// BoundType ..
func (p *LoginPluginRequest) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *LoginPluginRequest) Marshal(io encoding.IO) {
	io.Varint32(&p.MessageID)
	io.Identifier(&p.Channel)
	io.Bytes(&p.Data)
}
