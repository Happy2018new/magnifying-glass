package encoding

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"math/big"
	"slices"
	"unsafe"

	"github.com/Happy2018new/magnifying-glass/minecraft/nbt"
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding/basic_encoding"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// Reader implements reading operations for
// reading types from Minecraft packets.
//
// Each Packet implementation
// has one passed to it.
//
// Reader's uses should always be encapsulated
// with a deferred recovery.
// Reader panics on invalid data.
type Reader struct {
	*basic_encoding.BasicReader
}

// NewReader creates a new Reader using the
// io.ByteReader passed as underlying source
// to read bytes from.
func NewReader(r interface {
	io.Reader
	io.ByteReader
}) *Reader {
	return &Reader{basic_encoding.NewBasicReader(r)}
}

// Bool reads a bool from the underlying buffer.
func (r *Reader) Bool(x *bool) {
	u, err := r.Reader().ReadByte()
	if err != nil {
		r.panic(err)
	}
	*x = *(*bool)(unsafe.Pointer(&u))
}

// String reads a string from the underlying buffer.
func (r *Reader) String(x *string) {
	var length int16
	r.Varint16(&length)
	l := int(length)
	data := make([]byte, l)
	if _, err := r.Reader().Read(data); err != nil {
		r.panic(err)
	}
	*x = *(*string)(unsafe.Pointer(&data))
}

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

// Position reads BlockPos by read a int64 from the underlying buffer.
func (r *Reader) Position(x *BlockPos) {
	var val uint64
	r.Uint64(&val)
	x[0] = int32(val >> 38)
	x[1] = int32(val << 52 >> 52)
	x[2] = int32(val << 26 >> 38)
}

// Angle reads a rotational float32 from a single byte.
func (r *Reader) Angle(x *float32) {
	var v uint8
	r.Uint8(&v)
	*x = float32(v) * (360.0 / 256.0)
}

// UUID reads a uuid.UUID from the underlying buffer.
func (r *Reader) UUID(x *uuid.UUID) {
	b := make([]byte, 16)
	if _, err := r.Reader().Read(b); err != nil {
		r.panic(err)
	}
	*x = uuid.UUID(b)
}

// Bitset reads a Java standard bitset that is []uint64
// with prefixed varint64 as the length of this slice
// into x. The encoding is little-endian.
func (r *Reader) Bitset(x *Bitset) {
	var bitsSliceOrigin []uint64
	FuncSliceVarint32Length(r, &bitsSliceOrigin, r.Uint64)

	length := len(bitsSliceOrigin)
	bitsSlice := make([]big.Word, length)
	for i := range length {
		bitsSlice[i] = big.Word(bitsSliceOrigin[i])
	}

	*x = NewBitset(length * 8)
	x.bits.SetBits(bitsSlice)
}

// FixedBitset reads a Minecraft Java fixed Bitset
// that is []byte into x.
// The size of []byte is giving by size, which meet
// x.Len() = celi(size / 8).
// The encoding is little-endian.
func (r *Reader) FixedBitset(x *Bitset, size int) {
	fixedSize := uint32(math.Ceil(float64(size) / 8))
	bitsSlice := make([]byte, fixedSize)

	r.Reader().Read(bitsSlice)
	slices.Reverse(bitsSlice)

	*x = NewBitset(size)
	x.bits.SetBytes(bitsSlice)
}

// SoundEvent reads a SoundEvent from the reader.
func (r *Reader) SoundEvent(x *SoundEvent) {
	r.Identifier(&x.SoundName)
	OptionalFunc(r, &x.FixedRange, r.Float32)
}

// TeleportFlags reads a TeleportFlags from the reader.
func (r *Reader) TeleportFlags(x *TeleportFlags) {
	r.FixedBitset((*Bitset)(x), TeleportFlagBitsetSize)
}

// RecipeDisplay reads a RecipeDisplay from the reader.
func (r *Reader) RecipeDisplay(x *RecipeDisplay) {
	var t int32
	r.Varint32(&t)
	if !lookupRecipeDisplay(t, x) {
		r.UnknownEnumOption(t, "recipe display type")
	}
	(*x).Marshal(r)
}

// SlotDisplay reads a SlotDisplay from the reader.
func (r *Reader) SlotDisplay(x *SlotDisplay) {
	var t int32
	r.Varint32(&t)
	if !lookupSlotDisplay(t, x) {
		r.UnknownEnumOption(t, "slot display type")
	}
	(*x).Marshal(r)
}

