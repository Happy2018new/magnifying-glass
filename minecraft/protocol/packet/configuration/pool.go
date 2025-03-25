package packet_configuration

import packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"

// NewClientPool returns a new pool containing packets sent by a client.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewClientPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{
		IDServerBoundClientInformation:              &ClientInformation{},
		IDServerBoundCookieResponse:                 &CookieResponse{},
		IDServerBoundPluginMessage:                  &ServerboundPluginMessage{},
		IDServerBoundAcknowledgeFinishConfiguration: &AcknowledgeFinishConfiguration{},
		IDServerBoundKeepAlive:                      &ServerboundKeepAlive{},
		IDServerBoundPong:                           &Pong{},
		IDServerBoundResourcePackResponse:           &ResourcePackResponse{},
		IDServerBoundKnownPacks:                     &ServerboundKnownPacks{},
	}
}

// NewServerPool returns a new pool containing packets sent by a server.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewServerPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{
		IDClientBoundCookieRequest:       &CookieRequest{},
		IDClientBoundPluginMessage:       &ClientBoundPluginMessage{},
		IDClientBoundDisconnect:          &Disconnect{},
		IDClientBoundFinishConfiguration: &FinishConfiguration{},
		IDClientBoundKeepAlive:           &ClientBoundKeepAlive{},
		IDClientBoundPing:                &Ping{},
		IDClientBoundResetChat:           &ResetChat{},
		IDClientBoundRegistryData:        &RegistryData{},
		IDClientBoundRemoveResourcePack:  &RemoveResourcePack{},
		IDClientBoundAddResourcePack:     &AddResourcePack{},
		IDClientBoundStoreCookie:         &StoreCookie{},
		IDClientBoundTransfer:            &Transfer{},
		IDClientBoundFeatureFlags:        &FeatureFlags{},
		IDClientBoundUpdateTags:          &UpdateTags{},
		IDClientBoundKnownPacks:          &ClientBoundKnownPacks{},
		IDClientBoundCustomReportDetails: &CustomReportDetails{},
		IDClientBoundServerLinks:         &ServerLinks{},
	}
}
