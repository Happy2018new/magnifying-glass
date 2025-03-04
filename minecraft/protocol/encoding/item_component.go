package encoding

import "magnifying-glass/minecraft/nbt"

const (
	ItemComponentCustomDataType uint8 = iota
)

// ItemComponent is the structured components
// of Minecraft Java item stack.
// It is similar to the NBT data that record
// in Minecraft Bedrock item stack.
type ItemComponent interface {
	Name() string
	Marshaler
}

// lookupItemComponentType looks up
// the ID of a ItemComponent.
func lookupItemComponentType(x ItemComponent, id *uint8) bool {
	switch x.(type) {
	case *ItemComponentCustomData:
		*id = ItemComponentCustomDataType
	default:
		return false
	}
	return true
}

// lookupItemComponent looks up the
// ItemComponent matching an ID.
func lookupItemComponent(id uint8, x *ItemComponent) bool {
	switch id {
	case ItemComponentCustomDataType:
		*x = &ItemComponentCustomData{}
	default:
		return false
	}
	return true
}

// Customizable data that doesn't
// fit any specific component.
type ItemComponentCustomData struct {
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentCustomData) Name() string {
	return "minecraft:custom_data"
}

func (i *ItemComponentCustomData) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}

// Maximum stack size for the item.
type ItemComponentMaxStackSize struct {
	// Ranges from 1 to 99.
	MaxStackSize int32
}

func (i *ItemComponentMaxStackSize) Name() string {
	return "minecraft:max_stack_size"
}

func (i *ItemComponentMaxStackSize) Marshal(io IO) {
	io.Varint32(&i.MaxStackSize)
}

// The maximum damage the item can take before breaking.
type ItemComponentMaxDamage struct {
	MaxDamage int32
}

func (i *ItemComponentMaxDamage) Name() string {
	return "minecraft:max_damage"
}

func (i *ItemComponentMaxDamage) Marshal(io IO) {
	io.Varint32(&i.MaxDamage)
}

// The current damage of the item.
type ItemComponentDamage struct {
	Damage int32
}

func (i *ItemComponentDamage) Name() string {
	return "minecraft:damage"
}

func (i *ItemComponentDamage) Marshal(io IO) {
	io.Varint32(&i.Damage)
}

// Marks the item as unbreakable.
type ItemComponentUnbreakable struct {
	// Whether the Unbreakable indicator should
	// be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentUnbreakable) Name() string {
	return "minecraft:unbreakable"
}

func (i *ItemComponentUnbreakable) Marshal(io IO) {
	io.Bool(&i.ShowInTooltip)
}

// Item's custom name.
// Normally shown in italic,
// and changeable at an anvil.
type ItemComponentCustomName struct {
	CustomName TextComponentComplex
}

func (i *ItemComponentCustomName) Name() string {
	return "minecraft:custom_name"
}

func (i *ItemComponentCustomName) Marshal(io IO) {
	Single(io, &i.CustomName)
}

// Override for the item's default name.
// Shown when the item has no custom name.
type ItemComponentItemName struct {
	ItemName TextComponentComplex
}

func (i *ItemComponentItemName) Name() string {
	return "minecraft:item_name"
}

func (i *ItemComponentItemName) Marshal(io IO) {
	Single(io, &i.ItemName)
}

// Item's model.
type ItemComponentItemModel struct {
	Model Identifier
}

func (i *ItemComponentItemModel) Name() string {
	return "minecraft:item_model"
}

func (i *ItemComponentItemModel) Marshal(io IO) {
	Single(io, &i.Model)
}

// Item's lore.
type ItemComponentLore struct {
	Lines []TextComponentComplex
}

func (i *ItemComponentLore) Name() string {
	return "minecraft:lore"
}

func (i *ItemComponentLore) Marshal(io IO) {
	SliceVarint32Length(io, &i.Lines)
}

const (
	ItemComponentRarityCommon int32 = iota
	ItemComponentRarityUncommon
	ItemComponentRarityRare
	ItemComponentRarityEpic
)

// Item's rarity.
// This affects the default color of the item's name.
type ItemComponentRarity struct {
	// Can be one of the following:
	// 0 - Common (white)
	// 1 - Uncommon (yellow)
	// 2 - Rare (aqua)
	// 3 - Epic (pink)
	Rarity int32
}

func (i *ItemComponentRarity) Name() string {
	return "minecraft:rarity"
}