// ChunkData reads a ChunkData from the reader.
func (r *Reader) ChunkData(x *ChunkData) {
	r.NBT(&x.Heightmaps, nbt.NetworkBigEndian)
	FuncSliceVarint32Length(r, &x.Data, r.Uint8)
	SliceVarint32Length(r, &x.BlockEntities)
}

// LightData reads a LightData from the reader.
func (r *Reader) LightData(x *LightData) {
	r.Bitset(&x.SkyLightMask)
	r.Bitset(&x.BlockLightMask)
	r.Bitset(&x.EmptySkyLightMask)
	r.Bitset(&x.EmptyBlockLightMask)
	FuncSliceVarint32Length(r, &x.SkyLightArrays, r.Uint8)
	FuncSliceVarint32Length(r, &x.BlockLightArrays, r.Uint8)
}

// NBT reads a compound tag into a map from the underlying buffer.
func (r *Reader) NBT(m *map[string]any, encoding nbt.Encoding) {
	dec := nbt.NewDecoderWithEncoding(r.Reader(), encoding)
	dec.AllowZero = true

	*m = make(map[string]any)
	if err := dec.Decode(m); err != nil {
		r.panic(err)
	}
}

// NBTList reads a list of NBT tags from the underlying buffer.
func (r *Reader) NBTList(m *[]any, encoding nbt.Encoding) {
	if err := nbt.NewDecoderWithEncoding(r.Reader(), encoding).Decode(m); err != nil {
		r.panic(err)
	}
}

// NBTString reads a string tag into a string from the underlying buffer.
func (r *Reader) NBTString(s *string, encoding nbt.Encoding) {
	dec := nbt.NewDecoderWithEncoding(r.Reader(), encoding)
	dec.AllowZero = true

	*s = ""
	if err := dec.Decode(s); err != nil {
		r.panic(err)
	}
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

// ItemComponent reads an ItemComponent from the reader.
func (r *Reader) ItemComponent(x *ItemComponent) {
	var t int32
	r.Varint32(&t)
	if !lookupItemComponent(t, x) {
		r.UnknownEnumOption(t, "item component type")
	}
	(*x).Marshal(r)
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

// Vec3 reads three float32s into an mgl32.Vec3 from the underlying buffer.
func (r *Reader) Vec3(x *mgl32.Vec3) {
	r.Float32(&x[0])
	r.Float32(&x[1])
	r.Float32(&x[2])
}

// Vec4 reads four float32s into an mgl32.Vec4 from the underlying buffer.
func (r *Reader) Vec4(x *mgl32.Vec4) {
	r.Float32(&x[0])
	r.Float32(&x[1])
	r.Float32(&x[2])
	r.Float32(&x[3])
}

// RGB reads a color.RGBA x from a 0xRRGGBB uint32.
func (r *Reader) RGB(x *color.RGBA) {
	var v uint32
	r.Uint32(&v)
	*x = color.RGBA{
		R: byte((v >> 16) & 0xff),
		G: byte((v >> 8) & 0xff),
		B: byte(v & 0xff),
		A: 255,
	}
}

// RGBA reads a color.RGBA x from a 0xAARRGGBB uint32.
func (r *Reader) RGBA(x *color.RGBA) {
	var v uint32
	r.Uint32(&v)
	*x = color.RGBA{
		A: byte((v >> 24) & 0xff),
		R: byte((v >> 16) & 0xff),
		G: byte((v >> 8) & 0xff),
		B: byte(v & 0xff),
	}
}

// Bytes reads the leftover bytes into a byte slice.
func (r *Reader) Bytes(p *[]byte) {
	var err error
	*p, err = io.ReadAll(r.Reader())
	if err != nil {
		r.panic(err)
	}
}

// UnknownEnumOption panics with an unknown enum option error.
func (r *Reader) UnknownEnumOption(value any, enum string) {
	r.panicf("unknown value '%v' for enum type '%v'", value, enum)
}

// InvalidValue panics with an error indicating that the value passed is not valid for a specific field.
func (r *Reader) InvalidValue(value any, forField, reason string) {
	r.panicf("invalid value '%v' for %v: %v", value, forField, reason)
}

// panicf panics with the format and values passed and assigns the error created to the Reader.
func (r *Reader) panicf(format string, a ...any) {
	panic(fmt.Errorf(format, a...))
}

// panic panics with the error passed, similarly to panicf.
func (r *Reader) panic(err error) {
	panic(err)
}
