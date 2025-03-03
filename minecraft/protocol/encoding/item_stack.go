package encoding

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
	ComponentsToAdd any
	// Components that need to remove.
	// int32 refer to the type of
	// components.
	ComponentsToRemove []int32
}
