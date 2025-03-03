package encoding

// ConsumeEffect reads an ConsumeEffect from the reader.
func (r *Reader) ConsumeEffect(x *ConsumeEffect) {
	var t int32
	r.Varint32(&t)
	if !lookupConsumeEffect(t, x) {
		r.UnknownEnumOption(t, "consume effect type")
	}
	(*x).Marshal(r)
}
