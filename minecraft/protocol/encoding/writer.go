package encoding

import (
	"fmt"
	"image/color"
	"io"
	"math"
	"reflect"
	"slices"
	"sort"
	"unsafe"

	"github.com/Happy2018new/magnifying-glass/minecraft/nbt"
	"github.com/Happy2018new/magnifying-glass/minecraft/protocol/encoding/basic_encoding"

	"github.com/go-gl/mathgl/mgl32"
	"github.com/google/uuid"
)

// Writer implements writing methods for data types
// from Minecraft packets.
//
// Each Packet implementation has one passed to it
// when writing.
//
// Writer implements methods where values are passed
// using a pointer, so that Reader and Writer have a
// synonymous interface and both implement the IO
// interface.
type Writer struct {
	*basic_encoding.BasicWriter
}

// NewWriter creates a new initialised Writer with an
// underlying io.ByteWriter to write to.
func NewWriter(w interface {
	io.Writer
	io.ByteWriter
}, shieldID int32) *Writer {
	return &Writer{basic_encoding.NewBasicWriter(w)}
}

// Bool writes a bool as either 0 or 1 to the underlying buffer.
func (w *Writer) Bool(x *bool) {
	_ = w.Writer().WriteByte(*(*byte)(unsafe.Pointer(x)))
}

// String writes a string, prefixed with a varint16, to the underlying buffer.
func (w *Writer) String(x *string) {
	l := int16(len(*x))
	w.Varint16(&l)
	_, _ = w.Writer().Write([]byte(*x))
}

// TextComponentString writes a TextComponentString to the writer.
func (w *Writer) TextComponentString(x *TextComponentString) {
	w.String((*string)(x))
}

// TextComponentComplex writes a TextComponentComplex to the writer.
func (w *Writer) TextComponentComplex(x *TextComponentComplex) {
	w.NBT((*map[string]any)(x), nbt.NetworkBigEndian)
}

// TextComponentComplexOptional writes a TextComponentComplexOptional to the writer.
func (w *Writer) TextComponentComplexOptional(x *TextComponentComplexOptional) {
	w.Bool(&x.Existed)
	if x.Existed {
		w.NBT(&x.Data, nbt.NetworkBigEndian)
	}
}

// JsonTextComponent writes a JsonTextComponent to the writer.
func (w *Writer) JsonTextComponent(x *JsonTextComponent) {
	w.NBTString((*string)(x), nbt.NetworkBigEndian)
}

// Identifier writes an Identifier to the writer.
func (w *Writer) Identifier(x *Identifier) {
	w.String((*string)(x))
}

