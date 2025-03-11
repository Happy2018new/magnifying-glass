package packet_login

import packet_interface "magnifying-glass/minecraft/protocol/packet/interface"

// NewClientPool returns a new pool containing packets sent by a client.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewClientPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{
		IDServerBoundLoginStart:          &LoginStart{},
		IDServerBoundEncryptionResponse:  &EncryptionResponse{},
		IDServerBoundLoginPluginResponse: &LoginPluginResponse{},
		IDServerBoundLoginAcknowledged:   &LoginAcknowledged{},
		IDServerBoundCookieResponse:      &CookieResponse{},
	}
}

// NewServerPool returns a new pool containing packets sent by a server.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewServerPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{
		IDClientBoundDisconnect:         &Disconnect{},
		IDClientBoundEncryptionRequest:  &EncryptionRequest{},
		IDClientBoundLoginSuccess:       &LoginSuccess{},
		IDClientBoundSetCompression:     &SetCompression{},
		IDClientBoundLoginPluginRequest: &LoginPluginRequest{},
		IDClientBoundCookieRequest:      &CookieRequest{},
	}
}
