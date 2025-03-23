package packet_play

import (
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding"
	packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"
)

// Furnace property key ID enumerate
const (
	// Key: Fire icon (fuel left)
	//
	// Value: The enchantment's xp level requirement
	PropertyKeyFurnaceFireIcon int16 = iota
	// Key: Maximum fuel burn time
	//
	// Value: fuel burn time or 0 (in-game ticks)
	PropertyKeyFurnaceMaximumFuelBurnTime
	// Key: Progress arrow
	//
	// Value: counting from 0 to maximum progress (in-game ticks)
	PropertyKeyFurnaceProgressArrow
	// Key: Maximum progress
	//
	// Value: always 200 on the vanilla server
	PropertyKeyFurnaceMaximumProgress
)

// Enchantment table property Key ID enumerate
const (
	// Key: Level requirement for top enchantment slot
	//
	// Value: The enchantment's xp level requirement
	PropertyKeyEnchTableRequireTopSlot int16 = iota
	// Key: Level requirement for middle enchantment slot
	//
	// Value: The enchantment's xp level requirement
	PropertyKeyEnchTableRequireMiddleSlot
	// Key: Level requirement for bottom enchantment slot
	//
	// Value: The enchantment's xp level requirement
	PropertyKeyEnchTableRequireBottomSlot
	// Key: The enchantment seed
	//
	// Value:
	// 		- Used for drawing the enchantment names
	// 		  [in SGA (https://en.wikipedia.org/wiki/Standard_Galactic_Alphabet)]
	// 		  clientside.
	// 		  The same seed is used to calculate enchantments,
	// 		  but some of the data isn't sent to the client to
	// 		  prevent easily guessing the entire list
	// 		  (the seed value here is the regular seed bitwise
	// 		  and 0xFFFFFFF0).
	PropertyKeyEnchTableSeed
	// Key: Enchantment ID shown on mouse hover over top enchantment slot
	//
	// Value: The enchantment id (set to -1 to hide it), see below for values
	PropertyKeyEnchTableEnchIDTopSlot
	// Key: Enchantment ID shown on mouse hover over middle enchantment slot
	//
	// Value: The enchantment id (set to -1 to hide it), see below for values
	PropertyKeyEnchTableEnchIDMiddleSlot
	// Key: Enchantment ID shown on mouse hover over bottom enchantment slot
	//
	// Value: The enchantment id (set to -1 to hide it), see below for values
	PropertyKeyEnchTableEnchIDBottomSlot
	// Key: Enchantment level shown on mouse hover over the top slot
	//
	// Value: The enchantment level (1 = I, 2 = II, 6 = VI, etc.), or -1 if no enchant
	PropertyKeyEnchTableEnchLevelTopSlot
	// Key: Enchantment level shown on mouse hover over the middle slot
	//
	// Value: The enchantment level (1 = I, 2 = II, 6 = VI, etc.), or -1 if no enchant
	PropertyKeyEnchTableEnchLevelMiddleSlot
	// Key: Enchantment level shown on mouse hover over the bottom slot
	//
	// Value: The enchantment level (1 = I, 2 = II, 6 = VI, etc.), or -1 if no enchant
	PropertyKeyEnchTableEnchLevelBottomSlot
)

// Beacon property Key ID enumerate
//
// Helpful links:
//   - Potion effect ID (https://minecraft.wiki/w/Data_values#Status_effects)
const (
	// Key: Power level
	//
	// Value: 0-4, controls what effect buttons are enabled
	PropertyKeyBeaconPowerLevel int16 = iota
	// Key: First potion effect
	//
	// Value: Potion effect ID for the first effect, or -1 if no effect
	PropertyKeyBeaconFirstPotionEffect
	// Key: Second potion effect
	//
	// Value: Potion effect ID for the second effect, or -1 if no effect
	PropertyKeyBeaconSecondPotionEffect
)

// Anvil property Key ID enumerate
const (
	// Key: Repair cost
	//
	// Value: The repair's cost in xp levels
	PropertyKeyAnvilRepairCost int16 = iota
)

// Brewing stand property Key ID enumerate
const (
	// Key: Brew time
	//
	// Value: 0 – 400, with 400 making the arrow empty, and 0 making the arrow full
	PropertyKeyBrewingStandBrewTime int16 = iota
	// Key: Fuel time
	//
	// Value: 0 - 20, with 0 making the arrow empty, and 20 making the arrow full
	PropertyKeyBrewingStandFuelTime
)

// Stonecutter property Key ID enumerate
const (
	// Key: Selected recipe
	//
	// Value: The index of the selected recipe. -1 means none is selected
	PropertyKeyStonecutterSelectedRecipe int16 = iota
)

// Loom Key ID enumerate
const (
	// Key: Selected pattern
	//
	// Value:
	// 		- The index of the selected pattern. 0 means none is selected,
	// 		  0 is also the internal id of the "base" pattern.
	PropertyKeyLoomSelectedPattern int16 = iota
)

// Lectern Key ID enumerate
const (
	// Key: Page number
	//
	// Value: The current page number, starting from 0
	PropertyKeyLecternPageNumber int16 = iota
)

// This packet is used to inform the client
// that part of a GUI window should be updated.
type SetContainerProperty struct {
	// WindowID ..
	WindowID int32
	// The property to be updated,
	// see constant enum above.
	Property int16
	// The new value for the property,
	// see constant enum above.
	Value int16
}

// ID ..
func (p *SetContainerProperty) ID() int32 {
	return IDClientBoundSetContainerProperty
}

// Resource ..
func (p *SetContainerProperty) Resource() string {
	return "container_set_data"
}

// BoundType ..
func (p *SetContainerProperty) BoundType() uint8 {
	return packet_interface.BoundTypeClient
}

func (p *SetContainerProperty) Marshal(io encoding.IO) {
	io.Varint32(&p.WindowID)
	io.Int16(&p.Property)
	io.Int16(&p.Value)
}
