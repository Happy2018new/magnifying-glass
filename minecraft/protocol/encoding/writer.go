package encoding

import (
	"fmt"
	"image/color"
	"io"
	"magnifying-glass/minecraft/nbt"
	"magnifying-glass/minecraft/protocol/encoding/basic_encoding"
	"math"
	"slices"
	"unsafe"

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

// ByteSlice writes a []byte, prefixed with a varuint32, to the underlying buffer.
func (w *Writer) ByteSlice(x *[]byte) {
	l := uint32(len(*x))
	w.Varuint32(&l)
	_, _ = w.Writer().Write(*x)
}

// Bytes appends a []byte to the underlying buffer.
func (w *Writer) Bytes(x *[]byte) {
	_, _ = w.Writer().Write(*x)
}

// Angle writes a rotational float32 as a single byte to the underlying buffer.
func (w *Writer) Angle(x *float32) {
	_ = w.Writer().WriteByte(byte(*x / (360.0 / 256.0)))
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

// Vec2 writes an mgl32.Vec2 as 2 float32s to the underlying buffer.
func (w *Writer) Vec2(x *mgl32.Vec2) {
	w.Float32(&x[0])
	w.Float32(&x[1])
}

// // BlockPos writes a BlockPos as 3 varint32s to the underlying buffer.
// func (w *Writer) BlockPos(x *BlockPos) {
// 	w.Varint32(&x[0])
// 	w.Varint32(&x[1])
// 	w.Varint32(&x[2])
// }

// // UBlockPos writes a BlockPos as 2 varint32s and a varuint32 to the underlying buffer.
// func (w *Writer) UBlockPos(x *BlockPos) {
// 	w.Varint32(&x[0])
// 	y := uint32(x[1])
// 	w.Varuint32(&y)
// 	w.Varint32(&x[2])
// }

// // ChunkPos writes a ChunkPos as 2 varint32s to the underlying buffer.
// func (w *Writer) ChunkPos(x *ChunkPos) {
// 	w.Varint32(&x[0])
// 	w.Varint32(&x[1])
// }

// // SubChunkPos writes a SubChunkPos as 3 varint32s to the underlying buffer.
// func (w *Writer) SubChunkPos(x *SubChunkPos) {
// 	w.Varint32(&x[0])
// 	w.Varint32(&x[1])
// 	w.Varint32(&x[2])
// }

// // SoundPos writes an mgl32.Vec3 that serves as a position for a sound.
// func (w *Writer) SoundPos(x *mgl32.Vec3) {
// 	b := BlockPos{int32((*x)[0] * 8), int32((*x)[1] * 8), int32((*x)[2] * 8)}
// 	w.UBlockPos(&b)
// }

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

// VarRGBA writes a color.RGBA x as a varuint32 to the underlying buffer.
func (w *Writer) VarRGBA(x *color.RGBA) {
	val := uint32(x.R) | uint32(x.G)<<8 | uint32(x.B)<<16 | uint32(x.A)<<24
	w.Varuint32(&val)
}

// UUID writes a UUID to the underlying buffer.
func (w *Writer) UUID(x *uuid.UUID) {
	b := [16]byte(*x)
	_, _ = w.Writer().Write(b[:])
}

// // PlayerInventoryAction writes a PlayerInventoryAction.
// func (w *Writer) PlayerInventoryAction(x *UseItemTransactionData) {
// 	w.Varint32(&x.LegacyRequestID)
// 	if x.LegacyRequestID < -1 && (x.LegacyRequestID&1) == 0 {
// 		Slice(w, &x.LegacySetItemSlots)
// 	}
// 	Slice(w, &x.Actions)
// 	w.Varuint32(&x.ActionType)
// 	w.Varuint32(&x.TriggerType)
// 	w.BlockPos(&x.BlockPosition)
// 	w.Varint32(&x.BlockFace)
// 	w.Varint32(&x.HotBarSlot)
// 	w.ItemInstance(&x.HeldItem)
// 	w.Vec3(&x.Position)
// 	w.Vec3(&x.ClickedPosition)
// 	w.Varuint32(&x.BlockRuntimeID)
// 	w.Varuint32(&x.ClientPrediction)
// }

// // GameRule writes a GameRule x to the Writer.
// func (w *Writer) GameRule(x *GameRule) {
// 	w.String(&x.Name)
// 	w.Bool(&x.CanBeModifiedByPlayer)

// 	switch v := x.Value.(type) {
// 	case bool:
// 		id := uint32(1)
// 		w.Varuint32(&id)
// 		w.Bool(&v)
// 	case uint32:
// 		id := uint32(2)
// 		w.Varuint32(&id)
// 		w.Varuint32(&v)
// 	case float32:
// 		id := uint32(3)
// 		w.Varuint32(&id)
// 		w.Float32(&v)
// 	default:
// 		w.UnknownEnumOption(fmt.Sprintf("%T", v), "game rule type")
// 	}
// }

// // EntityMetadata writes an entity metadata map x to the underlying buffer.
// func (w *Writer) EntityMetadata(x *map[uint32]any) {
// 	l := uint32(len(*x))
// 	w.Varuint32(&l)

// 	// Entity metadata needs to be sorted for some functionality to work. NPCs, for example, need to have their fields
// 	// set in increasing order, or the text or buttons won't be shown to the client. See #88.
// 	// Sorting this is probably not very fast, but it'll have to do for now: We can change entity metadata to a slice
// 	// later on.
// 	keys := make([]int, 0, l)
// 	for k := range *x {
// 		keys = append(keys, int(k))
// 	}
// 	sort.Ints(keys)
// 	for _, k := range keys {
// 		key := uint32(k)
// 		value := (*x)[uint32(k)]
// 		w.Varuint32(&key)
// 		switch v := value.(type) {
// 		case byte:
// 			entityDataTypeByte := EntityDataTypeByte
// 			w.Varuint32(&entityDataTypeByte)
// 			w.Uint8(&v)
// 		case int16:
// 			entityDataTypeInt16 := EntityDataTypeInt16
// 			w.Varuint32(&entityDataTypeInt16)
// 			w.Int16(&v)
// 		case int32:
// 			entityDataTypeInt32 := EntityDataTypeInt32
// 			w.Varuint32(&entityDataTypeInt32)
// 			w.Varint32(&v)
// 		case float32:
// 			entityDataTypeFloat32 := EntityDataTypeFloat32
// 			w.Varuint32(&entityDataTypeFloat32)
// 			w.Float32(&v)
// 		case string:
// 			entityDataTypeString := EntityDataTypeString
// 			w.Varuint32(&entityDataTypeString)
// 			w.String(&v)
// 		case map[string]any:
// 			entityDataTypeCompoundTag := EntityDataTypeCompoundTag
// 			w.Varuint32(&entityDataTypeCompoundTag)
// 			w.NBT(&v, nbt.NetworkLittleEndian)
// 		case BlockPos:
// 			entityDataTypeBlockPos := EntityDataTypeBlockPos
// 			w.Varuint32(&entityDataTypeBlockPos)
// 			w.BlockPos(&v)
// 		case int64:
// 			entityDataTypeInt64 := EntityDataTypeInt64
// 			w.Varuint32(&entityDataTypeInt64)
// 			w.Varint64(&v)
// 		case mgl32.Vec3:
// 			entityDataTypeVec3 := EntityDataTypeVec3
// 			w.Varuint32(&entityDataTypeVec3)
// 			w.Vec3(&v)
// 		default:
// 			w.UnknownEnumOption(reflect.TypeOf(value), "entity metadata")
// 		}
// 	}
// }

// // ItemDescriptorCount writes an ItemDescriptorCount i to the underlying buffer.
// func (w *Writer) ItemDescriptorCount(i *ItemDescriptorCount) {
// 	var id byte
// 	switch i.Descriptor.(type) {
// 	case *InvalidItemDescriptor:
// 		id = ItemDescriptorInvalid
// 	case *DefaultItemDescriptor:
// 		id = ItemDescriptorDefault
// 	case *MoLangItemDescriptor:
// 		id = ItemDescriptorMoLang
// 	case *ItemTagItemDescriptor:
// 		id = ItemDescriptorItemTag
// 	case *DeferredItemDescriptor:
// 		id = ItemDescriptorDeferred
// 	case *ComplexAliasItemDescriptor:
// 		id = ItemDescriptorComplexAlias
// 	default:
// 		w.UnknownEnumOption(fmt.Sprintf("%T", i.Descriptor), "item descriptor type")
// 		return
// 	}
// 	w.Uint8(&id)

// 	i.Descriptor.Marshal(w)
// 	w.Varint32(&i.Count)
// }

// // ItemInstance writes an ItemInstance i to the underlying buffer.
// func (w *Writer) ItemInstance(i *ItemInstance) {
// 	x := &i.Stack
// 	w.Varint32(&x.NetworkID)
// 	if x.NetworkID == 0 {
// 		// The item was air, so there's no more data to follow. Return immediately.
// 		return
// 	}

// 	w.Uint16(&x.Count)
// 	w.Varuint32(&x.MetadataValue)

// 	hasNetID := i.StackNetworkID != 0
// 	w.Bool(&hasNetID)

// 	if hasNetID {
// 		w.Varint32(&i.StackNetworkID)
// 	}

// 	w.Varint32(&x.BlockRuntimeID)

// 	buf := new(bytes.Buffer)
// 	bufWriter := NewWriter(buf, w.shieldID)

// 	var length int16
// 	if len(x.NBTData) != 0 {
// 		length = int16(-1)
// 		version := uint8(1)

// 		bufWriter.Int16(&length)
// 		bufWriter.Uint8(&version)
// 		bufWriter.NBT(&x.NBTData, nbt.LittleEndian)
// 	} else {
// 		bufWriter.Int16(&length)
// 	}

// 	FuncSliceUint32Length(bufWriter, &x.CanBePlacedOn, bufWriter.StringUTF)
// 	FuncSliceUint32Length(bufWriter, &x.CanBreak, bufWriter.StringUTF)

// 	if x.NetworkID == bufWriter.shieldID {
// 		var blockingTick int64
// 		bufWriter.Int64(&blockingTick)
// 	}

// 	b := buf.Bytes()
// 	w.ByteSlice(&b)
// }

// // Item writes an ItemStack x to the underlying buffer.
// func (w *Writer) Item(x *ItemStack) {
// 	w.Varint32(&x.NetworkID)
// 	if x.NetworkID == 0 {
// 		// The item was air, so there's no more data to follow. Return immediately.
// 		return
// 	}

// 	w.Uint16(&x.Count)
// 	w.Varuint32(&x.MetadataValue)
// 	w.Varint32(&x.BlockRuntimeID)

// 	var extraData []byte
// 	buf := bytes.NewBuffer(extraData)
// 	bufWriter := NewWriter(buf, w.shieldID)

// 	var length int16
// 	if len(x.NBTData) != 0 {
// 		length = int16(-1)
// 		version := uint8(1)

// 		bufWriter.Int16(&length)
// 		bufWriter.Uint8(&version)
// 		bufWriter.NBT(&x.NBTData, nbt.LittleEndian)
// 	} else {
// 		bufWriter.Int16(&length)
// 	}

// 	FuncSliceUint32Length(bufWriter, &x.CanBePlacedOn, bufWriter.StringUTF)
// 	FuncSliceUint32Length(bufWriter, &x.CanBreak, bufWriter.StringUTF)

// 	if x.NetworkID == bufWriter.shieldID {
// 		var blockingTick int64
// 		bufWriter.Int64(&blockingTick)
// 	}

// 	extraData = buf.Bytes()
// 	w.ByteSlice(&extraData)
// }

// // StackRequestAction writes a StackRequestAction to the writer.
// func (w *Writer) StackRequestAction(x *StackRequestAction) {
// 	var id byte
// 	if !lookupStackRequestActionType(*x, &id) {
// 		w.UnknownEnumOption(fmt.Sprintf("%T", *x), "stack request action type")
// 	}
// 	w.Uint8(&id)
// 	(*x).Marshal(w)
// }

// // MaterialReducer writes a material reducer to the writer.
// func (w *Writer) MaterialReducer(m *MaterialReducer) {
// 	mix := (m.InputItem.NetworkID << 16) | int32(m.InputItem.MetadataValue)
// 	w.Varint32(&mix)
// 	Slice(w, &m.Outputs)
// }

// // Recipe writes a Recipe to the writer.
// func (w *Writer) Recipe(x *Recipe) {
// 	var recipeType int32
// 	if !lookupRecipeType(*x, &recipeType) {
// 		w.UnknownEnumOption(fmt.Sprintf("%T", *x), "crafting recipe type")
// 	}
// 	w.Varint32(&recipeType)
// 	(*x).Marshal(w)
// }

// // EventType writes an Event to the writer.
// func (w *Writer) EventType(x *Event) {
// 	var t int32
// 	if !lookupEventType(*x, &t) {
// 		w.UnknownEnumOption(fmt.Sprintf("%T", x), "event packet event type")
// 	}
// 	w.Varint32(&t)
// }

// // TransactionDataType writes an InventoryTransactionData type to the writer.
// func (w *Writer) TransactionDataType(x *InventoryTransactionData) {
// 	var id uint32
// 	if !lookupTransactionDataType(*x, &id) {
// 		w.UnknownEnumOption(fmt.Sprintf("%T", x), "inventory transaction data type")
// 	}
// 	w.Varuint32(&id)
// }

// // AbilityValue writes an ability value to the writer.
// func (w *Writer) AbilityValue(x *any) {
// 	switch val := (*x).(type) {
// 	case bool:
// 		valType, defaultVal := uint8(1), float32(0)
// 		w.Uint8(&valType)
// 		w.Bool(&val)
// 		w.Float32(&defaultVal)
// 	case float32:
// 		valType, defaultVal := uint8(2), false
// 		w.Uint8(&valType)
// 		w.Bool(&defaultVal)
// 		w.Float32(&val)
// 	default:
// 		w.InvalidValue(*x, "ability value type", "must be bool or float32")
// 	}
// }

// // CompressedBiomeDefinitions reads a list of compressed biome definitions from the reader. Minecraft decided to make their
// // own type of compression for this, so we have to implement it ourselves. It uses a dictionary of repeated byte sequences
// // to reduce the size of the data. The compressed data is read byte-by-byte, and if the byte is 0xff then it is assumed
// // that the next two bytes are an int16 for the dictionary index. Otherwise, the byte is copied to the output. The dictionary
// // index is then used to look up the byte sequence to be appended to the output.
// func (w *Writer) CompressedBiomeDefinitions(x *map[string]any) {
// 	decompressed, err := nbt.Marshal(x)
// 	if err != nil {
// 		w.panicf("error marshaling nbt: %v", err)
// 	}

// 	var compressed []byte
// 	buf := bytes.NewBuffer(compressed)
// 	bufWriter := NewWriter(buf, w.shieldID)

// 	header := []byte("COMPRESSED")
// 	bufWriter.Bytes(&header)

// 	// TODO: Dictionary compression implementation
// 	var dictionaryLength uint16
// 	bufWriter.Uint16(&dictionaryLength)
// 	for _, b := range decompressed {
// 		bufWriter.Uint8(&b)
// 		if b == 0xff {
// 			dictionaryIndex := int16(1)
// 			bufWriter.Int16(&dictionaryIndex)
// 		}
// 	}

// 	compressed = buf.Bytes()
// 	length := uint32(len(compressed))
// 	w.Varuint32(&length)
// 	w.Bytes(&compressed)
// }

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

// TeleportFlags writes a TeleportFlags to the writer.
func (w *Writer) TeleportFlags(x *TeleportFlags) {
	w.FixedBitset((*Bitset)(x), TeleportFlagBitsetSize)
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

// // ShieldID returns the shield ID provided to the writer.
// func (w *Writer) ShieldID() int32 {
// 	return w.shieldID
// }

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
