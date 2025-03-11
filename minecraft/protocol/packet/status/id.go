package packet_status

const (
	IDClientBoundStatusResponse int32 = iota
	IDClientBoundPongResponse
)

const (
	IDServerBoundStatusRequest int32 = iota
	IDServerBoundPingRequest
)