// WriteEntityMetadata writes an entity metadata
// map x to the underlying buffer.
func (w *Writer) EntityMetadata(x *EntityMetadata) {
	l := uint32(len(*x))

	// I don't know if it is necessary to do the key
	// sorting on Java edition of Minecraft.
	// Howerver, let me do the same things to make
	// everything is sorted.
	keys := make([]int, 0, l)
	for k := range *x {
		keys = append(keys, int(k))
	}
	sort.Ints(keys)

	for _, k := range keys {
		index := uint8(k)
		value := (*x)[index]
		w.Uint8(&index)
		switch v := value.(type) {
		case byte:
			entityDataTypeByte := EntityDataTypeByte
			w.Varint32(&entityDataTypeByte)
			w.Uint8(&v)
		case int32:
			entityDataTypeVarint32 := EntityDataTypeVarint32
			w.Varint32(&entityDataTypeVarint32)
			w.Varint32(&v)
		case int64:
			entityDataTypeVarint64 := EntityDataTypeVarint64
			w.Varint32(&entityDataTypeVarint64)
			w.Varint64(&v)
		case float32:
			entityDataTypeFloat32 := EntityDataTypeFloat32
			w.Varint32(&entityDataTypeFloat32)
			w.Float32(&v)
		case string:
			entityDataTypeString := EntityDataTypeString
			w.Varint32(&entityDataTypeString)
			w.String(&v)
		case TextComponentComplex:
			entityDataTypeTextCompound := EntityDataTypeTextCompound
			w.Varint32(&entityDataTypeTextCompound)
			w.TextComponentComplex(&v)
		case TextComponentComplexOptional:
			entityDataTypeOptionalTextCompound := EntityDataTypeOptionalTextCompound
			w.Varint32(&entityDataTypeOptionalTextCompound)
			w.TextComponentComplexOptional(&v)
		case ItemStack:
			entityDataTypeItemStack := EntityDataTypeItemStack
			w.Varint32(&entityDataTypeItemStack)
			w.ItemStack(&v)
		case bool:
			entityDataTypeBoolean := EntityDataTypeBoolean
			w.Varint32(&entityDataTypeBoolean)
			w.Bool(&v)
		case EntityDataRotations:
			entityDataTypeRotations := EntityDataTypeRotations
			w.Varint32(&entityDataTypeRotations)
			v.Marshal(w)
		case BlockPos:
			entityDataTypePosition := EntityDataTypePosition
			w.Varint32(&entityDataTypePosition)
			w.Position(&v)
		case EntityDataOptionalPosition:
			entityDataTypeOptionalPosition := EntityDataTypeOptionalPosition
			w.Varint32(&entityDataTypeOptionalPosition)
			v.Marshal(w)
		case EntityDataDirection:
			entityDataTypeDirection := EntityDataTypeDirection
			w.Varint32(&entityDataTypeDirection)
			v.Marshal(w)
		case EntityDataOptionalUUID:
			entityDataTypeOptionalUUID := EntityDataTypeOptionalUUID
			w.Varint32(&entityDataTypeOptionalUUID)
			v.Marshal(w)
		case EntityDataBlockState:
			entityDataTypeBlockState := EntityDataTypeBlockState
			w.Varint32(&entityDataTypeBlockState)
			v.Marshal(w)
		case EntityDataOptionalBlockState:
			entityDataTypeOptionalBlockState := EntityDataTypeOptionalBlockState
			w.Varint32(&entityDataTypeOptionalBlockState)
			v.Marshal(w)
		case map[string]any:
			entityDataTypeTagNBT := EntityDataTypeTagNBT
			w.Varint32(&entityDataTypeTagNBT)
			w.NBT(&v, nbt.NetworkBigEndian)
		case Particle:
			entityDataTypeParticle := EntityDataTypeParticle
			w.Varint32(&entityDataTypeParticle)
			w.Particle(&v)
		case []Particle:
			entityDataTypeParticles := EntityDataTypeParticles
			w.Varint32(&entityDataTypeParticles)
			FuncSliceVarint32Length(w, &v, w.Particle)
		case EntityDataVillagerData:
			entityDataTypeVillagerData := EntityDataTypeVillagerData
			w.Varint32(&entityDataTypeVillagerData)
			v.Marshal(w)
		case EntityDataOptionalVarint32:
			entityDataTypeOptionalVarint32 := EntityDataTypeOptionalVarint32
			w.Varint32(&entityDataTypeOptionalVarint32)
			v.Marshal(w)
		case EntityDataPose:
			entityDataTypePose := EntityDataTypePose
			w.Varint32(&entityDataTypePose)
			v.Marshal(w)
		case EntityDataCatVariant:
			entityDataTypeCatVariant := EntityDataTypeCatVariant
			w.Varint32(&entityDataTypeCatVariant)
			v.Marshal(w)
		case EntityDataWolfVariant:
			entityDataTypeWolfVariant := EntityDataTypeWolfVariant
			w.Varint32(&entityDataTypeWolfVariant)
			v.Marshal(w)
		case EntityDataForgVariant:
			entityDataTypeFrogVariant := EntityDataTypeFrogVariant
			w.Varint32(&entityDataTypeFrogVariant)
			v.Marshal(w)
		case EntityDataOptionalGlobalPosition:
			entityDataTypeOptionalGlobalPosition := EntityDataTypeOptionalGlobalPosition
			w.Varint32(&entityDataTypeOptionalGlobalPosition)
			v.Marshal(w)
		case EntityDataPaintingVariant:
			entityDataTypePaintingVariant := EntityDataTypePaintingVariant
			w.Varint32(&entityDataTypePaintingVariant)
			v.Marshal(w)
		case EntityDataSnifferState:
			entityDataTypeSnifferState := EntityDataTypeSnifferState
			w.Varint32(&entityDataTypeSnifferState)
			v.Marshal(w)
		case EntityDataArmadilloState:
			entityDataTypeArmadilloState := EntityDataTypeArmadilloState
			w.Varint32(&entityDataTypeArmadilloState)
			v.Marshal(w)
		case mgl32.Vec3:
			entityDataTypeVec3 := EntityDataTypeVec3
			w.Varint32(&entityDataTypeVec3)
			w.Vec3(&v)
		case mgl32.Vec4:
			entityDataTypeQuaternion := EntityDataTypeQuaternion
			w.Varint32(&entityDataTypeQuaternion)
			w.Vec4(&v)
		default:
			w.UnknownEnumOption(reflect.TypeOf(value), "entity metadata")
		}
	}

	entityMetadataTagEnd := EntityMetadataTagEnd
	w.Uint8(&entityMetadataTagEnd)
}

