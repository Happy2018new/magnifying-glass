package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Mods and plugins can use this to send their data.
// Minecraft itself uses some plugin channels.
// These internal channels are in the minecraft namespace.
//
// More documentation on this:
// (https://dinnerbone.com/blog/2012/01/13/minecraft-plugin-channels-messaging).
//
// Note that the length of Data is known only from the
// packet length,
// since the packet has no length field of any kind.
//
// Helpful links is as follows.
//   - Plugin channels (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Plugin_channels)
type ServerboundPluginMessage struct {
	// Name of the
	// plugin channel (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Plugin_channels)
	// used to send the data.
	Channel encoding.Identifier
	// Any data, depending on the channel.
	// minecraft: channels are documented
	// here (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Plugin_channels).
	//
	// The length of this array must be inferred
	// from the packet length.
	Data []byte
}

// ID ..
func (p *ServerboundPluginMessage) ID() int32 {
	return IDServerBoundPluginMessage
}

// Resource ..
func (p *ServerboundPluginMessage) Resource() string {
	return "custom_payload"
}

// BoundType ..
func (p *ServerboundPluginMessage) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *ServerboundPluginMessage) Marshal(io encoding.IO) {
	io.Identifier(&p.Channel)
	io.Bytes(&p.Data)
}
