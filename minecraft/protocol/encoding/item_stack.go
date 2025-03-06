package encoding

const (
	ItemComponentTypeCustomData int32 = iota
	ItemComponentTypeMaxStackSize
	ItemComponentTypeMaxDamage
	ItemComponentTypeDamage
	ItemComponentTypeUnbreakable
	ItemComponentTypeCustomName
	ItemComponentTypeItemName
	ItemComponentTypeItemModel
	ItemComponentTypeLore
	ItemComponentTypeRarity
	ItemComponentTypeEnchantments
	ItemComponentTypeCanPlaceOn
	ItemComponentTypeCanBreak
	ItemComponentTypeAttributeModifiers
	ItemComponentTypeCustomModelData
	ItemComponentTypeHideAdditionalTooltip
	ItemComponentTypeHideTooltip
	ItemComponentTypeRepairCost
	ItemComponentTypeCreativeSlotLock
	ItemComponentTypeEnchantmentGlintOverride
	ItemComponentTypeIntangibleProjectile
	ItemComponentTypeFood
	ItemComponentTypeConsumable
	ItemComponentTypeUseRemainder
	ItemComponentTypeUseCooldown
	ItemComponentTypeDamageResistant
	ItemComponentTypeTool
	ItemComponentTypeEnchantable
	ItemComponentTypeEquippable
	ItemComponentTypeRepairable
	ItemComponentTypeGlider
	ItemComponentTypeTooltipStyle
	ItemComponentTypeDeathProtection
	ItemComponentTypeStoredEnchantments
	ItemComponentTypeDyedColor
	ItemComponentTypeMapColor
	ItemComponentTypeMapID
	ItemComponentTypeMapDecorations
	ItemComponentTypeMapPostProcessing
	ItemComponentTypeChargedProjectiles
	ItemComponentTypeBundleContents
	ItemComponentTypePotionContents
	ItemComponentTypeSuspiciousStewEffects
	ItemComponentTypeWritableBookContent
	ItemComponentTypeWrittenBookContent
	ItemComponentTypeTrim
	ItemComponentTypeDebugStickState
	ItemComponentTypeEntityData
	ItemComponentTypeBucketEntityData
	ItemComponentTypeBlockEntityData
	ItemComponentTypeInstrument
	ItemComponentTypeOminousBottleAmplifier
	ItemComponentTypeJukeboxPlayable
	ItemComponentTypeRecipes
	ItemComponentTypeLodestoneTracker
	ItemComponentTypeFireworkExplosion
	ItemComponentTypeFireworks
	ItemComponentTypeProfile
	ItemComponentTypeNoteBlockSound
	ItemComponentTypeBannerPatterns
	ItemComponentTypeBaseColor
	ItemComponentTypePotDecorations
	ItemComponentTypeContainer
	ItemComponentTypeBlockState
	ItemComponentTypeBees
	ItemComponentTypeLock
	ItemComponentTypeContainerLoot
)

// The Slot (ItemStack) data structure defines
// how an item is represented when inside an
// inventory window of any kind, such as a chest
// or furnace.
type ItemStack struct {
	// The item count. Every following field is only
	// present if this value is greater than zero.
	ItemCount int32
	// The item ID. Item IDs are distinct from block IDs.
	// See Data Generators (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Data_Generators)
	// for more information.
	ItemID int32
	// Number of elements present in the first data component array.
	AddComponentsCount int32
	// Number of elements present in the second data component array.
	// This serve as a way to remove the default component values
	// that are present on some items.
	RemoveComponentsCount int32
	// Components that needs to add.
	ComponentsToAdd []ItemComponent
	// Components that need to remove.
	// int32 refer to the type of
	// each component.
	// See the constant enum above for
	// more details.
	ComponentsToRemove []int32
}
