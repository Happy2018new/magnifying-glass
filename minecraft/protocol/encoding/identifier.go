package encoding

// Identifiers are a namespaced location, in the form of minecraft:thing.
// If the namespace is not provided, it defaults to minecraft (i.e. thing is minecraft:thing).
//
// Custom content should always be in its own namespace, not the default one.
// Both the namespace and value can use all lowercase alphanumeric characters
// (a-z and 0-9), dot (.), dash (-), and underscore (_).
//
// In addition, values can use slash (/).
// The naming convention is lower_case_with_underscores.
//
// More information (https://minecraft.net/en-us/article/minecraft-snapshot-17w43a).
//
// For ease of determining whether a namespace or value is valid,
// here are regular expressions for each:
// 		- Namespace: [a-z0-9.-_]
// 		- Value: [a-z0-9.-_/]
type Identifier string

const IDSetTypeTagDefined int32 = 0

// Represents a set of IDs in a certain
// registry (implied by context),
// either directly (enumerated IDs) or
// indirectly (tag name).
type IDSet struct {
	// Value used to determine the data that follows.
	// It can be either:
	// 		0 - Represents a named set of IDs defined by a tag.
	// 		Anything else - Represents an ad-hoc set of IDs enumerated inline.
	Type int32
	// The registry tag defining the ID set.
	// Only present if Type is 0.
	TagName Identifier
	// An array of registry IDs. Only present if Type is not 0.
	// The size of the array is equal to Type - 1.
	IDs []int32
}

func (i *IDSet) Marshal(io IO) {
	io.Varint32(&i.Type)
	if i.Type == IDSetTypeTagDefined {
		io.Identifier(&i.TagName)
	} else {
		FuncSliceOfLen(io, uint32(i.Type)-1, &i.IDs, io.Varint32)
	}
}
