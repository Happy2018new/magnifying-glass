package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Used to enable and disable features,
// generally experimental ones, on the client.
//
// There is one special feature flag, which is in most versions:
//   - minecraft:vanilla - enables vanilla features
//
// For the other feature flags, which may change between versions,
// see Experiments#Java_Edition (https://minecraft.wiki/w/Experiments#Java_Edition).
type FeatureFlags struct {
	// FeatureFlags ..
	FeatureFlags []encoding.Identifier
}

// ID ..
func (p *FeatureFlags) ID() int32 {
	return IDClientBoundFeatureFlags
}

// Resource ..
func (p *FeatureFlags) Resource() string {
	return "update_enabled_features"
}

// BoundType ..
func (p *FeatureFlags) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *FeatureFlags) Marshal(io encoding.IO) {
	encoding.FuncSliceVarint32Length(io, &p.FeatureFlags, io.Identifier)
}