func (i *ItemComponentRarity) Marshal(io IO) {
	io.Varint32(&i.Rarity)
}

// The enchantments of the item.
type ItemComponentEnchantments struct {
	// Multiple enchantments of this item.
	Enchantments []Enchantment
	// Whether the list of enchantments should
	// be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentEnchantments) Name() string {
	return "minecraft:enchantments"
}

func (i *ItemComponentEnchantments) Marshal(io IO) {
	SliceVarint32Length(io, &i.Enchantments)
	io.Bool(&i.ShowInTooltip)
}

// List of blocks this block can be
// placed on when in adventure mode.
type ItemComponentCanPlaceOn struct {
	// See BlockPredicate for more details.
	BlockPredicates []BlockPredicate
	// Whether the Unbreakable indicator
	// should be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentCanPlaceOn) Name() string {
	return "minecraft:can_place_on"
}

func (i *ItemComponentCanPlaceOn) Marshal(io IO) {
	SliceVarint32Length(io, &i.BlockPredicates)
	io.Bool(&i.ShowInTooltip)
}

// List of blocks this item can
// break when in adventure mode.
type ItemComponentCanBreak struct {
	// See BlockPredicate for more details.
	BlockPredicates []BlockPredicate
	// Whether the Unbreakable indicator
	// should be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentCanBreak) Name() string {
	return "minecraft:can_break"
}

func (i *ItemComponentCanBreak) Marshal(io IO) {
	SliceVarint32Length(io, &i.BlockPredicates)
	io.Bool(&i.ShowInTooltip)
}

// The attribute modifiers of the item.
type ItemComponentAttributeModifiers struct {
	// AttributeModifier ..
	AttributeModifier []AttributeModifier
	// Whether the list of attribute modifiers
	// should be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentAttributeModifiers) Name() string {
	return "minecraft:attribute_modifiers"
}

func (i *ItemComponentAttributeModifiers) Marshal(io IO) {
	SliceVarint32Length(io, &i.AttributeModifier)
	io.Bool(&i.ShowInTooltip)
}

// Value for the item predicate when using custom item models.
// More info can be found here (https://minecraft.wiki/w/Tutorials/Models#Item_predicates).
type ItemComponentCustomModelData struct {
	// Floats ..
	Floats []float32
	// Flags ..
	Flags []bool
	// Strings ..
	Strings []string
	// Colors ..
	Colors []int32
}

func (i *ItemComponentCustomModelData) Name() string {
	return "minecraft:custom_model_data"
}

func (i *ItemComponentCustomModelData) Marshal(io IO) {
	FuncSliceVarint32Length(io, &i.Floats, io.Float32)
	FuncSliceVarint32Length(io, &i.Flags, io.Bool)
	FuncSliceVarint32Length(io, &i.Strings, io.String)
	FuncSliceVarint32Length(io, &i.Colors, io.Int32)
}

// Hides the special item's tooltip of crossbow ("Projectile:"),
// banner pattern layers, goat horn instrument and others.
type ItemComponentHideAdditionalTooltip struct{}

func (i *ItemComponentHideAdditionalTooltip) Name() string {
	return "minecraft:hide_additional_tooltip"
}

func (i *ItemComponentHideAdditionalTooltip) Marshal(io IO) {}

// Hides the item's tooltip altogether.
type ItemComponentHideTooltip struct{}

func (i *ItemComponentHideTooltip) Name() string {
	return "minecraft:hide_tooltip"
}

func (i *ItemComponentHideTooltip) Marshal(io IO) {}

// Accumulated anvil usage cost.
//
// The client displays "Too Expensive" if the value is
// greater than 40 and the player is not in creative mode
// (more specifically, if they don't have the insta-build flag enabled).
//
// This behavior can be overridden by setting the level with the
// Set Container Property packet.
//
// Helpful links:
// 		- insta-build flag enabled (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Player_Abilities_(clientbound))
//		- Set Container Property (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Set_Container_Property)
type ItemComponentRepairCost struct {
	// Cost ..
	Cost int32
}

func (i *ItemComponentRepairCost) Name() string {
	return "minecraft:repair_cost"
}

func (i *ItemComponentRepairCost) Marshal(io IO) {
	io.Varint32(&i.Cost)
}

// Marks the item as non-interactive on the creative
// inventory (the first 5 rows of items).
// This is used internally by the client on the paper
// icon in the saved hot-bars tab.
type ItemComponentCreativeSlotLock struct{}

