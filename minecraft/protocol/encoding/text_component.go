package encoding

// TextComponentString means this is a text
// component but encoding with TAG_String.
type TextComponentString string

// TextComponentComplex means this is a text
// component but encoding with TAG_Compound.
type TextComponentComplex map[string]any

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

// JsonTextComponent means this is a JSON text
// component but encoding with TAG_String.
type JsonTextComponent string
