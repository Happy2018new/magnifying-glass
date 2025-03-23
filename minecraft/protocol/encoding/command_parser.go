package encoding

// ConstCommandParserEnum record all the Minecraft Java
// command parser (each command parer name and its id).
// Dump from (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Command_Data#Parsers).
var ConstCommandParserEnum = NewMinecraftEnum(
	0,
	[]string{
		// 0;
		// Boolean value (true or false, case-sensitive).
		"brigadier:bool",
		"brigadier:float",   // 1; See corresponding struct for details.
		"brigadier:double",  // 2; See corresponding struct for details.
		"brigadier:integer", // 3; See corresponding struct for details.
		"brigadier:long",    // 4; See corresponding struct for details.
		"brigadier:string",  // 5; See corresponding struct for details.
		"minecraft:entity",  // 6; See corresponding struct for details.
		// 7;
		// A player, online or not.
		// Can also use a selector,
		// which may match one or
		// more players (but not entities).
		"minecraft:game_profile",
		// 8;
		// A location, represented as 3 numbers
		// (which must be integers).
		// May use relative locations with ~.
		"minecraft:block_pos",
		// 9;
		// A column location, represented as 2 numbers
		// (which must be integers).
		// May use relative locations with ~.
		"minecraft:column_pos",
		// 10;
		// A location, represented as 3 numbers
		// (which may have a decimal point,
		// but will be moved to the center of
		// a block if none is specified).
		// May use relative locations with ~.
		"minecraft:vec3",
		// 11;
		// A location, represented as 2 numbers
		// (which may have a decimal point, but will
		// be moved to the center of a block if none
		// is specified).
		// May use relative locations with ~.
		"minecraft:vec2",
		// 12;
		// A block state, optionally including
		// NBT and state information.
		"minecraft:block_state",
		// 13;
		// A block, or a block tag.
		"minecraft:block_predicate",
		// 14;
		// An item, optionally including NBT.
		"minecraft:item_stack",
		// 15;
		// An item, or an item tag.
		"minecraft:item_predicate",
		// 16;
		// A chat color that Case-insensitive.
		// One of the names from
		// Text formatting#Colors (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Text_formatting#Colors),
		// or reset.
		"minecraft:color",
		// 17;
		// A JSON text component.
		"minecraft:component",
		// 18;
		// A JSON object containing the
		// text component styling fields (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Text_formatting#Styling_fields).
		"minecraft:style",
		// 19;
		// A regular message, potentially including selectors.
		"minecraft:message",
		// 20;
		// An NBT value, parsed using JSON-NBT rules.
		"minecraft:nbt",
		// 21;
		// Represents a partial nbt tag,
		// usable in data modify command.
		"minecraft:nbt_tag",
		// 22;
		// A path within an NBT value,
		// allowing for array and member accesses.
		"minecraft:nbt_path",
		// 23;
		// A scoreboard objective.
		"minecraft:objective",
		// 24;
		// A single score criterion.
		"minecraft:objective_criteria",
		// 25;
		// A scoreboard operator.
		"minecraft:operation",
		// 26;
		// 	A particle effect
		// (
		// 		an identifier with extra
		//  	information following
		//  	it for specific particles,
		//  	mirroring the
		// 		Particle packet (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Command_Data#Particle)
		// )
		"minecraft:particle",
		"minecraft:angle", // 27; No comments
		// 28;
		// An angle, represented as 2 numbers
		// (which may have a decimal point,
		// but will be moved to the center of
		// a block if none is specified).
		// May use relative locations with ~.
		"minecraft:rotation",
		// 29;
		// A scoreboard display position slot.
		// "list", "sidebar", "belowName", and
		// "sidebar.team.${color}"" for all chat colors
		// (reset is not included).
		"minecraft:scoreboard_slot",
		"minecraft:score_holder", // 30; See corresponding struct for details.
		// 31;
		// A collection of up to 3 axes.
		"minecraft:swizzle",
		// 32;
		// The name of a team. Parsed as an unquoted string.
		"minecraft:team",
		// 33;
		// A name for an inventory slot.
		"minecraft:item_slot",
		// 34;
		// An Identifier.
		"minecraft:resource_location",
		// 35;
		// A function.
		"minecraft:function",
		// 36;
		// The entity anchor related to the
		// facing argument in the teleport command,
		// is feet or eyes.
		"minecraft:entity_anchor",
		// 37;
		// An integer range of values with a min and a max.
		"minecraft:int_range",
		// 38;
		// A floating-point range of values with a min and a max.
		"minecraft:float_range",
		// 39;
		// Represents a dimension.
		"minecraft:dimension",
		// 40;
		// Represents a gamemode
		// (survival, creative, adventure or spectator).
		"minecraft:gamemode",
		"minecraft:time",                // 41; See corresponding struct for details.
		"minecraft:resource_or_tag",     // 42; See corresponding struct for details.
		"minecraft:resource_or_tag_key", // 43; See corresponding struct for details.
		"minecraft:resource",            // 44; See corresponding struct for details.
		"minecraft:resource_key",        // 45; See corresponding struct for details.
		// 46;
		// Mirror type (none, left_right or front_back).
		"minecraft:template_mirror",
		// 47;
		// Rotation type
		// (none, clockwise_90, 180 or counterclockwise_90).
		"minecraft:template_rotation",
		// 48;
		// Post-worldgen heightmap type
		// (motion_blocking, motion_blocking_no_leaves,
		// ocean_floor and world_surface).
		"minecraft:heightmap",
		// 49;
		// Represents a UUID value.
		"minecraft:uuid",
	},
)

