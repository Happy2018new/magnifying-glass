package packet_status

import packet_interface "magnifying-glass/minecraft/protocol/packet/interface"

// NewClientPool returns a new pool containing packets sent by a client.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewClientPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{
		IDServerBoundStatusRequest: &StatusRequest{},
		IDServerBoundPingRequest:   &PingRequest{},
	}
}

// NewServerPool returns a new pool containing packets sent by a server.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewServerPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{
		IDClientBoundStatusResponse: &StatusResponse{},
		IDClientBoundPongResponse:   &PongResponse{},
	}
}