// ItemStack writes an ItemStack to the writer.
func (w *Writer) ItemStack(x *ItemStack) {
	w.Varint32(&x.ItemCount)
	if x.ItemCount == 0 {
		return
	}
	w.Varint32(&x.ItemID)
	w.Varint32(&x.AddComponentsCount)
	w.Varint32(&x.RemoveComponentsCount)
	FuncSliceOfLen(w, uint32(x.AddComponentsCount), &x.ComponentsToAdd, w.ItemComponent)
	FuncSliceOfLen(w, uint32(x.RemoveComponentsCount), &x.ComponentsToRemove, w.Varint32)
}

// Position writes BlockPos as a int64 to the writer.
func (w *Writer) Position(x *BlockPos) {
	partX := (uint64(x[0]) & 0x3FFFFF) << 38
	partZ := (uint64(x[2]) & 0x3FFFFF) << 12
	partY := uint64(x[1]) & 0xFFF
	val := partX | partZ | partY
	w.Uint64(&val)
}

// Angle writes a rotational float32 as a single byte to the underlying buffer.
func (w *Writer) Angle(x *float32) {
	_ = w.Writer().WriteByte(byte(*x / (360.0 / 256.0)))
}

// UUID writes a UUID to the underlying buffer.
func (w *Writer) UUID(x *uuid.UUID) {
	b := [16]byte(*x)
	_, _ = w.Writer().Write(b[:])
}

// Bitset writes Bitset as Java standard bitset that is
// []uint64 with prefixed varint64 as its length into the
// underlying buffer. The encoding is little-endian.
func (w *Writer) Bitset(x *Bitset) {
	bitsSliceOrigin := x.bits.Bits()
	length := len(bitsSliceOrigin)

	bitsSlice := make([]uint64, length)
	for i := range length {
		bitsSlice[i] = uint64(bitsSliceOrigin[i])
	}

	FuncSliceVarint32Length(w, &bitsSlice, w.Uint64)
}

// FixedBitset writes Minecraft Java fixed Bitset
// as []byte without any prefixed things into the
// underlying buffer.
// The size of []byte is giving by size, which meet
// len([]byte) = celi(x.Len() / 8).
// The encoding is little-endian.
func (w *Writer) FixedBitset(x *Bitset, size int) {
	if x.Len() != size {
		w.panicf("bitset size mismatch: expected %v, got %v", size, x.Len())
	}
	fixedSize := uint32(math.Ceil(float64(size) / 8))

	bitsSlice := x.bits.FillBytes(make([]byte, fixedSize))
	slices.Reverse(bitsSlice)

	w.Writer().Write(bitsSlice)
}

// SoundEvent writes a SoundEvent to the writer.
func (w *Writer) SoundEvent(x *SoundEvent) {
	w.Identifier(&x.SoundName)
	OptionalFunc(w, &x.FixedRange, w.Float32)
}

// TeleportFlags writes a TeleportFlags to the writer.
func (w *Writer) TeleportFlags(x *TeleportFlags) {
	w.FixedBitset((*Bitset)(x), TeleportFlagBitsetSize)
}

// RecipeDisplay writes a RecipeDisplay to the writer.
func (w *Writer) RecipeDisplay(x *RecipeDisplay) {
	var t int32
	if !lookupRecipeDisplayType(*x, &t) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "recipe display type")
	}
	w.Varint32(&t)
	(*x).Marshal(w)
}

// SlotDisplay writes a SlotDisplay to the writer.
func (w *Writer) SlotDisplay(x *SlotDisplay) {
	var t int32
	if !lookupSlotDisplayType(*x, &t) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "slot display type")
	}
	w.Varint32(&t)
	(*x).Marshal(w)
}

// ChunkData writes a ChunkData to the writer.
func (w *Writer) ChunkData(x *ChunkData) {
	w.NBT(&x.Heightmaps, nbt.NetworkBigEndian)
	FuncSliceVarint32Length(w, &x.Data, w.Uint8)
	SliceVarint32Length(w, &x.BlockEntities)
}

