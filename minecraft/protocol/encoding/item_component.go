package encoding

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/nbt"

	"github.com/google/uuid"
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
func lookupItemComponentType(x ItemComponent, id *int32) bool {
	switch x.(type) {
	case *ItemComponentCustomData:
		*id = ItemComponentTypeCustomData
	case *ItemComponentMaxStackSize:
		*id = ItemComponentTypeMaxStackSize
	case *ItemComponentMaxDamage:
		*id = ItemComponentTypeMaxDamage
	case *ItemComponentDamage:
		*id = ItemComponentTypeDamage
	case *ItemComponentUnbreakable:
		*id = ItemComponentTypeUnbreakable
	case *ItemComponentCustomName:
		*id = ItemComponentTypeCustomName
	case *ItemComponentItemName:
		*id = ItemComponentTypeItemName
	case *ItemComponentItemModel:
		*id = ItemComponentTypeItemModel
	case *ItemComponentLore:
		*id = ItemComponentTypeLore
	case *ItemComponentRarity:
		*id = ItemComponentTypeRarity
	case *ItemComponentEnchantments:
		*id = ItemComponentTypeEnchantments
	case *ItemComponentCanPlaceOn:
		*id = ItemComponentTypeCanPlaceOn
	case *ItemComponentCanBreak:
		*id = ItemComponentTypeCanBreak
	case *ItemComponentAttributeModifiers:
		*id = ItemComponentTypeAttributeModifiers
	case *ItemComponentCustomModelData:
		*id = ItemComponentTypeCustomModelData
	case *ItemComponentHideAdditionalTooltip:
		*id = ItemComponentTypeHideAdditionalTooltip
	case *ItemComponentHideTooltip:
		*id = ItemComponentTypeHideTooltip
	case *ItemComponentRepairCost:
		*id = ItemComponentTypeRepairCost
	case *ItemComponentCreativeSlotLock:
		*id = ItemComponentTypeCreativeSlotLock
	case *ItemComponentEnchantmentGlintOverride:
		*id = ItemComponentTypeEnchantmentGlintOverride
	case *ItemComponentIntangibleProjectile:
		*id = ItemComponentTypeIntangibleProjectile
	case *ItemComponentFood:
		*id = ItemComponentTypeFood
	case *ItemComponentConsumable:
		*id = ItemComponentTypeConsumable
	case *ItemComponentUseRemainder:
		*id = ItemComponentTypeUseRemainder
	case *ItemComponentUseCooldown:
		*id = ItemComponentTypeUseCooldown
	case *ItemComponentDamageResistant:
		*id = ItemComponentTypeDamageResistant
	case *ItemComponentTool:
		*id = ItemComponentTypeTool
	case *ItemComponentEnchantable:
		*id = ItemComponentTypeEnchantable
	case *ItemComponentEquippable:
		*id = ItemComponentTypeEquippable
	case *ItemComponentRepairable:
		*id = ItemComponentTypeRepairable
	case *ItemComponentGlider:
		*id = ItemComponentTypeGlider
	case *ItemComponentTooltipStyle:
		*id = ItemComponentTypeTooltipStyle
	case *ItemComponentDeathProtection:
		*id = ItemComponentTypeDeathProtection
	case *ItemComponentStoredEnchantments:
		*id = ItemComponentTypeStoredEnchantments
	case *ItemComponentDyedColor:
		*id = ItemComponentTypeDyedColor
	case *ItemComponentMapColor:
		*id = ItemComponentTypeMapColor
	case *ItemComponentMapID:
		*id = ItemComponentTypeMapID
	case *ItemComponentMapDecorations:
		*id = ItemComponentTypeMapDecorations
	case *ItemComponentMapPostProcessing:
		*id = ItemComponentTypeMapPostProcessing
	case *ItemComponentChargedProjectiles:
		*id = ItemComponentTypeChargedProjectiles
	case *ItemComponentBundleContents:
		*id = ItemComponentTypeBundleContents
	case *ItemComponentPotionContents:
		*id = ItemComponentTypePotionContents
	case *ItemComponentSuspiciousStewEffects:
		*id = ItemComponentTypeSuspiciousStewEffects
	case *ItemComponentWritableBookContent:
		*id = ItemComponentTypeWritableBookContent
	case *ItemComponentWrittenBookContent:
		*id = ItemComponentTypeWrittenBookContent
	case *ItemComponentTrim:
		*id = ItemComponentTypeTrim
	case *ItemComponentDebugStickState:
		*id = ItemComponentTypeDebugStickState
	case *ItemComponentEntityData:
		*id = ItemComponentTypeEntityData
	case *ItemComponentBucketEntityData:
		*id = ItemComponentTypeBucketEntityData
	case *ItemComponentBlockEntityData:
		*id = ItemComponentTypeBlockEntityData
	case *ItemComponentInstrument:
		*id = ItemComponentTypeInstrument
	case *ItemComponentOminousBottleAmplifier:
		*id = ItemComponentTypeOminousBottleAmplifier
	case *ItemComponentJukeboxPlayable:
		*id = ItemComponentTypeJukeboxPlayable
	case *ItemComponentRecipes:
		*id = ItemComponentTypeRecipes
	case *ItemComponentLodestoneTracker:
		*id = ItemComponentTypeLodestoneTracker
	case *ItemComponentFireworkExplosion:
		*id = ItemComponentTypeFireworkExplosion
	case *ItemComponentFireworks:
		*id = ItemComponentTypeFireworks
	case *ItemComponentProfile:
		*id = ItemComponentTypeProfile
	case *ItemComponentNoteBlockSound:
		*id = ItemComponentTypeNoteBlockSound
	case *ItemComponentBannerPatterns:
		*id = ItemComponentTypeBannerPatterns
	case *ItemComponentBaseColor:
		*id = ItemComponentTypeBaseColor
	case *ItemComponentPotDecorations:
		*id = ItemComponentTypePotDecorations
	case *ItemComponentContainer:
		*id = ItemComponentTypeContainer
	case *ItemComponentBlockState:
		*id = ItemComponentTypeBlockState
	case *ItemComponentBees:
		*id = ItemComponentTypeBees
	case *ItemComponentLock:
		*id = ItemComponentTypeLock
	case *ItemComponentContainerLoot:
		*id = ItemComponentTypeContainerLoot
	default:
		return false
	}
	return true
}

