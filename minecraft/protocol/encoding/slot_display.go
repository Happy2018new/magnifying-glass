package encoding

const (
	SlotDisplayTypeEmpty int32 = iota
	SlotDisplayTypeAnyFuel
	SlotDisplayTypeItem
	SlotDisplayTypeItemStack
	SlotDisplayTypeTag
	SlotDisplayTypeSmithingTrim
	SlotDisplayTypeWithRemainder
	SlotDisplayTypeComposite
)

// SlotDisplay ..
type SlotDisplay interface {
	Name() string
	Marshaler
}

// lookupSlotDisplayType looks up
// the ID of a SlotDisplay.
func lookupSlotDisplayType(x SlotDisplay, id *int32) bool {
	switch x.(type) {
	case *SlotDisplayEmpty:
		*id = SlotDisplayTypeEmpty
	case *SlotDisplayAnyFuel:
		*id = SlotDisplayTypeAnyFuel
	case *SlotDisplayItem:
		*id = SlotDisplayTypeItem
	case *SlotDisplayItemStack:
		*id = SlotDisplayTypeItemStack
	case *SlotDisplayTag:
		*id = SlotDisplayTypeTag
	case *SlotDisplaySmithingTrim:
		*id = SlotDisplayTypeSmithingTrim
	case *SlotDisplayWithRemainder:
		*id = SlotDisplayTypeWithRemainder
	case *SlotDisplayComposite:
		*id = SlotDisplayTypeComposite
	default:
		return false
	}
	return true
}

// lookupSlotDisplay looks up the
// SlotDisplay matching an ID.
func lookupSlotDisplay(id int32, x *SlotDisplay) bool {
	switch id {
	case SlotDisplayTypeEmpty:
		*x = &SlotDisplayEmpty{}
	case SlotDisplayTypeAnyFuel:
		*x = &SlotDisplayAnyFuel{}
	case SlotDisplayTypeItem:
		*x = &SlotDisplayItem{}
	case SlotDisplayTypeItemStack:
		*x = &SlotDisplayItemStack{}
	case SlotDisplayTypeTag:
		*x = &SlotDisplayTag{}
	case SlotDisplayTypeSmithingTrim:
		*x = &SlotDisplaySmithingTrim{}
	case SlotDisplayTypeWithRemainder:
		*x = &SlotDisplayWithRemainder{}
	case SlotDisplayTypeComposite:
		*x = &SlotDisplayComposite{}
	default:
		return false
	}
	return true
}

// SlotDisplay ..
type SlotDisplayEmpty struct{}

func (s *SlotDisplayEmpty) Name() string {
	return "minecraft:empty"
}

func (s *SlotDisplayEmpty) Marshal(io IO) {}

// SlotDisplayAnyFuel ..
type SlotDisplayAnyFuel struct{}

func (s *SlotDisplayAnyFuel) Name() string {
	return "minecraft:any_fuel"
}

func (s *SlotDisplayAnyFuel) Marshal(io IO) {}

// SlotDisplayItem ..
type SlotDisplayItem struct {
	// ID in the minecraft:item registry.
	ItemType int32
}

func (s *SlotDisplayItem) Name() string {
	return "minecraft:item"
}

func (s *SlotDisplayItem) Marshal(io IO) {
	io.Varint32(&s.ItemType)
}

// SlotDisplayItemStack ..
type SlotDisplayItemStack struct {
	// ItemStack ..
	ItemStack ItemStack
}

func (s *SlotDisplayItemStack) Name() string {
	return "minecraft:item_stack"
}

func (s *SlotDisplayItemStack) Marshal(io IO) {
	io.ItemStack(&s.ItemStack)
}

// SlotDisplayTag ..
type SlotDisplayTag struct {
	// Tag in the minecraft:item registry.
	// Not prefixed by '#'!.
	Tag Identifier
}

func (s *SlotDisplayTag) Name() string {
	return "minecraft:tag"
}

func (s *SlotDisplayTag) Marshal(io IO) {
	io.Identifier(&s.Tag)
}

// SlotDisplaySmithingTrim ..
type SlotDisplaySmithingTrim struct {
	// Base ..
	Base SlotDisplay
	// Material ..
	Material SlotDisplay
	// Pattern ..
	Pattern SlotDisplay
}

func (s *SlotDisplaySmithingTrim) Name() string {
	return "minecraft:smithing_trim"
}

func (s *SlotDisplaySmithingTrim) Marshal(io IO) {
	io.SlotDisplay(&s.Base)
	io.SlotDisplay(&s.Material)
	io.SlotDisplay(&s.Pattern)
}

// SlotDisplayWithRemainder ..
type SlotDisplayWithRemainder struct {
	// Ingredient ..
	Ingredient SlotDisplay
	// Remainder ..
	Remainder SlotDisplay
}

func (s *SlotDisplayWithRemainder) Name() string {
	return "minecraft:with_remainder"
}

func (s *SlotDisplayWithRemainder) Marshal(io IO) {
	io.SlotDisplay(&s.Ingredient)
	io.SlotDisplay(&s.Remainder)
}

// SlotDisplayComposite ..
type SlotDisplayComposite struct {
	// Options ..
	Options []SlotDisplay
}

func (s *SlotDisplayComposite) Name() string {
	return "minecraft:composite"
}

func (s *SlotDisplayComposite) Marshal(io IO) {
	FuncSliceVarint32Length(io, &s.Options, io.SlotDisplay)
}