const (
	CommandParserTypeFloat            int32 = 1
	CommandParserTypeDouble           int32 = 2
	CommandParserTypeInteger          int32 = 3
	CommandParserTypeLong             int32 = 4
	CommandParserTypeString           int32 = 5
	CommandParserTypeEntity           int32 = 6
	CommandParserTypeScoreHolder      int32 = 30
	CommandParserTypeTime             int32 = 41
	CommandParserTypeResourceOrTag    int32 = 42
	CommandParserTypeResourceOrTagKey int32 = 43
	CommandParserTypeResource         int32 = 44
	CommandParserTypeResourceKey      int32 = 45
)

// CommandParser ..
type CommandParser interface {
	Name() string
	Marshaler
}

// lookupCommandParserType looks up
// the ID of a CommandParser.
func lookupCommandParserType(x CommandParser, id *int32) bool {
	switch val := x.(type) {
	case *CommandParserFloat:
		*id = CommandParserTypeFloat
	case *CommandParserDouble:
		*id = CommandParserTypeDouble
	case *CommandParserInteger:
		*id = CommandParserTypeInteger
	case *CommandParserLong:
		*id = CommandParserTypeLong
	case *CommandParserString:
		*id = CommandParserTypeString
	case *CommandParserEntity:
		*id = CommandParserTypeEntity
	case *CommandParserScoreHolder:
		*id = CommandParserTypeScoreHolder
	case *CommandParserTime:
		*id = CommandParserTypeTime
	case *CommandParserResourceOrTag:
		*id = CommandParserTypeResourceOrTag
	case *CommandParserResourceOrTagKey:
		*id = CommandParserTypeResourceOrTagKey
	case *CommandParserResource:
		*id = CommandParserTypeResource
	case *CommandParserResourceKey:
		*id = CommandParserTypeResourceKey
	case *CommandParserDefault:
		parserID := val.ID()
		if _, ok := ConstCommandParserEnum.Value(parserID); !ok {
			return false
		}
		*id = parserID
	default:
		return false
	}
	return true
}

// lookupCommandParser looks up the
// CommandParser matching an ID.
func lookupCommandParser(id int32, x *CommandParser) bool {
	switch id {
	case CommandParserTypeFloat:
		*x = &CommandParserFloat{}
	case CommandParserTypeDouble:
		*x = &CommandParserDouble{}
	case CommandParserTypeInteger:
		*x = &CommandParserInteger{}
	case CommandParserTypeLong:
		*x = &CommandParserLong{}
	case CommandParserTypeString:
		*x = &CommandParserString{}
	case CommandParserTypeEntity:
		*x = &CommandParserEntity{}
	case CommandParserTypeScoreHolder:
		*x = &CommandParserScoreHolder{}
	case CommandParserTypeTime:
		*x = &CommandParserTime{}
	case CommandParserTypeResourceOrTag:
		*x = &CommandParserResourceOrTag{}
	case CommandParserTypeResourceOrTagKey:
		*x = &CommandParserResourceOrTagKey{}
	case CommandParserTypeResource:
		*x = &CommandParserResource{}
	case CommandParserTypeResourceKey:
		*x = &CommandParserResourceKey{}
	default:
		parserName, ok := ConstCommandParserEnum.Value(id)
		if !ok {
			return false
		}
		*x = &CommandParserDefault{
			ParserName: parserName,
			ParserID:   id,
		}
	}
	return true
}

