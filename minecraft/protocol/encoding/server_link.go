package encoding

const (
	// Displayed on connection error screen;
	// included as a comment in the disconnection report.
	LinkLabelBugReport int32 = iota
	LinkLabelCommunityGuidelines
	LinkLabelSupport
	LinkLabelStatus
	LinkLabelFeedback
	LinkLabelCommunity
	LinkLabelWebsite
	LinkLabelForums
	LinkLabelNews
	LinkLabelAnnouncements
)

// ServerLink is used in Server Link packet,
// which contains a single link that the vanilla
// client will display in the menu available from
// the pause menu. Link labels can be built-in or
// custom (i.e., any text).
type ServerLink struct {
	// True if Label is an enum (built-in label),
	// false if it's a text component (custom label).
	IsBuiltIn bool
	// BuiltInLabel is refer the lable of the link.
	// It is only preset if IsBuiltIn is true,
	// which means this is a built-in label enumerate.
	BuiltInLabel int32
	// CustomLabel is refer to the lable of the link.
	// Different from BuiltInLabelEnum,
	// it is preset if IsBuiltIn is false,
	// and that means this link is not a built-in constant.
	CustomLabel TextComponentComplex
	// Valid URL.
	URL string
}

func (s *ServerLink) Marshal(io IO) {
	io.Bool(&s.IsBuiltIn)
	if s.IsBuiltIn {
		io.Varint32(&s.BuiltInLabel)
	} else {
		io.TextComponentComplex(&s.CustomLabel)
	}
	io.String(&s.URL)
}
