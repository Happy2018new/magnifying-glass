package encoding

// MinecraftEnum package a Minecraft Java
// enumerate by type T.
type MinecraftEnum[T comparable] struct {
	iotaNumber int32
	constSlice []T
	mapping    map[T]int32
}

// NewMinecraftEnum create a new Minecraft Java
// enumerate by content that data type is []T.
//
// iotaNumber means the iota of those
// enumerates.
//
// For example, if content = []string{"a","b","c"}
// and iotaNumber = 5,
// then it means:
// 		- The id of "a" is 5
// 		- The id of "b" is 6
// 		- The id of "c" is 7
// Therefore:
// 		- Value(6) will get "b"
// 		- Key("c") will get 7
func NewMinecraftEnum[T comparable](iotaNumber int32, content []T) *MinecraftEnum[T] {
	mapping := make(map[T]int32)
	for key, value := range content {
		mapping[value] = int32(key) + iotaNumber
	}
	return &MinecraftEnum[T]{
		constSlice: content,
		mapping:    mapping,
		iotaNumber: iotaNumber,
	}
}

// Value returns the enumeration whose id is key.
func (m *MinecraftEnum[T]) Value(key int32) (result T, exist bool) {
	realKey := key - m.iotaNumber
	if int(realKey) >= len(m.constSlice) {
		return result, false
	}
	return m.constSlice[realKey], true
}

// Value returns the enumeration id of value.
// If not exist, return -1.
func (m *MinecraftEnum[T]) Key(value T) int32 {
	result, exist := m.mapping[value]
	if !exist {
		return -1
	}
	return result
}
