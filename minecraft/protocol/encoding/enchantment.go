package encoding

// Single enchantment of a item.
type Enchantment struct {
	// The ID of the enchantment in
	// the enchantment registry.
	EnchID int32
	// The level of the enchantment.
	EnchLevel int32
}

func (e *Enchantment) Marshal(io IO) {
	io.Varint32(&e.EnchID)
	io.Varint32(&e.EnchLevel)
}