// LightData writes a LightData to the writer.
func (w *Writer) LightData(x *LightData) {
	w.Bitset(&x.SkyLightMask)
	w.Bitset(&x.BlockLightMask)
	w.Bitset(&x.EmptySkyLightMask)
	w.Bitset(&x.EmptyBlockLightMask)
	SliceVarint32Length(w, &x.SkyLightArrays)
	SliceVarint32Length(w, &x.BlockLightArrays)
}

// NBT writes a map as NBT to the underlying buffer using the encoding passed.
func (w *Writer) NBT(x *map[string]any, encoding nbt.Encoding) {
	if err := nbt.NewEncoderWithEncoding(w.Writer(), encoding).Encode(*x); err != nil {
		panic(err)
	}
}

// NBTList writes a slice as NBT to the underlying buffer using the encoding passed.
func (w *Writer) NBTList(x *[]any, encoding nbt.Encoding) {
	if err := nbt.NewEncoderWithEncoding(w.Writer(), encoding).Encode(*x); err != nil {
		panic(err)
	}
}

// NBTString writes a string as NBT to the underlying buffer using the encoding passed.
func (w *Writer) NBTString(x *string, encoding nbt.Encoding) {
	if err := nbt.NewEncoderWithEncoding(w.Writer(), encoding).Encode(*x); err != nil {
		panic(err)
	}
}

// ConsumeEffect writes an ConsumeEffect to the writer.
func (w *Writer) ConsumeEffect(x *ConsumeEffect) {
	var id int32
	if !lookupConsumeEffectType(*x, &id) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "consume effect type")
	}
	w.Varint32(&id)
	(*x).Marshal(w)
}

// ItemComponent writes an ItemComponent to the writer.
func (w *Writer) ItemComponent(x *ItemComponent) {
	var t int32
	if !lookupItemComponentType(*x, &t) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "item component type")
	}
	w.Varint32(&t)
	(*x).Marshal(w)
}

// Particle writes a Particle to the writer.
func (w *Writer) Particle(x *Particle) {
	var t int32
	if !lookupParticleType(*x, &t) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "particle type")
	}
	w.Varint32(&t)
	(*x).Marshal(w)
}

// BossBarActionType writes the type of BossBarAction to the writer.
func (w *Writer) BossBarActionType(x *BossBarAction) {
	var t int32
	if !lookupBossBarActionType(*x, &t) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "boss bar action type")
	}
	w.Varint32(&t)
}

// CommandParserType writes the type of CommandParserType to the writer.
func (w *Writer) CommandParserType(x *CommandParser) {
	var t int32
	if !lookupCommandParserType(*x, &t) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "command parser type")
	}
	w.Varint32(&t)
}

// Vec4 writes an mgl32.Vec4 as 4 float32s to the underlying buffer.
func (w *Writer) Vec4(x *mgl32.Vec4) {
	w.Float32(&x[0])
	w.Float32(&x[1])
	w.Float32(&x[2])
	w.Float32(&x[3])
}

// Vec3 writes an mgl32.Vec3 as 3 float32s to the underlying buffer.
func (w *Writer) Vec3(x *mgl32.Vec3) {
	w.Float32(&x[0])
	w.Float32(&x[1])
	w.Float32(&x[2])
}

// RGB writes a color.RGBA x as a uint32 0xRRGGBB the underlying buffer.
func (w *Writer) RGB(x *color.RGBA) {
	val := uint32(x.R)<<16 | uint32(x.G)<<8 | uint32(x.B)
	w.Uint32(&val)
}

// RGBA writes a color.RGBA x as a uint32 0xAARRGGBB to the underlying buffer.
func (w *Writer) RGBA(x *color.RGBA) {
	val := uint32(x.A)<<24 | uint32(x.R)<<16 | uint32(x.G)<<8 | uint32(x.B)
	w.Uint32(&val)
}

// Bytes appends a []byte to the underlying buffer.
func (w *Writer) Bytes(x *[]byte) {
	_, _ = w.Writer().Write(*x)
}

// UnknownEnumOption panics with an unknown enum option error.
func (w *Writer) UnknownEnumOption(value any, enum string) {
	w.panicf("unknown value '%v' for enum type '%v'", value, enum)
}

// InvalidValue panics with an invalid value error.
func (w *Writer) InvalidValue(value any, forField, reason string) {
	w.panicf("invalid value '%v' for %v: %v", value, forField, reason)
}

// panicf panics with the format and values passed.
func (w *Writer) panicf(format string, a ...any) {
	panic(fmt.Errorf(format, a...))
}
