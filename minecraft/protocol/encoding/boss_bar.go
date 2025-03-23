package encoding

const (
	BossBarActionTypeAdd int32 = iota
	BossBarActionTypeRemove
	BossBarActionTypeUpdateHealth
	BossBarActionTypeUpdateTitle
	BossBarActionTypeUpdateStyle
	BossBarActionTypeUpdateFlags
)

const (
	BossBarColorPink int32 = iota
	BossBarColorBlue
	BossBarColorRed
	BossBarColorGreen
	BossBarColorYellow
	BossBarColorPurple
	BossBarColorWhite
)

const (
	BossBarDivisionNoDivision int32 = iota
	BossBarDivisionSixNotches
	BossBarDivisionTenNotches
	BossBarDivisionTwelveNotches
	BossBarDivisionTwentyNotches
)

const (
	BossBarFlagShouldDarkenSky uint8 = 1 << iota
	BossBarFlagIsDragonBar
	BossBarFlagCreateFog
)

// BossBarAction ..
type BossBarAction interface {
	Marshaler
}

// BossBarActionType looks up
// the ID of a BossBarAction.
func lookupBossBarActionType(x BossBarAction, id *int32) bool {
	switch x.(type) {
	case *BossBarActionAdd:
		*id = BossBarActionTypeAdd
	case *BossBarActionRemove:
		*id = BossBarActionTypeRemove
	case *BossBarActionUpdateHealth:
		*id = BossBarActionTypeUpdateHealth
	case *BossBarActionUpdateTitle:
		*id = BossBarActionTypeUpdateTitle
	case *BossBarActionUpdateStyle:
		*id = BossBarActionTypeUpdateStyle
	case *BossBarActionUpdateFlags:
		*id = BossBarActionTypeUpdateFlags
	default:
		return false
	}
	return true
}

// lookupBossBarAction looks up the
// BossBarAction matching an ID.
func lookupBossBarAction(id int32, x *BossBarAction) bool {
	switch id {
	case BossBarActionTypeAdd:
		*x = &BossBarActionAdd{}
	case BossBarActionTypeRemove:
		*x = &BossBarActionRemove{}
	case BossBarActionTypeUpdateHealth:
		*x = &BossBarActionUpdateHealth{}
	case BossBarActionTypeUpdateTitle:
		*x = &BossBarActionUpdateTitle{}
	case BossBarActionTypeUpdateStyle:
		*x = &BossBarActionUpdateStyle{}
	case BossBarActionTypeUpdateFlags:
		*x = &BossBarActionUpdateFlags{}
	default:
		return false
	}
	return true
}

// BossBarActionAdd ..
type BossBarActionAdd struct {
	// Effects ..
	Title TextComponent
	// From 0 to 1. Values greater than
	// 1 do not crash a vanilla client,
	// and start rendering part of a second
	// health bar at around 1.5.
	//
	// For more information, see the following links.
	//		- rendering part of a second health bar (https://i.johni0702.de/nA.png)
	Health float32
	// Color ID (See constant enum above).
	Color int32
	// Type of division (See constant enum above).
	Division int32
	// Bit mask.
	// 0x01: should darken sky,
	// 0x02: is dragon bar (used to play end music),
	// 0x04: create fog (previously was also controlled by 0x02).
	Flags byte
}

func (b *BossBarActionAdd) Marshal(io IO) {
	io.TextComponent(&b.Title)
	io.Float32(&b.Health)
	io.Varint32(&b.Color)
	io.Varint32(&b.Division)
	io.Uint8(&b.Flags)
}

// BossBarActionRemove ..
type BossBarActionRemove struct{}

func (b *BossBarActionRemove) Marshal(io IO) {}

// BossBarActionUpdateHealth ..
type BossBarActionUpdateHealth struct {
	// Same as BossBarActionAdd.Health.
	Health float32
}

func (b *BossBarActionUpdateHealth) Marshal(io IO) {
	io.Float32(&b.Health)
}

// BossBarActionUpdateTitle ..
type BossBarActionUpdateTitle struct {
	// Same as BossBarActionAdd.Title.
	Title TextComponent
}

func (b *BossBarActionUpdateTitle) Marshal(io IO) {
	io.TextComponent(&b.Title)
}

// BossBarActionUpdateStyle ..
type BossBarActionUpdateStyle struct {
	// Same as BossBarActionAdd.Color.
	// Color ID (See constant enum above.)
	Color int32
	// Same as BossBarActionAdd.Division.
	Dividers int32
}

func (b *BossBarActionUpdateStyle) Marshal(io IO) {
	io.Varint32(&b.Color)
	io.Varint32(&b.Dividers)
}

// BossBarActionUpdateFlags ..
type BossBarActionUpdateFlags struct {
	// Same as BossBarActionAdd.Flags.
	Flags byte
}

func (b *BossBarActionUpdateFlags) Marshal(io IO) {
	io.Uint8(&b.Flags)
}
