package encoding

// ReportDetail is used in Custom
// Report Details packet.
//
// ReportDetail contains a key-value text
// entry that are included in any crash or
// disconnection report generated during
// connection to the server.
type ReportDetail struct {
	// Title ..
	Title string
	// Description ..
	Description string
}

func (r *ReportDetail) Marshal(io IO) {
	io.String(&r.Title)
	io.String(&r.Description)
}