// lookupItemComponent looks up the
// ItemComponent matching an ID.
func lookupItemComponent(id int32, x *ItemComponent) bool {
	switch id {
	case ItemComponentTypeCustomData:
		*x = &ItemComponentCustomData{}
	case ItemComponentTypeMaxStackSize:
		*x = &ItemComponentMaxStackSize{}
	case ItemComponentTypeMaxDamage:
		*x = &ItemComponentMaxDamage{}
	case ItemComponentTypeDamage:
		*x = &ItemComponentDamage{}
	case ItemComponentTypeUnbreakable:
		*x = &ItemComponentUnbreakable{}
	case ItemComponentTypeCustomName:
		*x = &ItemComponentCustomName{}
	case ItemComponentTypeItemName:
		*x = &ItemComponentItemName{}
	case ItemComponentTypeItemModel:
		*x = &ItemComponentItemModel{}
	case ItemComponentTypeLore:
		*x = &ItemComponentLore{}
	case ItemComponentTypeRarity:
		*x = &ItemComponentRarity{}
	case ItemComponentTypeEnchantments:
		*x = &ItemComponentEnchantments{}
	case ItemComponentTypeCanPlaceOn:
		*x = &ItemComponentCanPlaceOn{}
	case ItemComponentTypeCanBreak:
		*x = &ItemComponentCanBreak{}
	case ItemComponentTypeAttributeModifiers:
		*x = &ItemComponentAttributeModifiers{}
	case ItemComponentTypeCustomModelData:
		*x = &ItemComponentCustomModelData{}
	case ItemComponentTypeHideAdditionalTooltip:
		*x = &ItemComponentHideAdditionalTooltip{}
	case ItemComponentTypeHideTooltip:
		*x = &ItemComponentHideTooltip{}
	case ItemComponentTypeRepairCost:
		*x = &ItemComponentRepairCost{}
	case ItemComponentTypeCreativeSlotLock:
		*x = &ItemComponentCreativeSlotLock{}
	case ItemComponentTypeEnchantmentGlintOverride:
		*x = &ItemComponentEnchantmentGlintOverride{}
	case ItemComponentTypeIntangibleProjectile:
		*x = &ItemComponentIntangibleProjectile{}
	case ItemComponentTypeFood:
		*x = &ItemComponentFood{}
	case ItemComponentTypeConsumable:
		*x = &ItemComponentConsumable{}
	case ItemComponentTypeUseRemainder:
		*x = &ItemComponentUseRemainder{}
	case ItemComponentTypeUseCooldown:
		*x = &ItemComponentUseCooldown{}
	case ItemComponentTypeDamageResistant:
		*x = &ItemComponentDamageResistant{}
	case ItemComponentTypeTool:
		*x = &ItemComponentTool{}
	case ItemComponentTypeEnchantable:
		*x = &ItemComponentEnchantable{}
	case ItemComponentTypeEquippable:
		*x = &ItemComponentEquippable{}
	case ItemComponentTypeRepairable:
		*x = &ItemComponentRepairable{}
	case ItemComponentTypeGlider:
		*x = &ItemComponentGlider{}
	case ItemComponentTypeTooltipStyle:
		*x = &ItemComponentTooltipStyle{}
	case ItemComponentTypeDeathProtection:
		*x = &ItemComponentDeathProtection{}
	case ItemComponentTypeStoredEnchantments:
		*x = &ItemComponentStoredEnchantments{}
	case ItemComponentTypeDyedColor:
		*x = &ItemComponentDyedColor{}
	case ItemComponentTypeMapColor:
		*x = &ItemComponentMapColor{}
	case ItemComponentTypeMapID:
		*x = &ItemComponentMapID{}
	case ItemComponentTypeMapDecorations:
		*x = &ItemComponentMapDecorations{}
	case ItemComponentTypeMapPostProcessing:
		*x = &ItemComponentMapPostProcessing{}
	case ItemComponentTypeChargedProjectiles:
		*x = &ItemComponentChargedProjectiles{}
	case ItemComponentTypeBundleContents:
		*x = &ItemComponentBundleContents{}
	case ItemComponentTypePotionContents:
		*x = &ItemComponentPotionContents{}
	case ItemComponentTypeSuspiciousStewEffects:
		*x = &ItemComponentSuspiciousStewEffects{}
	case ItemComponentTypeWritableBookContent:
		*x = &ItemComponentWritableBookContent{}
	case ItemComponentTypeWrittenBookContent:
		*x = &ItemComponentWrittenBookContent{}
	case ItemComponentTypeTrim:
		*x = &ItemComponentTrim{}
	case ItemComponentTypeDebugStickState:
		*x = &ItemComponentDebugStickState{}
	case ItemComponentTypeEntityData:
		*x = &ItemComponentEntityData{}
	case ItemComponentTypeBucketEntityData:
		*x = &ItemComponentBucketEntityData{}
	case ItemComponentTypeBlockEntityData:
		*x = &ItemComponentBlockEntityData{}
	case ItemComponentTypeInstrument:
		*x = &ItemComponentInstrument{}
	case ItemComponentTypeOminousBottleAmplifier:
		*x = &ItemComponentOminousBottleAmplifier{}
	case ItemComponentTypeJukeboxPlayable:
		*x = &ItemComponentJukeboxPlayable{}
	case ItemComponentTypeRecipes:
		*x = &ItemComponentRecipes{}
	case ItemComponentTypeLodestoneTracker:
		*x = &ItemComponentLodestoneTracker{}
	case ItemComponentTypeFireworkExplosion:
		*x = &ItemComponentFireworkExplosion{}
	case ItemComponentTypeFireworks:
		*x = &ItemComponentFireworks{}
	case ItemComponentTypeProfile:
		*x = &ItemComponentProfile{}
	case ItemComponentTypeNoteBlockSound:
		*x = &ItemComponentNoteBlockSound{}
	case ItemComponentTypeBannerPatterns:
		*x = &ItemComponentBannerPatterns{}
	case ItemComponentTypeBaseColor:
		*x = &ItemComponentBaseColor{}
	case ItemComponentTypePotDecorations:
		*x = &ItemComponentPotDecorations{}
	case ItemComponentTypeContainer:
		*x = &ItemComponentContainer{}
	case ItemComponentTypeBlockState:
		*x = &ItemComponentBlockState{}
	case ItemComponentTypeBees:
		*x = &ItemComponentBees{}
	case ItemComponentTypeLock:
		*x = &ItemComponentLock{}
	case ItemComponentTypeContainerLoot:
		*x = &ItemComponentContainerLoot{}
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
	io.TextComponentComplex(&i.CustomName)
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
	io.TextComponentComplex(&i.ItemName)
}

// Item's model.
type ItemComponentItemModel struct {
	Model Identifier
}

func (i *ItemComponentItemModel) Name() string {
	return "minecraft:item_model"
}

func (i *ItemComponentItemModel) Marshal(io IO) {
	io.Identifier(&i.Model)
}

// Item's lore.
type ItemComponentLore struct {
	Lines []TextComponentComplex
}

func (i *ItemComponentLore) Name() string {
	return "minecraft:lore"
}

func (i *ItemComponentLore) Marshal(io IO) {
	FuncSliceVarint32Length(io, &i.Lines, io.TextComponentComplex)
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
//   - insta-build flag enabled (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Player_Abilities_(clientbound))
//   - Set Container Property (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Protocol#Set_Container_Property)
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
	IDOrXFunc(io, &i.Sound, io.SoundEvent)
	io.Bool(&i.HasConsumeParticles)
	FuncSliceVarint32Length(io, &i.Effects, io.ConsumeEffect)
}

// This specifies the item produced after using
// the current item.
// In the Notchian server, this is used for stews,
// which turn into bowls.
type ItemComponentUseRemainder struct {
	// How long it takes to consume the item.
	Remainder ItemStack
}

func (i *ItemComponentUseRemainder) Name() string {
	return "minecraft:use_remainder"
}

func (i *ItemComponentUseRemainder) Marshal(io IO) {
	io.ItemStack(&i.Remainder)
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
	OptionalFunc(io, &i.CooldownGroup, io.Identifier)
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
	io.Identifier(&i.Types)
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
	IDOrXFunc(io, &i.EquipSound, io.SoundEvent)
	OptionalFunc(io, &i.Model, io.Identifier)
	OptionalFunc(io, &i.CameraOverlay, io.Identifier)
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

// Custom textures for the item tooltip.
type ItemComponentTooltipStyle struct {
	// Style ..
	Style Identifier
}

func (i *ItemComponentTooltipStyle) Name() string {
	return "minecraft:tooltip_style"
}

func (i *ItemComponentTooltipStyle) Marshal(io IO) {
	io.Identifier(&i.Style)
}

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
	Projectiles []ItemStack
}

func (i *ItemComponentChargedProjectiles) Name() string {
	return "minecraft:charged_projectiles"
}

func (i *ItemComponentChargedProjectiles) Marshal(io IO) {
	FuncSliceVarint32Length(io, &i.Projectiles, io.ItemStack)
}

// Contents of a bundle.
type ItemComponentBundleContents struct {
	// The projectiles.
	Items []ItemStack
}

func (i *ItemComponentBundleContents) Name() string {
	return "minecraft:bundle_contents"
}

func (i *ItemComponentBundleContents) Marshal(io IO) {
	FuncSliceVarint32Length(io, &i.Items, io.ItemStack)
}

// Visual and effects of a potion item.
type ItemComponentPotionContents struct {
	// Whether this potion has an
	// ID in the potion registry.
	//
	// If this field existed, it
	// means it has the default
	// effects associated with the
	// potion type.
	PotionID Optional[int32]
	// Whether this potion has a custom color.
	// If this field not existed, it means it
	// uses the default color associated with
	// the potion type.
	//
	// CustomColor is the RGB components of the
	// color, encoded as an integer.
	CustomColor Optional[int32]
	// Any custom effects the potion might have.
	// See Potion Effect (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Slot_Data#Potion_Effect).
	CustomEffects []PotionEffect
	// CustomName ..
	CustomName string
}

func (i *ItemComponentPotionContents) Name() string {
	return "minecraft:potion_contents"
}

func (i *ItemComponentPotionContents) Marshal(io IO) {
	OptionalFunc(io, &i.PotionID, io.Varint32)
	OptionalFunc(io, &i.CustomColor, io.Int32)
	SliceVarint32Length(io, &i.CustomEffects)
	io.String(&i.CustomName)
}

// Effects granted by a suspicious stew.
type ItemComponentSuspiciousStewEffects struct {
	// Effect ..
	Effect []Effect
}

func (i *ItemComponentSuspiciousStewEffects) Name() string {
	return "minecraft:suspicious_stew_effects"
}

func (i *ItemComponentSuspiciousStewEffects) Marshal(io IO) {
	SliceVarint32Length(io, &i.Effect)
}

// Content of a writable book.
type ItemComponentWritableBookContent struct {
	// Page ..
	Page []BookPage
}

func (i *ItemComponentWritableBookContent) Name() string {
	return "minecraft:writable_book_content"
}

func (i *ItemComponentWritableBookContent) Marshal(io IO) {
	SliceVarint32Length(io, &i.Page)
}

// Content of a writable book.
type ItemComponentWrittenBookContent struct {
	// The raw title of the book.
	RawTitle string
	// The title after going through chat filters.
	FilteredTitle Optional[string]
	// Author ..
	Author string
	// Generation ..
	Generation int32
	// Page ..
	Page []BookPage
	// Whether entity selectors have already been resolved.
	Resolved bool
}

func (i *ItemComponentWrittenBookContent) Name() string {
	return "minecraft:written_book_content"
}

func (i *ItemComponentWrittenBookContent) Marshal(io IO) {
	io.String(&i.RawTitle)
	OptionalFunc(io, &i.FilteredTitle, io.String)
	io.String(&i.Author)
	io.Varint32(&i.Generation)
	SliceVarint32Length(io, &i.Page)
	io.Bool(&i.Resolved)
}

// Armor's trim pattern and color.
type ItemComponentTrim struct {
	// ID in the minecraft:trim_material
	// registry, or an inline definition.
	TrimMaterial IDOrX[ArmorTrimMaterial]
	// ID in the minecraft:trim_pattern
	// registry, or an inline definition.
	TrimPattern IDOrX[ArmorTrimMaterial]
	// Whether the trim information should
	// be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentTrim) Name() string {
	return "minecraft:trim"
}

func (i *ItemComponentTrim) Marshal(io IO) {
	IDOrXMarshaler(io, &i.TrimMaterial)
	IDOrXMarshaler(io, &i.TrimPattern)
	io.Bool(&i.ShowInTooltip)
}

// State of the debug stick.
type ItemComponentDebugStickState struct {
	// States of previously interacted blocks.
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentDebugStickState) Name() string {
	return "minecraft:debug_stick_state"
}

func (i *ItemComponentDebugStickState) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}