// CommandParserDefault refer to the default
// form of a Minecraft Java command parser.
//
// There are a lot of empty command parser struct
// in Minecraft Java protocol.
//
// So, we use CommandParserDefault to
// record the name and ID of these parsers,
// whitout enumerate all of them as struct.
type CommandParserDefault struct {
	// The name of this command parser.
	ParserName string
	// The ID of this command parser.
	ParserID int32
}

func (c *CommandParserDefault) Name() string {
	return c.ParserName
}

func (c *CommandParserDefault) ID() int32 {
	return c.ParserID
}

func (c *CommandParserDefault) Marshal(io IO) {}

const (
	CommandParserFloatFlagHasMin uint8 = 1 << iota
	CommandParserFloatFlagHasMax
)

// Float (Specifies min and max values).
type CommandParserFloat struct {
	// Flags ..
	Flags byte
	// Only if flags contains CommandParserFloatFlagHasMin
	// (flags & 0x01).
	// If not specified, defaults to -Float.MAX_VALUE (≈ 3.4028235E38).
	Min float32
	// Only if flags contains CommandParserFloatFlagHasMax
	// (flags & 0x02).
	// If not specified, defaults to Float.MAX_VALUE (≈ 3.4028235E38).
	Max float32
}

func (c *CommandParserFloat) Name() string {
	return "brigadier:float"
}

func (c *CommandParserFloat) Marshal(io IO) {
	io.Uint8(&c.Flags)
	if c.Flags&CommandParserFloatFlagHasMin != 0 {
		io.Float32(&c.Min)
	}
	if c.Flags&CommandParserFloatFlagHasMax != 0 {
		io.Float32(&c.Max)
	}
}

const (
	CommandParserDoubleFlagHasMin uint8 = 1 << iota
	CommandParserDoubleFlagHasMax
)

// Double (Specifies min and max values).
type CommandParserDouble struct {
	// Flags ..
	Flags byte
	// Only if flags contains CommandParserDoubleFlagHasMin
	// (flags & 0x01).
	// If not specified, defaults to -Double.MAX_VALUE (≈ -1.7976931348623157E307).
	Min float64
	// Only if flags contains CommandParserDoubleFlagHasMax
	// (flags & 0x02).
	// If not specified, defaults to Double.MAX_VALUE (≈ 1.7976931348623157E307)
	Max float64
}

func (c *CommandParserDouble) Name() string {
	return "brigadier:double"
}

func (c *CommandParserDouble) Marshal(io IO) {
	io.Uint8(&c.Flags)
	if c.Flags&CommandParserDoubleFlagHasMin != 0 {
		io.Float64(&c.Min)
	}
	if c.Flags&CommandParserDoubleFlagHasMax != 0 {
		io.Float64(&c.Max)
	}
}

const (
	CommandParserIntegerFlagHasMin uint8 = 1 << iota
	CommandParserIntegerFlagHasMax
)

// Integer (Specifies min and max values).
type CommandParserInteger struct {
	// Flags ..
	Flags byte
	// Only if flags contains CommandParserIntegerFlagHasMin
	// (flags & 0x01).
	// If not specified, defaults to Integer.MIN_VALUE (-2147483648).
	Min int32
	// Only if flags contains CommandParserIntegerFlagHasMax
	// (flags & 0x02).
	// If not specified, defaults to Integer.MAX_VALUE (2147483647).
	Max int32
}

func (c *CommandParserInteger) Name() string {
	return "brigadier:integer"
}

func (c *CommandParserInteger) Marshal(io IO) {
	io.Uint8(&c.Flags)
	if c.Flags&CommandParserIntegerFlagHasMin != 0 {
		io.Int32(&c.Min)
	}
	if c.Flags&CommandParserIntegerFlagHasMax != 0 {
		io.Int32(&c.Max)
	}
}

const (
	CommandParserLongFlagHasMin uint8 = 1 << iota
	CommandParserLongFlagHasMax
)

