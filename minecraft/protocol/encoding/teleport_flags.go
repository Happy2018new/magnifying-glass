package encoding

const TeleportFlagBitsetSize = 32

const (
	TeleportFlagRelativeX = iota
	TeleportFlagRelativeY
	TeleportFlagRelativeZ
	TeleportFlagRelativeYaw
	TeleportFlagRelativePitch
	TeleportFlagRelativeVelocityX
	TeleportFlagRelativeVelocityY
	TeleportFlagRelativeVelocityZ
	// Rotate velocity according to the change in rotation,
	// before applying the velocity change in this packet.
	// Combining this with absolute rotation works as
	// expected—the difference in rotation is still used.
	TeleportFlagRotation
)

// A bit field represented as an int32,
// specifying how a teleportation is
// to be applied on each axis.
//
// In the lower 8 bits of the bit field,
// a set bit means the teleportation on
// the corresponding axis is relative,
// and an unset bit that it is absolute.
type TeleportFlags Bitset