// Data for the entity to be created from this item.
type ItemComponentEntityData struct {
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentEntityData) Name() string {
	return "minecraft:entity_data"
}

func (i *ItemComponentEntityData) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}

// Data of the entity contained in this bucket.
type ItemComponentBucketEntityData struct {
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentBucketEntityData) Name() string {
	return "minecraft:bucket_entity_data"
}

func (i *ItemComponentBucketEntityData) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}

// Data of the block entity to be created from this item.
type ItemComponentBlockEntityData struct {
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentBlockEntityData) Name() string {
	return "minecraft:block_entity_data"
}

func (i *ItemComponentBlockEntityData) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}

// The sound played when using a goat horn.
type ItemComponentInstrument struct {
	// ID in the minecraft:instrument
	// registry, or an inline definition.
	Instrument IDOrX[Instrument]
}

func (i *ItemComponentInstrument) Name() string {
	return "minecraft:instrument"
}

func (i *ItemComponentInstrument) Marshal(io IO) {
	IDOrXMarshaler(io, &i.Instrument)
}

// Amplifier for the effect of an ominous bottle.
type ItemComponentOminousBottleAmplifier struct {
	// Between 0 and 4.
	Amplifier int32
}

func (i *ItemComponentOminousBottleAmplifier) Name() string {
	return "minecraft:ominous_bottle_amplifier"
}

