package encoding

const (
	TextComponentDataTypeTagString = iota
	TextComponentDataTypeTagNBT
)

// Encoded as a NBT Tag, with the type of tag used depending on the case:
//   - As a String Tag: For components only containing text (no styling, no events etc.).
//   - As a Compound Tag: Every other case.
type TextComponent struct {
	// dataType represents the data type of this text component.
	// if it encoded as a string tag, then this field will be
	// TextComponentDataTypeTagString (0), otherwise will be 0
	// (TextComponentDataTypeTagNBT)
	dataType int
	// payload represents the real data of this TextComponent carried on.
	payload any
}

// IsRawString return true only this text
// component is just a plaintext component.
//
// Note that "IsRawString() != IsComplex()"
// is always true.
func (t TextComponent) IsRawString() bool {
	return (t.dataType == TextComponentDataTypeTagString)
}

// IsComplex return true only this text
// component is encoded as a Compound Tag.
//
// Note that "IsComplex() != IsRawString()"
// is always true.
func (t TextComponent) IsComplex() bool {
	return (t.dataType == TextComponentDataTypeTagNBT)
}

// LoadAsRawString load this text component as a raw string.
// Only success when this text component is a plaintext component.
func (t TextComponent) LoadAsRawString() (result string, success bool) {
	if t.dataType != TextComponentDataTypeTagString {
		return "", false
	}
	return t.payload.(string), true
}

// LoadAsComplex load this text component as a map.
// Only success when this text component is encoded as Compound Tag.
func (t TextComponent) LoadAsComplex() (result map[string]any, success bool) {
	if t.dataType != TextComponentDataTypeTagNBT {
		return nil, false
	}
	return t.payload.(map[string]any), true
}

// TextComponentOptional means this is a optional
// text component that with a prefixed bool.
type TextComponentOptional Optional[TextComponent]

func (t *TextComponentOptional) Marshal(io IO) {
	OptionalFunc(io, (*Optional[TextComponent])(t), io.TextComponent)
}

// JsonTextComponent means this is a JSON text
// component but encoding with TAG_String.
type JsonTextComponent string