func (i *ItemComponentCreativeSlotLock) Name() string {
	return "minecraft:creative_slot_lock"
}

func (i *ItemComponentCreativeSlotLock) Marshal(io IO) {}

// Overrides the item glint resulted from enchantments.
type ItemComponentEnchantmentGlintOverride struct {
	// HasGlint ..
	HasGlint bool
}

func (i *ItemComponentEnchantmentGlintOverride) Name() string {
	return "minecraft:enchantment_glint_override"
}

func (i *ItemComponentEnchantmentGlintOverride) Marshal(io IO) {
	io.Bool(&i.HasGlint)
}

// Marks the projectile as intangible
// (cannot be picked-up).
type ItemComponentIntangibleProjectile struct {
	// Always an empty Compound Tag.
	Empty map[string]any
}

func (i *ItemComponentIntangibleProjectile) Name() string {
	return "minecraft:intangible_projectile"
}

func (i *ItemComponentIntangibleProjectile) Marshal(io IO) {
	io.NBT(&i.Empty, nbt.NetworkBigEndian)
}

// Makes the item restore the player's
// hunger bar when consumed.
type ItemComponentFood struct {
	// Non-negative.
	Nutrition int32
	// How much saturation will be given
	// after consuming the item.
	SaturationModifier float32
	// Whether the item can always be eaten,
	// even at full hunger.
	CanAlwaysEat bool
}

func (i *ItemComponentFood) Name() string {
	return "minecraft:food"
}

func (i *ItemComponentFood) Marshal(io IO) {
	io.Varint32(&i.Nutrition)
	io.Float32(&i.SaturationModifier)
	io.Bool(&i.CanAlwaysEat)
}

const (
	ItemComponentConsumableAnimationNone int32 = iota
	ItemComponentConsumableAnimationEat
	ItemComponentConsumableAnimationDrink
	ItemComponentConsumableAnimationBlock
	ItemComponentConsumableAnimationBow
	ItemComponentConsumableAnimationSpear
	ItemComponentConsumableAnimationCrossbow
	ItemComponentConsumableAnimationSpyglass
	ItemComponentConsumableAnimationTootHorn
	ItemComponentConsumableAnimationBrush
)

// Makes the item consumable.
type ItemComponentConsumable struct {
	// How long it takes to consume the item.
	ConsumeSeconds float32
	// 0: none
	// 1: eat
	// 2: drink
	// 3: block
	// 4: bow
	// 5: spear
	// 6: crossbow
	// 7: spyglass
	// 8: toot_horn
	// 9: brush
	Animation int32
	// ID in the minecraft:sound_event registry,
	// or an inline definition.
	Sound IDOrX[SoundEvent]
	// HasConsumeParticles ..
	HasConsumeParticles bool
	// Effects to apply on consumption.
	// See ConsumeEffect for more information.
	Effects []ConsumeEffect
}

func (i *ItemComponentConsumable) Name() string {
	return "minecraft:consumable"
}

func (i *ItemComponentConsumable) Marshal(io IO) {
	io.Float32(&i.ConsumeSeconds)
	io.Varint32(&i.Animation)
	IDOrXMarshaler(io, &i.Sound)
	io.Bool(&i.HasConsumeParticles)
	FuncSliceVarint32Length(io, &i.Effects, io.ConsumeEffect)
}

// This specifies the item produced after using
// the current item.
// In the Notchian server, this is used for stews,
// which turn into bowls.
// TODO
type ItemComponentUseRemainder struct {
	// How long it takes to consume the item.
	Remainder any
}

func (i *ItemComponentUseRemainder) Name() string {
	return "minecraft:use_remainder"
}

func (i *ItemComponentUseRemainder) Marshal(io IO) {
	// TODO
}

// Cooldown to apply on use of the item.
type ItemComponentUseCooldown struct {
	// How long it takes to consume the item.
	Seconds float32
	// Group of items to apply the cooldown to.
	// Only present if Has cooldown group is true;
	// otherwise defaults to the item's identifier.
	CooldownGroup Optional[Identifier]
}

func (i *ItemComponentUseCooldown) Name() string {
	return "minecraft:use_cooldown"
}

func (i *ItemComponentUseCooldown) Marshal(io IO) {
	io.Float32(&i.Seconds)
	OptionalMarshaler(io, &i.CooldownGroup)
}
