package encoding

import (
	"magnifying-glass/minecraft/nbt"

	"github.com/go-gl/mathgl/mgl32"
)

// TextComponentString reads a TextComponentString from the reader.
func (r *Reader) TextComponentString(x *TextComponentString) {
	r.String((*string)(x))
}

// TextComponentComplex reads a TextComponentComplex from the reader.
func (r *Reader) TextComponentComplex(x *TextComponentComplex) {
	r.NBT((*map[string]any)(x), nbt.NetworkBigEndian)
}

// TextComponentComplexOptional reads a TextComponentComplexOptional from the reader.
func (r *Reader) TextComponentComplexOptional(x *TextComponentComplexOptional) {
	r.Bool(&x.Existed)
	if x.Existed {
		r.NBT(&x.Data, nbt.NetworkBigEndian)
	}
}

// JsonTextComponent reads a JsonTextComponent from the reader.
func (r *Reader) JsonTextComponent(x *JsonTextComponent) {
	r.NBTString((*string)(x), nbt.NetworkBigEndian)
}

// Identifier reads an Identifier from the reader.
func (r *Reader) Identifier(x *Identifier) {
	r.String((*string)(x))
}

// ConsumeEffect reads an ConsumeEffect from the reader.
func (r *Reader) ConsumeEffect(x *ConsumeEffect) {
	var t int32
	r.Varint32(&t)
	if !lookupConsumeEffect(t, x) {
		r.UnknownEnumOption(t, "consume effect type")
	}
	(*x).Marshal(r)
}

// Position reads BlockPos by read a int64 from the underlying buffer.
func (r *Reader) Position(x *BlockPos) {
	var val uint64
	r.Uint64(&val)
	x[0] = int32(val >> 38)
	x[1] = int32(val << 52 >> 52)
	x[2] = int32(val << 26 >> 38)
}

// ItemComponent reads an ItemComponent from the reader.
func (r *Reader) ItemComponent(x *ItemComponent) {
	var t int32
	r.Varint32(&t)
	if !lookupItemComponent(t, x) {
		r.UnknownEnumOption(t, "item component type")
	}
	(*x).Marshal(r)
}

// ItemStack reads an ItemStack from the reader.
func (r *Reader) ItemStack(x *ItemStack) {
	r.Varint32(&x.ItemCount)
	if x.ItemCount == 0 {
		return
	}
	r.Varint32(&x.ItemID)
	r.Varint32(&x.AddComponentsCount)
	r.Varint32(&x.RemoveComponentsCount)
	FuncSliceOfLen(r, uint32(x.AddComponentsCount), &x.ComponentsToAdd, r.ItemComponent)
	FuncSliceOfLen(r, uint32(x.RemoveComponentsCount), &x.ComponentsToRemove, r.Varint32)
}

// Particle reads a Particle from the reader.
func (r *Reader) Particle(x *Particle) {
	var t int32
	r.Varint32(&t)
	if !lookupParticle(t, x) {
		r.UnknownEnumOption(t, "particle type")
	}
	(*x).Marshal(r)
}

// EntityMetadata reads an entity metadata map
// from the underlying buffer into map x.
func (r *Reader) EntityMetadata(x *EntityMetadata) {
	*x = make(EntityMetadata)

	for {
		var index uint8
		var dataType int32

		r.Uint8(&index)
		if index == EntityMetadataTagEnd {
			break
		}

		r.Varint32(&dataType)
		switch dataType {
		case EntityDataTypeByte:
			var v byte
			r.Uint8(&v)
			(*x)[index] = v
		case EntityDataTypeVarint32:
			var v int32
			r.Varint32(&v)
			(*x)[index] = v
		case EntityDataTypeVarint64:
			var v int64
			r.Varint64(&v)
			(*x)[index] = v
		case EntityDataTypeFloat32:
			var v float32
			r.Float32(&v)
			(*x)[index] = v
		case EntityDataTypeString:
			var v string
			r.String(&v)
			(*x)[index] = v
		case EntityDataTypeTextCompound:
			var v TextComponentComplex
			r.TextComponentComplex(&v)
			(*x)[index] = v
		case EntityDataTypeOptionalTextCompound:
			var v TextComponentComplexOptional
			r.TextComponentComplexOptional(&v)
			(*x)[index] = v
		case EntityDataTypeItemStack:
			var v ItemStack
			r.ItemStack(&v)
			(*x)[index] = v
		case EntityDataTypeBoolean:
			var v bool
			r.Bool(&v)
			(*x)[index] = v
		case EntityDataTypeRotations:
			var v EntityDataRotations
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypePosition:
			var v BlockPos
			r.Position(&v)
			(*x)[index] = v
		case EntityDataTypeOptionalPosition:
			var v EntityDataOptionalPosition
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeDirection:
			var v EntityDataDirection
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeOptionalUUID:
			var v EntityDataOptionalUUID
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeBlockState:
			var v EntityDataBlockState
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeOptionalBlockState:
			var v EntityDataOptionalBlockState
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeTagNBT:
			var v map[string]any
			r.NBT(&v, nbt.NetworkBigEndian)
			(*x)[index] = v
		case EntityDataTypeParticle:
			var v Particle
			r.Particle(&v)
			(*x)[index] = v
		case EntityDataTypeParticles:
			var v []Particle
			FuncSliceVarint32Length(r, &v, r.Particle)
			(*x)[index] = v
		case EntityDataTypeVillagerData:
			var v EntityDataVillagerData
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeOptionalVarint32:
			var v EntityDataOptionalVarint32
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypePose:
			var v EntityDataPose
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeCatVariant:
			var v EntityDataCatVariant
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeWolfVariant:
			var v EntityDataWolfVariant
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeFrogVariant:
			var v EntityDataForgVariant
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeOptionalGlobalPosition:
			var v EntityDataOptionalGlobalPosition
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypePaintingVariant:
			var v EntityDataPaintingVariant
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeSnifferState:
			var v EntityDataSnifferState
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeArmadilloState:
			var v EntityDataArmadilloState
			v.Marshal(r)
			(*x)[index] = v
		case EntityDataTypeVec3:
			var v mgl32.Vec3
			r.Vec3(&v)
			(*x)[index] = v
		case EntityDataTypeQuaternion:
			var v mgl32.Vec4
			r.Vec4(&v)
			(*x)[index] = v
		default:
			r.UnknownEnumOption(dataType, "entity metadata")
		}
	}
}

// SoundEvent reads a SoundEvent from the reader.
func (r *Reader) SoundEvent(x *SoundEvent) {
	r.Identifier(&x.SoundName)
	OptionalFunc(r, &x.FixedRange, r.Float32)
}
