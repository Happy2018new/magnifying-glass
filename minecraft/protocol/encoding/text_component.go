package encoding

import (
	"magnifying-glass/minecraft/nbt"
)

// TextComponentString means this is a text
// component but encoding with TAG_String.
type TextComponentString string

func (t *TextComponentString) Marshal(io IO) {
	str := string(*t)
	io.NBTString(&str, nbt.NetworkBigEndian)
	*t = TextComponentString(str)
}

// TextComponentComplex means this is a text
// component but encoding with TAG_Compound.
type TextComponentComplex map[string]any

func (t *TextComponentComplex) Marshal(io IO) {
	mapping := map[string]any(*t)
	io.NBT(&mapping, nbt.NetworkBigEndian)
	*t = TextComponentComplex(mapping)
}

// TextComponentComplexOptional means this is a optional
// text component but encoding with TAG_Compound.
type TextComponentComplexOptional struct {
	// Existed is true means this field should
	// be read/write to the underlying buffer.
	Existed bool
	// Data refer to the real data of
	// TextComponentComplexOptional.
	Data map[string]any
}

func (t *TextComponentComplexOptional) Marshal(io IO) {
	io.Bool(&t.Existed)
	if t.Existed {
		io.NBT(&t.Data, nbt.NetworkBigEndian)
	}
}

// JsonTextComponent means this is a JSON text
// component but encoding with TAG_String.
type JsonTextComponent string

func (t *JsonTextComponent) Marshal(io IO) {
	str := string(*t)
	io.NBTString(&str, nbt.NetworkBigEndian)
	*t = JsonTextComponent(str)
}
