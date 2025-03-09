package encoding

const (
	RecipeDisplayTypeCraftingShapeless int32 = iota
	RecipeDisplayTypeCraftingShaped
	RecipeDisplayTypeFurnace
	RecipeDisplayTypeStonecutter
	RecipeDisplayTypeSmithing
)

// RecipeDisplay ..
type RecipeDisplay interface {
	Name() string
	Marshaler
}

// lookupRecipeDisplayType looks up
// the ID of a RecipeDisplay.
func lookupRecipeDisplayType(x RecipeDisplay, id *int32) bool {
	switch x.(type) {
	case *RecipeDisplayCraftingShapeless:
		*id = RecipeDisplayTypeCraftingShapeless
	case *RecipeDisplayCraftingShaped:
		*id = RecipeDisplayTypeCraftingShaped
	case *RecipeDisplayFurnace:
		*id = RecipeDisplayTypeFurnace
	case *RecipeDisplayStonecutter:
		*id = RecipeDisplayTypeStonecutter
	case *RecipeDisplaySmithing:
		*id = RecipeDisplayTypeSmithing
	default:
		return false
	}
	return true
}

// lookupRecipeDisplay looks up the
// RecipeDisplay matching an ID.
func lookupRecipeDisplay(id int32, x *RecipeDisplay) bool {
	switch id {
	case RecipeDisplayTypeCraftingShapeless:
		*x = &RecipeDisplayCraftingShapeless{}
	case RecipeDisplayTypeCraftingShaped:
		*x = &RecipeDisplayCraftingShaped{}
	case RecipeDisplayTypeFurnace:
		*x = &RecipeDisplayFurnace{}
	case RecipeDisplayTypeStonecutter:
		*x = &RecipeDisplayStonecutter{}
	case RecipeDisplayTypeSmithing:
		*x = &RecipeDisplaySmithing{}
	default:
		return false
	}
	return true
}

// RecipeDisplayCraftingShapeless ..
type RecipeDisplayCraftingShapeless struct {
	// Ingredients ..
	Ingredients []SlotDisplay
	// Result ..
	Result SlotDisplay
	// Crafting station icon to display
	// in the recipe notification.
	CraftingStation SlotDisplay
}

func (r *RecipeDisplayCraftingShapeless) Name() string {
	return "minecraft:crafting_shapeless"
}

func (r *RecipeDisplayCraftingShapeless) Marshal(io IO) {
	FuncSliceVarint32Length(io, &r.Ingredients, io.SlotDisplay)
	io.SlotDisplay(&r.Result)
	io.SlotDisplay(&r.CraftingStation)
}

// RecipeDisplayCraftingShaped ..
type RecipeDisplayCraftingShaped struct {
	// Width ..
	Width int32
	// Height ..
	Height int32
	// Note that len(Ingredients) = Width * Height.
	Ingredients []SlotDisplay
	// Result ..
	Result SlotDisplay
	// Crafting station icon to display
	// in the recipe notification.
	CraftingStation SlotDisplay
}

func (r *RecipeDisplayCraftingShaped) Name() string {
	return "minecraft:crafting_shaped"
}

func (r *RecipeDisplayCraftingShaped) Marshal(io IO) {
	io.Varint32(&r.Width)
	io.Varint32(&r.Height)
	FuncSliceVarint32Length(io, &r.Ingredients, io.SlotDisplay)
	io.SlotDisplay(&r.Result)
	io.SlotDisplay(&r.CraftingStation)
}

// RecipeDisplayFurnace ..
type RecipeDisplayFurnace struct {
	// Ingredient ..
	Ingredient SlotDisplay
	// Fuel ..
	Fuel SlotDisplay
	// Result ..
	Result SlotDisplay
	// Crafting station icon to display
	// in the recipe notification.
	CraftingStation SlotDisplay
	// Time to smelt in a regular furnace,
	// in ticks.
	CookingTime int32
	// Experience ..
	Experience float32
}

func (r *RecipeDisplayFurnace) Name() string {
	return "minecraft:furnace"
}

func (r *RecipeDisplayFurnace) Marshal(io IO) {
	io.SlotDisplay(&r.Ingredient)
	io.SlotDisplay(&r.Fuel)
	io.SlotDisplay(&r.Result)
	io.SlotDisplay(&r.CraftingStation)
	io.Varint32(&r.CookingTime)
	io.Float32(&r.Experience)
}

// RecipeDisplayStonecutter ..
type RecipeDisplayStonecutter struct {
	// Ingredient ..
	Ingredient SlotDisplay
	// Result ..
	Result SlotDisplay
	// Crafting station icon to display
	// in the recipe notification.
	CraftingStation SlotDisplay
}

func (r *RecipeDisplayStonecutter) Name() string {
	return "minecraft:stonecutter"
}

func (r *RecipeDisplayStonecutter) Marshal(io IO) {
	io.SlotDisplay(&r.Ingredient)
	io.SlotDisplay(&r.Result)
	io.SlotDisplay(&r.CraftingStation)
}

// RecipeDisplaySmithing ..
type RecipeDisplaySmithing struct {
	// Template ..
	Template SlotDisplay
	// Base ..
	Base SlotDisplay
	// Addition ..
	Addition SlotDisplay
	// Result ..
	Result SlotDisplay
	// Crafting station icon to display
	// in the recipe notification.
	CraftingStation SlotDisplay
}

func (r *RecipeDisplaySmithing) Name() string {
	return "minecraft:smithing"
}

func (r *RecipeDisplaySmithing) Marshal(io IO) {
	io.SlotDisplay(&r.Template)
	io.SlotDisplay(&r.Base)
	io.SlotDisplay(&r.Addition)
	io.SlotDisplay(&r.Result)
	io.SlotDisplay(&r.CraftingStation)
}
