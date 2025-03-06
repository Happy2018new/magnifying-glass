package encoding

import (
	"reflect"
	"sort"
)

// EntityMetadata refer to the metadata
// of Minecraft Java entity.
type EntityMetadata map[uint8]any

const (
	EntityDataTypeByte int32 = iota
	EntityDataTypeVarint32
	EntityDataTypeVarint64
	EntityDataTypeFloat32
	EntityDataTypeString
	EntityDataTypeTextCompound
	EntityDataTypeOptionalTextCompound
	EntityDataTypeSlot
	EntityDataTypeBoolean
	EntityDataTypeRotations
	EntityDataTypePosition
	EntityDataTypeOptionalPosition
	EntityDataTypeDirection
	EntityDataTypeOptionalUUID
	EntityDataTypeBlockState
	EntityDataTypeOptionalBlockState
	EntityDataTypeTagNBT
	EntityDataTypeParticle
	EntityDataTypeParticles
	EntityDataTypeVillagerData
	EntityDataTypeOptionalVarint32
	EntityDataTypePose
	EntityDataTypeCatVariant
	EntityDataTypeWolfVariant
	EntityDataTypeFrogVariant
	EntityDataTypeOptionalGlobalPosition
	EntityDataTypePaintingVariant
	EntityDataTypeSnifferState
	EntityDataTypeArmadilloState
	EntityDataTypeVec3
	EntityDataTypeQuaternion
)

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
		// case map[string]any:
		// 	entityDataTypeCompoundTag := EntityDataTypeCompoundTag
		// 	w.Varuint32(&entityDataTypeCompoundTag)
		// 	w.NBT(&v, nbt.NetworkLittleEndian)
		// case BlockPos:
		// 	entityDataTypeBlockPos := EntityDataTypeBlockPos
		// 	w.Varuint32(&entityDataTypeBlockPos)
		// 	w.BlockPos(&v)
		// case int64:
		// 	entityDataTypeInt64 := EntityDataTypeInt64
		// 	w.Varuint32(&entityDataTypeInt64)
		// 	w.Varint64(&v)
		// case mgl32.Vec3:
		// 	entityDataTypeVec3 := EntityDataTypeVec3
		// 	w.Varuint32(&entityDataTypeVec3)
		// 	w.Vec3(&v)
		default:
			w.UnknownEnumOption(reflect.TypeOf(value), "entity metadata")
		}
	}
}
