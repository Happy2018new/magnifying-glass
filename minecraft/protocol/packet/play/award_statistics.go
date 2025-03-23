package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Sent as a response to Client Status (id 1).
// Will only send the changed values if previously requested.
//
// See the following links for information.
//   - Client Status (https://minecraft.wiki/w/Java_Edition_protocol#Client_Status)
type AwardStatistics struct {
	// See AwardStatistic for more information.
	Statistic []encoding.AwardStatistic
}

// ID ..
func (p *AwardStatistics) ID() int32 {
	return IDClientBoundAwardStatistics
}

// Resource ..
func (p *AwardStatistics) Resource() string {
	return "award_stats"
}

// BoundType ..
func (p *AwardStatistics) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *AwardStatistics) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.Statistic)
}
