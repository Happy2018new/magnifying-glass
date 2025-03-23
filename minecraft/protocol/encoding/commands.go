package encoding

// 0: root,
// 1: literal,
// 2: argument.
//
// 3 is not used.
//
// NodeTypeRoot is special,
// maybe you need to use this
// to check:
// 		(flag & 1 == 0) && (flag & 2 == 0)
//
// Only the expression above is true,
// then the command node type is root.
const CommandNodeFlagNodeTypeRoot uint8 = 0

const (
	// See comments above for more information.
	CommandNodeFlagNodeTypeLiteral uint8 = 1 << iota
	// See comments above for more information.
	CommandNodeFlagNodeTypeArgument
	// 	Set if the node stack to this
	// point constitutes a valid command.
	CommandNodeFlagIsExecutable
	// Set if the node redirects to another node.
	CommandNodeFlagHasRedirect
	// Only present for argument nodes.
	CommandNodeFlagHasSuggestionsType
)

type CommandNode struct {
	// See constant enum above.
	Flags byte
	// Array of indices of child nodes.
	Children []int32
	// Only if Flags contains CommandNodeFlagHasRedirect
	// (flags & 0x08).
	// Index of redirect node.
	RedirectNode int32
	// Only for argument and literal nodes.
	Name string
	// Only for argument nodes.
	// Varies by parser.
	Parser CommandParser
	// Only if has CommandNodeFlagHasSuggestionsType
	// (flags & 0x10).
	SuggestionsType Identifier
}

func (c *CommandNode) Marshal(io IO) {
	io.Uint8(&c.Flags)
	FuncSliceVarint32Length(io, &c.Children, io.Varint32)
	if c.Flags&CommandNodeFlagHasRedirect != 0 {
		io.Varint32(&c.RedirectNode)
	}
	if c.Flags&CommandNodeFlagNodeTypeLiteral != 0 || c.Flags&CommandNodeFlagNodeTypeArgument != 0 {
		io.String(&c.Name)
	}
	if c.Flags&CommandNodeFlagNodeTypeArgument != 0 {
		io.CommandParserType(&c.Parser)
		c.Parser.Marshal(io)
	}
	if c.Flags&CommandNodeFlagHasSuggestionsType != 0 {
		io.Identifier(&c.SuggestionsType)
	}
}

// CommandSuggestMatch is used on
// Command Suggestions Response packet,
// which used to show command suggestion
// to the client side.
type CommandSuggestMatch struct {
	// One eligible value to insert,
	// note that each command is sent
	// separately instead of in a single
	// string, hence the need for Count.
	//
	// Note that for instance this doesn't
	// include a leading / on commands.
	Insert string
	// Tooltip to display.
	Tooltip Optional[TextComponentComplex]
}

func (c *CommandSuggestMatch) Marshal(io IO) {
	io.String(&c.Insert)
	OptionalFunc(io, &c.Tooltip, io.TextComponentComplex)
}
