package encoding

import "magnifying-glass/minecraft/nbt"

// ------------------------- ArmorTrim -------------------------

// ArmorTrimMaterialOverride ..
type ArmorTrimMaterialOverride struct {
	// ArmorMaterialType ..
	ArmorMaterialType int32
	// OverridenAssetName ..
	OverridenAssetName string
}

func (a *ArmorTrimMaterialOverride) Marshal(io IO) {
	io.Varint32(&a.ArmorMaterialType)
	io.String(&a.OverridenAssetName)
}

// ArmorTrimMaterial ..
type ArmorTrimMaterial struct {
	// AssetName ..
	AssetName string
	// Ingredient ..
	Ingredient int32
	// ItemModelIndex ..
	ItemModelIndex float32
	// Override ..
	Override []ArmorTrimMaterialOverride
	// Description ..
	Description TextComponentComplex
}

func (a *ArmorTrimMaterial) Marshal(io IO) {
	io.String(&a.AssetName)
	io.Varint32(&a.Ingredient)
	io.Float32(&a.ItemModelIndex)
	SliceVarint32Length(io, &a.Override)
	io.TextComponentComplex(&a.Description)
}

// ------------------------- AttributeModifier -------------------------

const (
	AttributeModifierOperationAdd int32 = iota
	AttributeModifierOperationMultiplyBase
	AttributeModifierOperationMultiplyTotal
)

const (
	AttributeModifierSlotAny int32 = iota
	AttributeModifierSlotMainhand
	AttributeModifierSlotOffhand
	AttributeModifierSlotHand
	AttributeModifierSlotFeet
	AttributeModifierSlotLegs
	AttributeModifierSlotChest
	AttributeModifierSlotHead
	AttributeModifierSlotArmor
	AttributeModifierSlotBody
)

// AttributeModifier ..
type AttributeModifier struct {
	// The attribute to be modified
	// (ID in the minecraft:attribute registry).
	AttributeID int32
	// The modifier's unique ID.
	ModifierID Identifier
	// The modifier's value.
	Value float64
	// The operation to be applied upon the value.
	// Can be one of the following:
	// 		0 - Add
	// 		1 - Multiply base
	// 		2 - Multiply total
	Operation int32
	// The item slot placement required for the
	// modifier to have effect.
	// Can be one of the following:
	// 		0 - Any
	// 		1 - Main hand
	// 		2 - Off hand
	// 		3 - Hand
	// 		4 - Feet
	// 		5 - Legs
	// 		6 - Chest
	// 		7 - Head
	// 		8 - Armor
	// 		9 - Body
	Slot int32
}

func (a *AttributeModifier) Marshal(io IO) {
	io.Varint32(&a.AttributeID)
	io.Identifier(&a.ModifierID)
	io.Float64(&a.Value)
	io.Varint32(&a.Operation)
	io.Varint32(&a.Slot)
}

// ------------------------- BannerPatterns -------------------------

const BannerPatternsTypeDirectly int32 = 0

// BannerPatterns ..
type BannerPatterns struct {
	// Identifier used to determine the data that follows.
	// It can be either:
	// 		0 - Directly represents a pattern, with the necessary data following.
	// 		Anything else - References a pattern in its registry, by the ID of Pattern Type - 1.
	PatternType int32
	// Identifier of the asset.
	// Only present if Pattern Type is 0.
	AssetID Identifier
	// Only present if Pattern Type is 0.
	TranslationKey string
	// Dye color.
	// Can be one of the following:
	// 		0 - White
	// 		1 - Orange
	// 		2 - Magenta
	// 		3 - Light Blue
	// 		4 - Yellow
	// 		5 - Lime
	// 		6 - Pink
	// 		7 - Gray
	// 		8 - Light Gray
	// 		9 - Cyan
	// 		10 - Purple
	// 		11 - Blue
	// 		12 - Brown
	// 		13 - Green
	// 		14 - Red
	// 		15 - Black
	Color int32
}

func (b *BannerPatterns) Marshal(io IO) {
	io.Varint32(&b.PatternType)
	if b.PatternType == BannerPatternsTypeDirectly {
		io.Identifier(&b.AssetID)
		io.String(&b.TranslationKey)
	}
	io.Varint32(&b.Color)
}

