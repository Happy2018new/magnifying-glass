package encoding

import "magnifying-glass/minecraft/nbt"

// TextComponentString reads an TextComponentString from the reader.
func (r *Reader) TextComponentString(x *TextComponentString) {
	r.String((*string)(x))
}

// TextComponentComplex reads an TextComponentComplex from the reader.
func (r *Reader) TextComponentComplex(x *TextComponentComplex) {
	r.NBT((*map[string]any)(x), nbt.NetworkBigEndian)
}

// TextComponentComplexOptional reads an TextComponentComplexOptional from the reader.
func (r *Reader) TextComponentComplexOptional(x *TextComponentComplexOptional) {
	r.Bool(&x.Existed)
	if x.Existed {
		r.NBT(&x.Data, nbt.NetworkBigEndian)
	}
}

// JsonTextComponent reads an JsonTextComponent from the reader.
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
	var val int64
	r.Int64(&val)
	x[0] = int32(val >> 38)
	x[1] = int32(val << 52 >> 52)
	x[2] = int32(val << 26 << 38)
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
