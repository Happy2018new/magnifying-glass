package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
	"github.com/google/uuid"
)

// BossBar ..
type BossBar struct {
	// Unique ID for this bar.
	UUID uuid.UUID
	// Action ..
	Action encoding.BossBarAction
}

// ID ..
func (p *BossBar) ID() int32 {
	return IDClientBoundBossBar
}

// Resource ..
func (p *BossBar) Resource() string {
	return "boss_event"
}

// BoundType ..
func (p *BossBar) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *BossBar) Marshal(io encoding.IO) {
	io.UUID(&p.UUID)
	io.BossBarActionType(&p.Action)
	p.Action.Marshal(io)
}
