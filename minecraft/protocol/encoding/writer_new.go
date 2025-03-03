package encoding

import "fmt"

// ConsumeEffectType writes an ConsumeEffect type to the writer.
func (w *Writer) ConsumeEffect(x *ConsumeEffect) {
	var id int32
	if !lookupConsumeEffectType(*x, &id) {
		w.UnknownEnumOption(fmt.Sprintf("%T", x), "consume effect type")
	}
	w.Varint32(&id)
	(*x).Marshal(w)
}
