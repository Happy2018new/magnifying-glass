package encoding

import (
	"fmt"
	"magnifying-glass/minecraft/nbt"
)

// TextComponentString writes an TextComponentString to the writer.
func (w *Writer) TextComponentString(x *TextComponentString) {
	w.String((*string)(x))
}

// TextComponentComplex writes an TextComponentComplex to the writer.
func (w *Writer) TextComponentComplex(x *TextComponentComplex) {
	w.NBT((*map[string]any)(x), nbt.NetworkBigEndian)
}

// TextComponentComplexOptional writes an TextComponentComplexOptional to the writer.
func (w *Writer) TextComponentComplexOptional(x *TextComponentComplexOptional) {
	w.Bool(&x.Existed)
	if x.Existed {
		w.NBT(&x.Data, nbt.NetworkBigEndian)
	}
}

// JsonTextComponent writes an JsonTextComponent to the writer.
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
	partA := (int64(x[0]) & 0x3FFFFF) << 38
	partB := (int64(x[1]) & 0x3FFFFF) << 12
	partC := int64(x[2]) & 0xFFF
	val := partA | partB | partC
	w.Int64(&val)
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
