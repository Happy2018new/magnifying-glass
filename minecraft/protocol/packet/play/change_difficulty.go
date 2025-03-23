package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

const (
	DifficultyPeaceful uint8 = iota
	DifficultyEasy
	DifficultyNormal
	DifficultyHard
)

// Changes the difficulty setting in the client's option menu.
type ChangeDifficulty struct {
	// 0: peaceful,
	// 1: easy,
	// 2: normal,
	// 3: hard.
	Difficulty byte
	// Difficulty locked ?
	DifficultyLocked bool
}

// ID ..
func (p *ChangeDifficulty) ID() int32 {
	return IDClientBoundChangeDifficulty
}

// Resource ..
func (p *ChangeDifficulty) Resource() string {
	return "change_difficulty"
}

// BoundType ..
func (p *ChangeDifficulty) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *ChangeDifficulty) Marshal(io encoding.IO) {
	io.Uint8(&p.Difficulty)
	io.Bool(&p.DifficultyLocked)
}
