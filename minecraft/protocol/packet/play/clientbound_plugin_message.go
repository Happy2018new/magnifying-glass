package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Mods and plugins can use this to send their data.
// Minecraft itself uses several plugin channels.
// These internal channels are in the minecraft namespace.
//
// More information on how it works on Dinnerbone's blog.
// More documentation about internal and popular registered
// channels are here (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Plugin_channels).
//
// In vanilla clients, the maximum data length is 1048576 bytes.
//
// Helpful links:
//   - plugin channels (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Plugin_channels)
//   - Dinnerbone's blog (https://dinnerbone.com/blog/2012/01/13/minecraft-plugin-channels-messaging/)
type ClientBoundPluginMessage struct {
	// Name of the
	// plugin channel (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Plugin_channels)
	// used to send the data.
	Channel encoding.Identifier
	// Any data. The length of this array must
	// be inferred from the packet length.
	Data []byte
}

// ID ..
func (p *ClientBoundPluginMessage) ID() int32 {
	return IDClientBoundPluginMessage
}

// Resource ..
func (p *ClientBoundPluginMessage) Resource() string {
	return "custom_payload"
}

// BoundType ..
func (p *ClientBoundPluginMessage) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ClientBoundPluginMessage) Marshal(io encoding.IO) {
	io.Identifier(&p.Channel)
	io.Bytes(&p.Data)
}
