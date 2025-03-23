package encoding

import "github.com/Happy2018new/magnifying-glass/minecraft/nbt"

// BlockStates refer to the property
// of a Minecraft Java block.
//
// TODO: I don't what's the difference
// between this and BlockPredicateProperty.
type BlockStates struct {
	// Name ..
	Name string
	// Value ..
	Value string
}

func (b *BlockStates) Marshal(io IO) {
	io.String(&b.Name)
	io.String(&b.Value)
}

// BlockPredicateProperty refer to the
// property of a Minecraft Java block.
//
// TODO: I don't what's the difference
// between this and BlockStates.
type BlockPredicateProperty struct {
	// Name of the block state property.
	Name string
	// Whether this is an exact value match,
	// as opposed to ranged.
	IsExactMatch bool
	// Value of the block state property.
	// Only present in exact match mode.
	ExactValue string
	// Minimum value of the block state property range.
	// Only present in ranged match mode.
	MinValue string
	// Maximum value of the block state property range.
	// Only present in ranged match mode.
	MaxValue string
}

func (b *BlockPredicateProperty) Marshal(io IO) {
	io.String(&b.Name)
	io.Bool(&b.IsExactMatch)
	if b.IsExactMatch {
		io.String(&b.ExactValue)
	} else {
		io.String(&b.MinValue)
		io.String(&b.MaxValue)
	}
}

// Describes a predicate used when block filtering is necessary.
//
// It can be parameterized to account for the type of block,
// the values of specific block state properties, NBT data
// related to block entities, or any combination of the three.
type BlockPredicate struct {
	// IDs in the minecraft:block registry.
	// Only present if Has Blocks is true.
	Blocks Optional[IDSet]
	// See Property structure below.
	// Only present if Has Properties is true.
	Properties Optional[[]BlockPredicateProperty]
	// Only present is Has NBT is true.
	NBT Optional[map[string]any]
}

func (b *BlockPredicate) Marshal(io IO) {
	OptionalMarshaler(io, &b.Blocks)
	OptionalSliceMarshaler(io, &b.Properties)
	OptionalFunc(io, &b.NBT, func(t *map[string]any) { io.NBT(t, nbt.NetworkBigEndian) })
}
