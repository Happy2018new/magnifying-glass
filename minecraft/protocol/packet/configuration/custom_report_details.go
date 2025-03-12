package packet_configuration

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Contains a list of key-value text entries that
// are included in any crash or disconnection report
// generated during connection to the server.
type CustomReportDetails struct {
	// See ReportDetail for more information.
	Details []encoding.ReportDetail
}

// ID ..
func (p *CustomReportDetails) ID() int32 {
	return IDClientBoundCustomReportDetails
}

// Resource ..
func (p *CustomReportDetails) Resource() string {
	return "custom_report_details"
}

// BoundType ..
func (p *CustomReportDetails) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *CustomReportDetails) Marshal(io encoding.IO) {
	encoding.SliceVarint32Length(io, &p.Details)
}
