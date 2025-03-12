package encoding

// ResourcePack describe a Minecraft
// Java resource pack,
// which contains its name, uuid
// and their version.
type ResourcePack struct {
	// Namespace ..
	Namespace string
	// ID ..
	ID string
	// Version ..
	Version string
}

func (r *ResourcePack) Marshal(io IO) {
	io.String(&r.Namespace)
	io.String(&r.ID)
	io.String(&r.Version)
}
