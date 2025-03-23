package encoding

// See the following links for more information.
// (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Registry_Data#Wolf_Variant).
type WolfVariant struct {
	// The texture for the wild version of this wolf.
	// The Notchian client uses the corresponding asset located at textures.
	WildTexture Identifier
	// The texture for the tamed version of this wolf.
	// The Notchian client uses the corresponding asset located at textures.
	TameTexture Identifier
	// The texture for the angry version of this wolf.
	// The Notchian client uses the corresponding asset located at textures.
	AngryTexture Identifier
	// Biomes in which this wolf can spawn in
	// (IDs in the minecraft:biome registry).
	Biomes IDSet
}

func (w *WolfVariant) Marshal(io IO) {
	io.Identifier(&w.WildTexture)
	io.Identifier(&w.TameTexture)
	io.Identifier(&w.AngryTexture)
	Single(io, &w.Biomes)
}

// See the following links for more information.
// (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Registry_Data#Painting_Variant).
type PaintingVariant struct {
	// The width of the painting, in blocks.
	Width int32
	// The height of the painting, in blocks.
	Height int32
	// The texture for the painting.
	// The Notchian client uses the corresponding
	// asset located at textures/painting.
	AssetID Identifier
	// The displayed title of the painting.
	Title Optional[TextComponent]
	// The displayed author of the painting.
	Author Optional[TextComponent]
}

func (p *PaintingVariant) Marshal(io IO) {
	io.Int32(&p.Width)
	io.Int32(&p.Height)
	io.Identifier(&p.AssetID)
	OptionalFunc(io, &p.Title, io.TextComponent)
	OptionalFunc(io, &p.Author, io.TextComponent)
}
