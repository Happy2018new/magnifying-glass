package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

const (
	ChatModeEnabed int32 = iota
	ChatModeCommandsOnly
	ChatModeHidden
)

const (
	DisplayedSkinPartFlagCapEnabled uint8 = 1 << iota
	DisplayedSkinPartFlagJacketEnabled
	DisplayedSkinPartFlagLeftSleeveEnabled
	DisplayedSkinPartFlagRightSleeveEnabled
	DisplayedSkinPartFlagLeftPantsLegEnabled
	DisplayedSkinPartFlagRightPantsLegEnabled
	DisplayedSkinPartFlagHatEnabled
)

const (
	MainHandLeft int32 = iota
	MainHandRight
)

const (
	ParticleStatusAll int32 = iota
	ParticleStatusDecreased
	ParticleStatusMinimal
)

// Sent when the player connects,
// or when settings are changed.
type ClientInformation struct {
	// e.g. en_GB.
	Locale string
	// Client-side render distance, in chunks.
	ViewDistance byte
	// 0: enabled,
	// 1: commands only,
	// 2: hidden.
	//
	// See
	// Chat#Client chat mode (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Chat#Client_chat_mode)
	// for more information.
	ChatMode int32
	// “Colors” multiplayer setting.
	// The vanilla server stores this value but does
	// nothing with it (see MC-64867).
	//
	// Third-party servers such as Hypixel disable all
	// coloring in chat and system messages when it is
	// false.
	//
	// See the following links for more information.
	// 		- MC-64867 (https://bugs.mojang.com/browse/MC-64867)
	ChatColors bool
	// Bit mask, see the enumerate above.
	DisplayedSkinParts byte
	// 0: Left,
	// 1: Right.
	MainHand int32
	// Enables filtering of text on signs and
	// written book titles.
	//
	// The vanilla client sets this according to the
	// "profanityFilterPreferences.profanityFilterOn"
	// account attribute indicated by the "/player/attributes"
	// Mojang API endpoint (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Mojang_API#Player_Attributes).
	//
	// In offline mode it is always false.
	EnableTextFiltering bool
	// Servers usually list online players,
	// this option should let you not show up
	// in that list.
	AllowServerListings bool
	// 0: all,
	// 1: decreased,
	// 2: minimal
	ParticleStatus int32
}

// ID ..
func (p *ClientInformation) ID() int32 {
	return IDServerBoundClientInformation
}

// Resource ..
func (p *ClientInformation) Resource() string {
	return "client_information"
}

// BoundType ..
func (p *ClientInformation) BoundType() uint8 {
	return packet_interface.BoundTypeServer
}

func (p *ClientInformation) Marshal(io encoding.IO) {
	io.String(&p.Locale)
	io.Uint8(&p.ViewDistance)
	io.Varint32(&p.ChatMode)
	io.Bool(&p.ChatColors)
	io.Uint8(&p.DisplayedSkinParts)
	io.Varint32(&p.MainHand)
	io.Bool(&p.EnableTextFiltering)
	io.Bool(&p.AllowServerListings)
	io.Varint32(&p.ParticleStatus)
}