func (i *ItemComponentOminousBottleAmplifier) Marshal(io IO) {
	io.Varint32(&i.Amplifier)
}

// The song this item will play when inserted into a jukebox.
type ItemComponentJukeboxPlayable struct {
	// Whether the jukebox song is specified
	// directly, or just referenced by name.
	DirectMode bool
	// The name of the jukebox song in its
	// respective registry.
	// Only present if Direct Mode is false.
	JukeboxSongName Identifier
	// ID in the minecraft:jukebox_song registry.
	// Only present if Direct Mode is true.
	JukeboxSong IDOrX[JukeboxSong]
	// Whether the song should be shown on the item's tooltip.
	ShowInTooltip bool
}

func (i *ItemComponentJukeboxPlayable) Name() string {
	return "minecraft:jukebox_playable"
}

func (i *ItemComponentJukeboxPlayable) Marshal(io IO) {
	io.Bool(&i.DirectMode)
	if !i.DirectMode {
		io.Identifier(&i.JukeboxSongName)
	} else {
		IDOrXMarshaler(io, &i.JukeboxSong)
	}
	io.Bool(&i.ShowInTooltip)
}

// The recipes this knowledge book unlocks.
type ItemComponentRecipes struct {
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentRecipes) Name() string {
	return "minecraft:recipes"
}