// Long (Specifies min and max values).
type CommandParserLong struct {
	// Flags ..
	Flags byte
	// Only if flags contains CommandParserLongFlagHasMin
	// (flags & 0x01).
	// If not specified, defaults to Long.MIN_VALUE (−9,223,372,036,854,775,808).
	Min int32
	// Only if flags contains CommandParserLongFlagHasMax
	// (flags & 0x02).
	// If not specified, defaults to Long.MAX_VALUE (9,223,372,036,854,775,807).
	Max int32
}

func (c *CommandParserLong) Name() string {
	return "brigadier:long"
}

func (c *CommandParserLong) Marshal(io IO) {
	io.Uint8(&c.Flags)
	if c.Flags&CommandParserLongFlagHasMin != 0 {
		io.Int32(&c.Min)
	}
	if c.Flags&CommandParserLongFlagHasMax != 0 {
		io.Int32(&c.Max)
	}
}

const (
	// Reads a single word.
	CommandParserBehaviorSingleWord int32 = iota
	// If it starts with a ‘"’, keeps reading until
	// another ‘"’ (allowing escaping with ‘\’).
	// Otherwise behaves the same as CommandParserBehaviorSingleWord.
	CommandParserBehaviorQuotablePhrase
	// Reads the rest of the content after the cursor.
	// Quotes will not be removed.
	CommandParserBehaviorGreedyPhrase
)

// A string.
type CommandParserString struct {
	// See constant enum above.
	Behavior int32
}

func (c *CommandParserString) Name() string {
	return "brigadier:string"
}

func (c *CommandParserString) Marshal(io IO) {
	io.Varint32(&c.Behavior)
}

const (
	// If set, only allows a single entity/player.
	CommandParserEntityFlagSingleEntity byte = 1 << iota
	// If set, only allows players.
	CommandParserEntityFlagPlayers
)

// A selector, player name, or UUID.
type CommandParserEntity struct {
	// See bit mask constant enum above.
	Flags byte
}

func (c *CommandParserEntity) Name() string {
	return "minecraft:entity"
}

func (c *CommandParserEntity) Marshal(io IO) {
	io.Uint8(&c.Flags)
}

// If set, allows multiple.
const CommandParserScoreMultiple byte = 1 << iota

// Something that can join a team.
// Allows selectors and *.
type CommandParserScoreHolder struct {
	// See bit mask constant enum above.
	Flags byte
}

func (c *CommandParserScoreHolder) Name() string {
	return "minecraft:entity"
}

func (c *CommandParserScoreHolder) Marshal(io IO) {
	io.Uint8(&c.Flags)
}

// Represents a time duration.
type CommandParserTime struct {
	// Minimum duration in ticks.
	Min int32
}

func (c *CommandParserTime) Name() string {
	return "minecraft:time"
}

func (c *CommandParserTime) Marshal(io IO) {
	io.Int32(&c.Min)
}

// An identifier or a tag name for a registry.
type CommandParserResourceOrTag struct {
	// The registry from where suggestions will be sourced from.
	Registry Identifier
}

func (c *CommandParserResourceOrTag) Name() string {
	return "minecraft:resource_or_tag"
}

func (c *CommandParserResourceOrTag) Marshal(io IO) {
	io.Identifier(&c.Registry)
}

// An identifier or a tag name for a registry.
type CommandParserResourceOrTagKey struct {
	// The registry from where suggestions will be sourced from.
	Registry Identifier
}

func (c *CommandParserResourceOrTagKey) Name() string {
	return "minecraft:resource_or_tag_key"
}

func (c *CommandParserResourceOrTagKey) Marshal(io IO) {
	io.Identifier(&c.Registry)
}

// An identifier for a registry.
type CommandParserResource struct {
	// The registry from where suggestions will be sourced from.
	Registry Identifier
}

func (c *CommandParserResource) Name() string {
	return "minecraft:resource"
}

func (c *CommandParserResource) Marshal(io IO) {
	io.Identifier(&c.Registry)
}

// An identifier for a registry.
type CommandParserResourceKey struct {
	// The registry from where suggestions will be sourced from.
	Registry Identifier
}

func (c *CommandParserResourceKey) Name() string {
	return "minecraft:resource_key"
}

func (c *CommandParserResourceKey) Marshal(io IO) {
	io.Identifier(&c.Registry)
}
