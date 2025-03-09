package encoding

// LightData maybe is used to do world render,
// to notify how client render the brightness
// of the block.
type LightData struct {
	// BitSet containing bits for each section in the world + 2.
	// Each set bit indicates that the corresponding 16×16×16
	// chunk section has data in the Sky Light array below.
	//
	// The least significant bit is for blocks 16 blocks to 1 block
	// below the min world height (one section below the world),
	// while the most significant bit covers blocks 1 to 16 blocks
	// above the max world height (one section above the world).
	SkyLightMask Bitset
	// 	BitSet containing bits for each section in the world + 2.
	// Each set bit indicates that the corresponding 16×16×16 chunk
	// section has data in the Block Light array below.
	//
	// The order of bits is the same as in Sky Light Mask.
	BlockLightMask Bitset
	// BitSet containing bits for each section in the world + 2.
	// Each set bit indicates that the corresponding 16×16×16 chunk
	// section has all zeros for its Sky Light data.
	//
	// The order of bits is the same as in Sky Light Mask.
	EmptySkyLightMask Bitset
	// BitSet containing bits for each section in the world + 2.
	// Each set bit indicates that the corresponding 16×16×16 chunk
	// section has all zeros for its Block Light data.
	//
	// The order of bits is the same as in Sky Light Mask.
	EmptyBlockLightMask Bitset
	// The length of any inner array is always 2048;
	// There is 1 array for each bit set to true in the sky light mask,
	// starting with the lowest value.
	//
	// Half a byte per light value.
	SkyLightArrays []byte
	// The length of any inner array is always 2048;
	// There is 1 array for each bit set to true in the block light mask,
	// starting with the lowest value.
	//
	// Half a byte per light value.
	BlockLightArrays []byte
}
