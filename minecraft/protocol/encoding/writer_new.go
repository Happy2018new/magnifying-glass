package encoding

import (
	"fmt"
	"magnifying-glass/minecraft/nbt"
	"reflect"
	"sort"

	"github.com/go-gl/mathgl/mgl32"
)

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

// ConsumeEffect writes an ConsumeEffect to the writer.
func (w *Writer) ConsumeEffect(x *ConsumeEffect) {
	var id int32
	if !lookupConsumeEffectType(*x, &id) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "consume effect type")
	}
	w.Varint32(&id)
	(*x).Marshal(w)
}

// Position writes BlockPos as a int64 to the writer.
func (w *Writer) Position(x *BlockPos) {
	partX := (uint64(x[0]) & 0x3FFFFF) << 38
	partZ := (uint64(x[2]) & 0x3FFFFF) << 12
	partY := uint64(x[1]) & 0xFFF
	val := partX | partZ | partY
	w.Uint64(&val)
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

// Particle writes a Particle to the writer.
func (w *Writer) Particle(x *Particle) {
	var t int32
	if !lookupParticleType(*x, &t) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "particle type")
	}
	w.Varint32(&t)
	(*x).Marshal(w)
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

// SoundEvent writes a SoundEvent to the writer.
func (w *Writer) SoundEvent(x *SoundEvent) {
	w.Identifier(&x.SoundName)
	OptionalFunc(w, &x.FixedRange, w.Float32)
}