func (i *ItemComponentRecipes) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}

// The lodestone this compass points to.
type ItemComponentLodestoneTracker struct {
	// Whether this lodestone points to a position,
	// otherwise it spins randomly.
	HasGlobalPosition bool
	// The dimension the compass points to.
	// Only present if Has Global Position is true.
	Dimension Identifier
	// The position the compass points to.
	// Only present if Has Global Position is true.
	Position BlockPos
	// Whether the component is removed when
	// the associated lodestone is broken.
	Tracked bool
}

func (i *ItemComponentLodestoneTracker) Name() string {
	return "minecraft:lodestone_tracker"
}

func (i *ItemComponentLodestoneTracker) Marshal(io IO) {
	io.Bool(&i.HasGlobalPosition)
	if i.HasGlobalPosition {
		io.Identifier(&i.Dimension)
		io.Position(&i.Position)
	}
	io.Bool(&i.Tracked)
}

// Properties of a firework star.
type ItemComponentFireworkExplosion struct {
	// See Firework Explosion.
	// (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Slot_Data#Firework_Explosion).
	Explosion FireworkExplosion
}

func (i *ItemComponentFireworkExplosion) Name() string {
	return "minecraft:firework_explosion"
}

func (i *ItemComponentFireworkExplosion) Marshal(io IO) {
	Single(io, &i.Explosion)
}

