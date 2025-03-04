package encoding

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
	Single(io, &a.ModifierID)
	io.Float64(&a.Value)
	io.Varint32(&a.Operation)
	io.Varint32(&a.Slot)
}
