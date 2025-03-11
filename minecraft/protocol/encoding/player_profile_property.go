package encoding

// The Property field looks like response of
// Mojang API#UUID to Profile and Skin/Cape (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Mojang_API#UUID_to_Profile_and_Skin/Cape),
// except using the protocol format instead of JSON.
//
// That is, each player will usually have one property
// with Name being “textures” and Value being a
// base64-encoded JSON string, as documented at
// Mojang API#UUID to Profile and Skin/Cape
// (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Mojang_API#UUID_to_Profile_and_Skin/Cape).
//
// An empty properties array is also acceptable,
// and will cause clients to display the player with one
// of the two default skins depending their UUID
// (again, see the Mojang API page).
type PlayerProfileProperty struct {
	// Name ..
	Name string
	// Value ..
	Value string
	// Signature ..
	Signature Optional[string]
}

func (p *PlayerProfileProperty) Marshal(io IO) {
	io.String(&p.Name)
	io.String(&p.Value)
	OptionalFunc(io, &p.Signature, io.String)
}