// The lodestone this compass points to.
type ItemComponentFireworks struct {
	// FlightDuration ..
	FlightDuration int32
	// See Firework Explosion
	// (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Slot_Data#Firework_Explosion).
	Explosions []FireworkExplosion
}

func (i *ItemComponentFireworks) Name() string {
	return "minecraft:fireworks"
}

func (i *ItemComponentFireworks) Marshal(io IO) {
	io.Int32(&i.FlightDuration)
	SliceVarint32Length(io, &i.Explosions)
}

// Game Profile of a player's head.
type ItemComponentProfile struct {
	// ProfileName ..
	ProfileName Optional[string]
	// UniqueID ..
	UniqueID Optional[uuid.UUID]
	// See PlayerProfileProperty for more information.
	Property []PlayerProfileProperty
}

func (i *ItemComponentProfile) Name() string {
	return "minecraft:profile"
}

func (i *ItemComponentProfile) Marshal(io IO) {
	OptionalFunc(io, &i.ProfileName, io.String)
	OptionalFunc(io, &i.UniqueID, io.UUID)
	SliceVarint32Length(io, &i.Property)
}

// Sound played by a note block when this
// player's head is placed on top of it.
type ItemComponentNoteBlockSound struct {
	// Sound ..
	Sound Identifier
}

