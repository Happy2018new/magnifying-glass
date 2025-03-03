package encoding

import "magnifying-glass/minecraft/nbt"

// BlockProperty refer to the property
// of a Minecraft Java block.
type BlockProperty struct {
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

func (b *BlockProperty) Marshal(io IO) {
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
	// Whether this predicate is tied
	// to specific types of blocks.
	HasBlocks bool
	// IDs in the minecraft:block registry.
	// Only present if Has Blocks is true.
	Blocks IDSet
	// hether this predicate is tied to
	// specific properties of a block.
	HasProperties bool
	// Number of elements in the following array.
	// Only present if Has Properties is true.
	NumberOfProperties int32
	// See Property structure below.
	// Only present if Has Properties is true.
	Properties []BlockProperty
	// Whether this predicate is tied
	// to specific block entity data.
	HasNBT bool
	// Only present is Has NBT is true.
	NBT map[string]any
}

func (b *BlockPredicate) Marshal(io IO) {
	io.Bool(&b.HasBlocks)
	if b.HasBlocks {
		Single(io, &b.Blocks)
	}

	io.Bool(&b.HasProperties)
	if b.HasProperties {
		io.Varint32(&b.NumberOfProperties)
		SliceOfLen(io, uint32(b.NumberOfProperties), &b.Properties)
	}

	io.Bool(&b.HasNBT)
	if b.HasNBT {
		io.NBT(&b.NBT, nbt.NetworkBigEndian)
	}
}
