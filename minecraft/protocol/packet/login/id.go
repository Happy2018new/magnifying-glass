package packet_login

const (
	IDClientBoundDisconnect int32 = iota
	IDClientBoundEncryptionRequest
	IDClientBoundLoginSuccess
	IDClientBoundSetCompression
	IDClientBoundLoginPluginRequest
	IDClientBoundCookieRequest
)

const (
	IDServerBoundLoginStart int32 = iota
	IDServerBoundEncryptionResponse
	IDServerBoundLoginPluginResponse
	IDServerBoundLoginAcknowledged
	IDServerBoundCookieResponse
)
