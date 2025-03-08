package encoding

import "math/big"

// Bitset is a representation of std::bitset<size> being
// sent over the network, allowing for more than 64 bits
// to be stored in a single integer.
//
// A Bitset has a fixed size, which is set at creation time.
type Bitset struct {
	// size is the amount of bits
	// that the Bitset can store.
	size int
	// bits is the underlying
	// record of Bitset.
	bits *big.Int
}

// NewBitset creates a new Bitset with a specific size.
// The size is the amount of bits that the Bitset can store.
// Attempting to set a bit at an index higher than the size will panic.
func NewBitset(size int) Bitset {
	return Bitset{size: size, bits: new(big.Int)}
}

// Set sets a bit at a specific index in the Bitset.
// If the index is higher than the size of the Bitset, a
// panic will occur.
func (b Bitset) Set(i int) {
	if i >= b.size {
		panic("index out of bounds")
	}
	b.bits.SetBit(b.bits, i, 1)
}

// Unset unsets a bit at a specific index in the Bitset.
// If the index is higher than the size of the Bitset,
// a panic will occur.
func (b Bitset) Unset(i int) {
	if i >= b.size {
		panic("index out of bounds")
	}
	b.bits.SetBit(b.bits, i, 0)
}

// Load returns if a bit at a specific index in the Bitset is set.
// If the index is higher than the size of the Bitset, a panic will occur.
func (b Bitset) Load(i int) bool {
	if i >= b.size {
		panic("index out of bounds")
	}
	return b.bits.Bit(i) == 1
}

// Len returns the size of the Bitset.
func (b Bitset) Len() int {
	return b.size
}
