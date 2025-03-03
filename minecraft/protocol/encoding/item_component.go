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

// Marks this item as damage resistant.
// The client won't render the item as
// being on-fire if this component is
// present.
type ItemComponentDamageResistant struct {
	// Tag specifying damage types the
	// item is immune to.
	// Not prefixed by '#'!.
	Types Identifier
}

func (i *ItemComponentDamageResistant) Name() string {
	return "minecraft:damage_resistant"
}

func (i *ItemComponentDamageResistant) Marshal(io IO) {
	Single(io, &i.Types)
}

// Alters the speed at which this
// item breaks certain blocks.
type ItemComponentTool struct {
	// Rule ..
	Rule []ItemComponentToolRule
	// The mining speed in case none of
	// the previous rule were matched.
	DefaultMiningSpeed float32
	// The amount of damage the item takes
	// per block break.
	DamagePerBlock int32
}

func (i *ItemComponentTool) Name() string {
	return "minecraft:tool"
}

func (i *ItemComponentTool) Marshal(io IO) {
	SliceVarint32Length(io, &i.Rule)
	io.Float32(&i.DefaultMiningSpeed)
	io.Varint32(&i.DamagePerBlock)
}

// Allows the item to be enchanted
// by an enchanting table.
type ItemComponentEnchantable struct {
	// Opaque internal value controlling how
	// expensive enchantments may be offered.
	Value int32
}

func (i *ItemComponentEnchantable) Name() string {
	return "minecraft:enchantable"
}

func (i *ItemComponentEnchantable) Marshal(io IO) {
	io.Varint32(&i.Value)
}

const (
	ItemSlotMainhand int32 = iota
	ItemSlotFeet
	ItemSlotLegs
	ItemSlotChest
	ItemSlotHead
	ItemSlotOffhand
	ItemSlotBody
)

// Allows the item to be equipped by the player.
type ItemComponentEquippable struct {
	// 0: mainhand
	// 1: feet
	// 2: legs
	// 3: chest
	// 4: head
	// 5: offhand
	// 6: body
	Slot int32
	// ID in the minecraft:sound_event registry,
	// or an inline definition.
	EquipSound IDOrX[SoundEvent]
	// Model ..
	Model Optional[Identifier]
	// CameraOverlay ..
	CameraOverlay Optional[Identifier]
	// IDs in the minecraft:entity_type registry.
	AllowedEntities Optional[IDSet]
	// Dispensable ..
	Dispensable bool
	// Swappable ..
	Swappable bool
	// DamageOnHurt ..
	DamageOnHurt bool
}

func (i *ItemComponentEquippable) Name() string {
	return "minecraft:equippable"
}

func (i *ItemComponentEquippable) Marshal(io IO) {
	io.Varint32(&i.Slot)
	IDOrXMarshaler(io, &i.EquipSound)
	OptionalMarshaler(io, &i.Model)
	OptionalMarshaler(io, &i.CameraOverlay)
	OptionalMarshaler(io, &i.AllowedEntities)
	io.Bool(&i.Dispensable)
	io.Bool(&i.Swappable)
	io.Bool(&i.DamageOnHurt)
}

// Items that can be combined with this
// item in an anvil to repair it.
type ItemComponentRepairable struct {
	// IDs in the minecraft:item registry.
	Items IDSet
}

func (i *ItemComponentRepairable) Name() string {
	return "minecraft:repairable"
}

func (i *ItemComponentRepairable) Marshal(io IO) {
	Single(io, &i.Items)
}

// Makes the item function like elytra.
type ItemComponentGlider struct{}

func (i *ItemComponentGlider) Name() string {
	return "minecraft:glider"
}

func (i *ItemComponentGlider) Marshal(io IO) {}

// Makes the item function like a totem of undying.
type ItemComponentDeathProtection struct {
	// Effects to apply on consumption.
	// See Consume Effect (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Slot_Data#Consume_Effect).
	Effects []ConsumeEffect
}

func (i *ItemComponentDeathProtection) Name() string {
	return "minecraft:death_protection"
}

func (i *ItemComponentDeathProtection) Marshal(io IO) {
	FuncSliceVarint32Length(io, &i.Effects, io.ConsumeEffect)
}

// The enchantments stored in this enchanted book.
type ItemComponentStoredEnchantments struct {
	// Enchantment ..
	Enchantment []Enchantment
	// Whether the list of enchantments should
	// be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentStoredEnchantments) Name() string {
	return "minecraft:stored_enchantments"
}

func (i *ItemComponentStoredEnchantments) Marshal(io IO) {
	SliceVarint32Length(io, &i.Enchantment)
	io.Bool(&i.ShowInTooltip)
}

// Color of dyed leather armor.
type ItemComponentDyedColor struct {
	// The RGB components of the color,
	// encoded as an integer.
	Color int32
	// Whether the armor's color should be
	// shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentDyedColor) Name() string {
	return "minecraft:dyed_color"
}

func (i *ItemComponentDyedColor) Marshal(io IO) {
	io.Int32(&i.Color)
	io.Bool(&i.ShowInTooltip)
}

// Color of the markings on the map item model.
type ItemComponentMapColor struct {
	// The RGB components of the color,
	// encoded as an integer.
	Color int32
}

func (i *ItemComponentMapColor) Name() string {
	return "minecraft:map_color"
}

func (i *ItemComponentMapColor) Marshal(io IO) {
	io.Int32(&i.Color)
}

// The ID of the map.
type ItemComponentMapID struct {
	// ID ..
	ID int32
}

func (i *ItemComponentMapID) Name() string {
	return "minecraft:map_id"
}

func (i *ItemComponentMapID) Marshal(io IO) {
	io.Varint32(&i.ID)
}

// Icons present on a map.
type ItemComponentMapDecorations struct {
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentMapDecorations) Name() string {
	return "minecraft:map_decorations"
}

func (i *ItemComponentMapDecorations) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}

const (
	MapPostProcessingLock int32 = iota
	MapPostProcessingScale
)

// Used internally by the client
// when expanding or locking a map.
//
// Display extra information on the
// item's tooltip when the component
// is present.
type ItemComponentMapPostProcessing struct {
	// Type of post processing. Can be either:
	// 		0 - Lock
	// 		1 - Scale
	Type int32
}

func (i *ItemComponentMapPostProcessing) Name() string {
	return "minecraft:map_post_processing"
}

func (i *ItemComponentMapPostProcessing) Marshal(io IO) {
	io.Varint32(&i.Type)
}

// Projectiles loaded into a charged crossbow.
type ItemComponentChargedProjectiles struct {
	// Projectiles ..
	// TODO
	Projectiles []any
}

func (i *ItemComponentChargedProjectiles) Name() string {
	return "minecraft:charged_projectiles"
}

func (i *ItemComponentChargedProjectiles) Marshal(io IO) {
	// TODO
}