// ------------------------- Bee -------------------------

// ItemComponentBeeData ..
type ItemComponentBeeData struct {
	// Custom data for the entity, always a Compound Tag.
	// Same structure as the minecraft:custom_data component.
	EntityData map[string]any
	// TicksInHive ..
	TicksInHive int32
	// MinTicksInHive ..
	MinTicksInHive int32
}

func (i *ItemComponentBeeData) Marshal(io IO) {
	io.NBT(&i.EntityData, nbt.NetworkBigEndian)
	io.Varint32(&i.TicksInHive)
	io.Varint32(&i.MinTicksInHive)
}

// ------------------------- Book -------------------------

// BookPage refer to the book
// page data of a book item.
type BookPage struct {
	// The raw text of the page.
	RawContent string
	// The content after passing through chat filters.
	FilteredContent Optional[string]
}

func (b *BookPage) Marshal(io IO) {
	io.String(&b.RawContent)
	OptionalFunc(io, &b.FilteredContent, io.String)
}

// ------------------------- Firework -------------------------

const (
	FireworkShapeSmallBall int32 = iota
	FireworkShapeLargeBall
	FireworkShapeStar
	FireworkShapeCreeper
	FireworkShapeBurst
)

// Represents a firework explosion,
// consisting of a shape, colors,
// and extra details.
type FireworkExplosion struct {
	// Can be one of the following:
	// 		0 - Small ball
	// 		1 - Large ball
	// 		2 - Star
	// 		3 - Creeper
	// 		4 - Burst
	Shape int32
	// The RGB components of the color,
	// encoded as an integer.
	Colors []int32
	// The RGB components of the color,
	// encoded as an integer.
	FadeColors []int32
	// HasTrail ..
	HasTrail bool
	// HasTwinkle ..
	HasTwinkle bool
}

func (f *FireworkExplosion) Marshal(io IO) {
	io.Varint32(&f.Shape)
	FuncSliceVarint32Length(io, &f.Colors, io.Int32)
	FuncSliceVarint32Length(io, &f.FadeColors, io.Int32)
	io.Bool(&f.HasTrail)
	io.Bool(&f.HasTwinkle)
}

// ------------------------- Instrument -------------------------

// Instrument ..
type Instrument struct {
	// The sound to be played.
	SoundEvent IDOrX[SoundEvent]
	// The maximum range of the sound.
	UseDurationa float32
	// The range of the instrument.
	Range float32
	// Description shown in the item tooltip.
	Description TextComponentComplex
}

func (i *Instrument) Marshal(io IO) {
	IDOrXFunc(io, &i.SoundEvent, io.SoundEvent)
	io.Float32(&i.UseDurationa)
	io.Float32(&i.Range)
	io.TextComponentComplex(&i.Description)
}

// ------------------------- ItemToolRule -------------------------

// ItemToolRule ..
type ItemComponentToolRule struct {
	// The blocks this rule applies to
	// (IDs in the minecraft:block registry).
	Blocks IDSet
	// The speed at which the tool breaks this rules' blocks.
	Speed Optional[float32]
	// Whether items should drop only if this is the correct tool.
	CorrectDropForBlocks Optional[bool]
}

func (i *ItemComponentToolRule) Marshal(io IO) {
	Single(io, &i.Blocks)
	OptionalFunc(io, &i.Speed, io.Float32)
	OptionalFunc(io, &i.CorrectDropForBlocks, io.Bool)
}

// ------------------------- JukeboxSong -------------------------

// JukeboxSong refer to a Minecraft song
// data that can be played on Jukebox.
type JukeboxSong struct {
	// The sound to be played.
	SoundEvent IDOrX[SoundEvent]
	// The description shown in the item lore.
	Description TextComponentComplex
	// The duration the songs should play for, in seconds.
	Duration float32
	// The output strength given by a comparator.
	// Between 0 and 15.
	Output int32
}

func (j *JukeboxSong) Marshal(io IO) {
	IDOrXFunc(io, &j.SoundEvent, io.SoundEvent)
	io.TextComponentComplex(&j.Description)
	io.Float32(&j.Duration)
	io.Varint32(&j.Output)
}

// ------------------------- PlayerProfile -------------------------

// PlayerProfileProperty ..
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

// ------------------------- END -------------------------