func (i *ItemComponentNoteBlockSound) Name() string {
	return "minecraft:note_block_sound"
}

func (i *ItemComponentNoteBlockSound) Marshal(io IO) {
	io.Identifier(&i.Sound)
}

// Patterns of a banner or banner applied to a shield.
type ItemComponentBannerPatterns struct {
	// Layer ..
	Layer []BannerPatterns
}

func (i *ItemComponentBannerPatterns) Name() string {
	return "minecraft:banner_patterns"
}

func (i *ItemComponentBannerPatterns) Marshal(io IO) {
	SliceVarint32Length(io, &i.Layer)
}

// Base color of the banner applied to a shield.
type ItemComponentBaseColor struct {
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

func (i *ItemComponentBaseColor) Name() string {
	return "minecraft:base_color"
}

func (i *ItemComponentBaseColor) Marshal(io IO) {
	io.Varint32(&i.Color)
}

// Decorations on the four sides of a pot.
type ItemComponentPotDecorations struct {
	// The ID of the items in the item registry.
	Decorations []int32
}

func (i *ItemComponentPotDecorations) Name() string {
	return "minecraft:pot_decorations"
}

func (i *ItemComponentPotDecorations) Marshal(io IO) {
	FuncSliceVarint32Length(io, &i.Decorations, io.Varint32)
}

// Items inside a container of any type.
type ItemComponentContainer struct {
	// Items ..
	Items []ItemStack
}

func (i *ItemComponentContainer) Name() string {
	return "minecraft:container"
}

func (i *ItemComponentContainer) Marshal(io IO) {
	FuncSliceVarint32Length(io, &i.Items, io.ItemStack)
}

// State of a block.
type ItemComponentBlockState struct {
	// Property ..
	Property []BlockStates
}

func (i *ItemComponentBlockState) Name() string {
	return "minecraft:block_state"
}

func (i *ItemComponentBlockState) Marshal(io IO) {
	SliceVarint32Length(io, &i.Property)
}

// Bees inside a hive.
type ItemComponentBees struct {
	// Bee ..
	Bee []ItemComponentBeeData
}

func (i *ItemComponentBees) Name() string {
	return "minecraft:bees"
}

func (i *ItemComponentBees) Marshal(io IO) {
	SliceVarint32Length(io, &i.Bee)
}

// Bees inside a hive.
type ItemComponentLock struct {
	// Always a String Tag.
	Key string
}

func (i *ItemComponentLock) Name() string {
	return "minecraft:lock"
}

func (i *ItemComponentLock) Marshal(io IO) {
	io.NBTString(&i.Key, nbt.NetworkBigEndian)
}

// Bees inside a hive.
type ItemComponentContainerLoot struct {
	// Always a Compound Tag.
	Data map[string]any
}

func (i *ItemComponentContainerLoot) Name() string {
	return "minecraft:container_loot"
}

func (i *ItemComponentContainerLoot) Marshal(io IO) {
	io.NBT(&i.Data, nbt.NetworkBigEndian)
}
