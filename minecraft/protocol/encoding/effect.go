package encoding

// Effect is the Minecraft effect
// that could apply to the Minecraft
// Java Entity.
type Effect struct {
	// The ID of the effect in the
	// potion effect type registry.
	TypeID int32
	// The duration of the effect.
	Duration int32
}

func (e *Effect) Marshal(io IO) {
	io.Varint32(&e.TypeID)
	io.Varint32(&e.Duration)
}

// PotionEffectDetail ..
type PotionEffectDetail struct {
	// Amplifier ..
	Amplifier int32
	// -1 for infinite.
	Duration int32
	// Produces more translucent particle effects if true.
	Ambient bool
	// Completely hides effect particles if false.
	ShowParticles bool
	// Shows the potion icon in the inventory screen if true.
	ShowIcon bool
	// Used to store the state of the previous potion effect
	// when a stronger one is applied.
	// This guarantees that the weaker one will persist,
	// in case it lasts longer.
	HiddenEffect Optional[*PotionEffectDetail]
}

func (p *PotionEffectDetail) Marshal(io IO) {
	io.Varint32(&p.Amplifier)
	io.Varint32(&p.Duration)
	io.Bool(&p.ShowParticles)
	io.Bool(&p.ShowIcon)
	OptionalPointerMarshaler(io, &p.HiddenEffect)
}

// Describes all the aspects of a potion effect.
type PotionEffect struct {
	// The ID of the effect in the potion effect type registry.
	TypeID int32
	// See PotionEffectDetail for more information.
	Details PotionEffectDetail
}

func (p *PotionEffect) Marshal(io IO) {
	io.Varint32(&p.TypeID)
	Single(io, &p.Details)
}

const (
	ConsumeEffectTypeApplyEffects int32 = iota
	ConsumeEffectTypeRemoveEffects
	ConsumeEffectTypeClearAllEffects
	ConsumeEffectTypeTeleportRandomly
	ConsumeEffectTypePlaySound
)

// ConsumeEffect ..
type ConsumeEffect interface {
	Name() string
	Marshaler
}

// lookupConsumeEffectType looks up
// the ID of a ConsumeEffect.
func lookupConsumeEffectType(x ConsumeEffect, id *int32) bool {
	switch x.(type) {
	case *ConsumeEffectApplyEffects:
		*id = ConsumeEffectTypeApplyEffects
	case *ConsumeEffectRemoveEffects:
		*id = ConsumeEffectTypeRemoveEffects
	case *ConsumeEffectClearAllEffects:
		*id = ConsumeEffectTypeClearAllEffects
	case *ConsumeEffectTeleportRandomly:
		*id = ConsumeEffectTypeTeleportRandomly
	case *ConsumeEffectPlaySound:
		*id = ConsumeEffectTypePlaySound
	default:
		return false
	}
	return true
}

// lookupConsumeEffect looks up the
// ConsumeEffect matching an ID.
func lookupConsumeEffect(id int32, x *ConsumeEffect) bool {
	switch id {
	case ConsumeEffectTypeApplyEffects:
		*x = &ConsumeEffectApplyEffects{}
	case ConsumeEffectTypeRemoveEffects:
		*x = &ConsumeEffectRemoveEffects{}
	case ConsumeEffectTypeClearAllEffects:
		*x = &ConsumeEffectClearAllEffects{}
	case ConsumeEffectTypeTeleportRandomly:
		*x = &ConsumeEffectTeleportRandomly{}
	case ConsumeEffectTypePlaySound:
		*x = &ConsumeEffectPlaySound{}
	default:
		return false
	}
	return true
}

// ConsumeEffectApplyEffects ..
type ConsumeEffectApplyEffects struct {
	// Effects ..
	Effects []PotionEffect
	// Probability ..
	Probability float32
}

func (i *ConsumeEffectApplyEffects) Name() string {
	return "minecraft:apply_effects"
}

func (i *ConsumeEffectApplyEffects) Marshal(io IO) {
	SliceVarint32Length(io, &i.Effects)
	io.Float32(&i.Probability)
}

// ConsumeEffectRemoveEffects ..
type ConsumeEffectRemoveEffects struct {
	// Effects ..
	Effects IDSet
}

func (i *ConsumeEffectRemoveEffects) Name() string {
	return "minecraft:remove_effects"
}

func (i *ConsumeEffectRemoveEffects) Marshal(io IO) {
	Single(io, &i.Effects)
}

// ConsumeEffectClearAllEffects ..
type ConsumeEffectClearAllEffects struct{}

func (i *ConsumeEffectClearAllEffects) Name() string {
	return "minecraft:clear_all_effects"
}

func (i *ConsumeEffectClearAllEffects) Marshal(io IO) {}

// ConsumeEffectTeleportRandomly ..
type ConsumeEffectTeleportRandomly struct {
	// Diameter ..
	Diameter float32
}

func (i *ConsumeEffectTeleportRandomly) Name() string {
	return "minecraft:teleport_randomly"
}

func (i *ConsumeEffectTeleportRandomly) Marshal(io IO) {
	io.Float32(&i.Diameter)
}

// ConsumeEffectTeleportRandomly ..
type ConsumeEffectPlaySound struct {
	// Sound ..
	Sound SoundEvent
}

func (i *ConsumeEffectPlaySound) Name() string {
	return "minecraft:play_sound"
}

func (i *ConsumeEffectPlaySound) Marshal(io IO) {
	io.SoundEvent(&i.Sound)
}
