package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

const SampleTypeTickTime int32 = iota

// Note that the vanilla client calculates the
// timings used for min/max/average display by
// subtracting the idle time from the full tick
// time.
// This can cause the displayed values to go negative
// if the idle time is (nonsensically) greater than
// the full tick time.
const (
	// Aggregate of the three times below.
	SamplesIndexTickTimeFull = iota
	// Server tick time: Main server tick logic.
	SamplesIndexTickTimeServer
	// Tasks time: Tasks scheduled to execute after the main logic.
	SamplesIndexTickTimeTasks
	// Idle time: Time idling to complete the full 50ms tick cycle.
	SamplesIndexTickTimeIdle
)

// Sample data that is sent  periodically
// after the client has subscribed with
// Debug Sample Subscription (https://minecraft.wiki/w/Java_Edition_protocol#Debug_Sample_Subscription).
//
// The vanilla server only sends debug samples
// to players that are server operators.
type DebugSample struct {
	// Array of type-dependent samples.
	Sample []int64
	// See constant enum above.
	SampleType int32
}

// ID ..
func (p *DebugSample) ID() int32 {
	return IDClientBoundDebugSample
}

// Resource ..
func (p *DebugSample) Resource() string {
	return "debug_sample"
}

// BoundType ..
func (p *DebugSample) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *DebugSample) Marshal(io encoding.IO) {
	encoding.FuncSliceVarint32Length(io, &p.Sample, io.Int64)
	io.Varint32(&p.SampleType)
}
