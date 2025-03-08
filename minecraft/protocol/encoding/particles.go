package encoding

import "image/color"

// ConstParticlesEnum record all the Minecraft Java
// particles (each particle name and particle id).
// Dump from (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Particles).
var ConstParticlesEnum = NewMinecraftEnum(
	0,
	[]string{
		"minecraft:angry_villager",                  // 0
		"minecraft:block",                           // 1
		"minecraft:block_marker",                    // 2
		"minecraft:bubble",                          // 3
		"minecraft:cloud",                           // 4
		"minecraft:crit",                            // 5
		"minecraft:damage_indicator",                // 6
		"minecraft:dragon_breath",                   // 7
		"minecraft:dripping_lava",                   // 8
		"minecraft:falling_lava",                    // 9
		"minecraft:landing_lava",                    // 10
		"minecraft:dripping_water",                  // 11
		"minecraft:falling_water",                   // 12
		"minecraft:dust",                            // 13
		"minecraft:dust_color_transition",           // 14
		"minecraft:effect",                          // 15
		"minecraft:elder_guardian",                  // 16
		"minecraft:enchanted_hit",                   // 17
		"minecraft:enchant",                         // 18
		"minecraft:end_rod",                         // 19
		"minecraft:entity_effect",                   // 20
		"minecraft:explosion_emitter",               // 21
		"minecraft:explosion",                       // 22
		"minecraft:gust",                            // 23
		"minecraft:small_gust",                      // 24
		"minecraft:gust_emitter_large",              // 25
		"minecraft:gust_emitter_small",              // 26
		"minecraft:sonic_boom",                      // 27
		"minecraft:falling_dust",                    // 28
		"minecraft:firework",                        // 29
		"minecraft:fishing",                         // 30
		"minecraft:flame",                           // 31
		"minecraft:infested",                        // 32
		"minecraft:cherry_leaves",                   // 33
		"minecraft:pale_oak_leaves",                 // 34
		"minecraft:sculk_soul",                      // 35
		"minecraft:sculk_charge",                    // 36
		"minecraft:sculk_charge_pop",                // 37
		"minecraft:soul_fire_flame",                 // 38
		"minecraft:soul",                            // 39
		"minecraft:flash",                           // 40
		"minecraft:happy_villager",                  // 41
		"minecraft:composter",                       // 42
		"minecraft:heart",                           // 43
		"minecraft:instant_effect",                  // 44
		"minecraft:item",                            // 45
		"minecraft:vibration",                       // 46
		"minecraft:trail",                           // 47
		"minecraft:item_slime",                      // 48
		"minecraft:item_cobweb",                     // 49
		"minecraft:item_snowball",                   // 50
		"minecraft:large_smoke",                     // 51
		"minecraft:lava",                            // 52
		"minecraft:mycelium",                        // 53
		"minecraft:note",                            // 54
		"minecraft:poof",                            // 55
		"minecraft:portal",                          // 56
		"minecraft:rain",                            // 57
		"minecraft:smoke",                           // 58
		"minecraft:white_smoke",                     // 59
		"minecraft:sneeze",                          // 60
		"minecraft:spit",                            // 61
		"minecraft:squid_ink",                       // 62
		"minecraft:sweep_attack",                    // 63
		"minecraft:totem_of_undying",                // 64
		"minecraft:underwater",                      // 65
		"minecraft:splash",                          // 66
		"minecraft:witch",                           // 67
		"minecraft:bubble_pop",                      // 68
		"minecraft:current_down",                    // 69
		"minecraft:bubble_column_up",                // 70
		"minecraft:nautilus",                        // 71
		"minecraft:dolphin",                         // 72
		"minecraft:campfire_cosy_smoke",             // 73
		"minecraft:campfire_signal_smoke",           // 74
		"minecraft:dripping_honey",                  // 75
		"minecraft:falling_honey",                   // 76
		"minecraft:landing_honey",                   // 77
		"minecraft:falling_nectar",                  // 78
		"minecraft:falling_spore_blossom",           // 79
		"minecraft:ash",                             // 80
		"minecraft:crimson_spore",                   // 81
		"minecraft:warped_spore",                    // 82
		"minecraft:spore_blossom_air",               // 83
		"minecraft:dripping_obsidian_tear",          // 84
		"minecraft:falling_obsidian_tear",           // 85
		"minecraft:landing_obsidian_tear",           // 86
		"minecraft:reverse_portal",                  // 87
		"minecraft:white_ash",                       // 88
		"minecraft:small_flame",                     // 89
		"minecraft:snowflake",                       // 90
		"minecraft:dripping_dripstone_lava",         // 91
		"minecraft:falling_dripstone_lava",          // 92
		"minecraft:dripping_dripstone_water",        // 93
		"minecraft:falling_dripstone_water",         // 94
		"minecraft:glow_squid_ink",                  // 95
		"minecraft:glow",                            // 96
		"minecraft:wax_on",                          // 97
		"minecraft:wax_off",                         // 98
		"minecraft:electric_spark",                  // 99
		"minecraft:scrape",                          // 100
		"minecraft:shriek",                          // 101
		"minecraft:egg_crack",                       // 102
		"minecraft:dust_plume",                      // 103
		"minecraft:trial_spawner_detection",         // 104
		"minecraft:trial_spawner_detection_ominous", // 105
		"minecraft:vault_connection",                // 106
		"minecraft:dust_pillar",                     // 107
		"minecraft:ominous_spawning",                // 108
		"minecraft:raid_omen",                       // 109
		"minecraft:trial_omen",                      // 110
		"minecraft:block_crumble",                   // 111
	},
)

const (
	ParticleTypeBlock               int32 = 1
	ParticleTypeBlockMarker         int32 = 2
	ParticleTypeDust                int32 = 13
	ParticleTypeDustColorTransition int32 = 14
	ParticleTypeEntityEffect        int32 = 20
	ParticleTypeFallingDust         int32 = 28
	ParticleTypeSculkCharge         int32 = 36
	ParticleTypeItem                int32 = 45
	ParticleTypeVibration           int32 = 46
	ParticleTypeTrail               int32 = 47
	ParticleTypeShriek              int32 = 101
	ParticleTypeDustPillar          int32 = 107
	ParticleTypeBlockCrumble        int32 = 111
)

// Particle ..
type Particle interface {
	Name() string
	Marshaler
}

// lookupParticleType looks up
// the ID of a Particle.
func lookupParticleType(x Particle, id *int32) bool {
	switch val := x.(type) {
	case *ParticleBlock:
		*id = ParticleTypeBlock
	case *ParticleBlockMarker:
		*id = ParticleTypeBlockMarker
	case *ParticleDust:
		*id = ParticleTypeDust
	case *ParticleDustColorTransition:
		*id = ParticleTypeDustColorTransition
	case *ParticleEntityEffect:
		*id = ParticleTypeEntityEffect
	case *ParticleFallingDust:
		*id = ParticleTypeFallingDust
	case *ParticleSculkCharge:
		*id = ParticleTypeSculkCharge
	case *ParticleItem:
		*id = ParticleTypeItem
	case *ParticleVibration:
		*id = ParticleTypeVibration
	case *ParticleTrail:
		*id = ParticleTypeTrail
	case *ParticleShriek:
		*id = ParticleTypeShriek
	case *ParticleDustPillar:
		*id = ParticleTypeDustPillar
	case *ParticleBlockCrumble:
		*id = ParticleTypeBlockCrumble
	case *ParticleDefault:
		particleID := val.ID()
		if _, ok := ConstParticlesEnum.Value(particleID); !ok {
			return false
		}
		*id = particleID
	default:
		return false
	}
	return true
}

// lookupParticle looks up the
// Particle matching an ID.
func lookupParticle(id int32, x *Particle) bool {
	switch id {
	case ParticleTypeBlock:
		*x = &ParticleBlock{}
	case ParticleTypeBlockMarker:
		*x = &ParticleBlockMarker{}
	case ParticleTypeDust:
		*x = &ParticleDust{}
	case ParticleTypeDustColorTransition:
		*x = &ParticleDustColorTransition{}
	case ParticleTypeEntityEffect:
		*x = &ParticleEntityEffect{}
	case ParticleTypeFallingDust:
		*x = &ParticleFallingDust{}
	case ParticleTypeSculkCharge:
		*x = &ParticleSculkCharge{}
	case ParticleTypeItem:
		*x = &ParticleItem{}
	case ParticleTypeVibration:
		*x = &ParticleVibration{}
	case ParticleTypeTrail:
		*x = &ParticleTrail{}
	case ParticleTypeShriek:
		*x = &ParticleShriek{}
	case ParticleTypeDustPillar:
		*x = &ParticleDustPillar{}
	case ParticleTypeBlockCrumble:
		*x = &ParticleBlockCrumble{}
	default:
		particleName, ok := ConstParticlesEnum.Value(id)
		if !ok {
			return false
		}
		*x = &ParticleDefault{
			ParticleName: particleName,
			ParticleID:   id,
		}
	}
	return true
}

// ParticleDefault refer to the default
// form of a Minecraft Java particle.
//
// There are a lot of particles but encoded
// with nothing in the Minecraft Java
// protocol.
// So that's why we use ParticleDefault to
// record the name and ID of these particles.
type ParticleDefault struct {
	// The name of this particle.
	ParticleName string
	// The ID of this particle.
	ParticleID int32
}

func (p *ParticleDefault) Name() string {
	return p.ParticleName
}

func (p *ParticleDefault) ID() int32 {
	return p.ParticleID
}

func (p *ParticleDefault) Marshal(io IO) {}

// ParticleBlock ..
type ParticleBlock struct {
	// The ID of the block state.
	BlockState int32
}

func (p *ParticleBlock) Name() string {
	return "minecraft:block"
}

func (p *ParticleBlock) Marshal(io IO) {
	io.Varint32(&p.BlockState)
}

// ParticleBlockMarker ..
type ParticleBlockMarker struct {
	// The ID of the block state.
	BlockState int32
}

func (p *ParticleBlockMarker) Name() string {
	return "minecraft:block_marker"
}

func (p *ParticleBlockMarker) Marshal(io IO) {
	io.Varint32(&p.BlockState)
}

// ParticleDust ..
type ParticleDust struct {
	// The start color, encoded as 0xRRGGBB;
	// top bits are ignored.
	Color color.RGBA
	// The scale, will be clamped between 0.01 and 4.
	Scale float32
}

func (p *ParticleDust) Name() string {
	return "minecraft:dust"
}

func (p *ParticleDust) Marshal(io IO) {
	io.RGB(&p.Color)
	io.Float32(&p.Scale)
}

// ParticleDustColorTransition ..
type ParticleDustColorTransition struct {
	// See FromColor for more information.
	FromColor color.RGBA
	// The start color, encoded as 0xRRGGBB;
	// top bits are ignored.
	ToColor color.RGBA
	// The scale, will be clamped between 0.01 and 4.
	Scale float32
}

func (p *ParticleDustColorTransition) Name() string {
	return "minecraft:dust_color_transition"
}

func (p *ParticleDustColorTransition) Marshal(io IO) {
	io.RGB(&p.FromColor)
	io.RGB(&p.ToColor)
	io.Float32(&p.Scale)
}

// ParticleEntityEffect ..
type ParticleEntityEffect struct {
	// The ARGB components of the
	// color encoded as an Int.
	Color color.RGBA
}

func (p *ParticleEntityEffect) Name() string {
	return "minecraft:entity_effect"
}

func (p *ParticleEntityEffect) Marshal(io IO) {
	io.RGBA(&p.Color)
}

// ParticleFallingDust ..
type ParticleFallingDust struct {
	// The ID of the block state.
	BlockState int32
}

func (p *ParticleFallingDust) Name() string {
	return "minecraft:falling_dust"
}

func (p *ParticleFallingDust) Marshal(io IO) {
	io.Varint32(&p.BlockState)
}

// ParticleSculkCharge ..
type ParticleSculkCharge struct {
	// How much the particle will
	// be rotated when displayed.
	Roll float32
}

func (p *ParticleSculkCharge) Name() string {
	return "minecraft:sculk_charge"
}

func (p *ParticleSculkCharge) Marshal(io IO) {
	io.Float32(&p.Roll)
}

// ParticleItem ..
type ParticleItem struct {
	// The item that will be used.
	Item ItemStack
}

func (p *ParticleItem) Name() string {
	return "minecraft:item"
}

func (p *ParticleItem) Marshal(io IO) {
	io.ItemStack(&p.Item)
}

const (
	PositionSourceTypeBlock int32 = iota
	PositionSourceTypeEntity
)

// ParticleVibration ..
type ParticleVibration struct {
	// The type of the vibration source
	// 		- 0 for minecraft:block
	// 		- 1 for minecraft:entity
	PositionSourceType int32
	// The position of the block the vibration originated from.
	// Only present if Position Type is minecraft:block.
	BlockPosition BlockPos
	// The ID of the entity the vibration originated from.
	// Only present if Position Type is minecraft:entity.
	EntityID int32
	// The height of the entity's eye relative to the entity.
	// Only present if Position Type is minecraft:entity.
	EntityEyeHeight float32
	// The amount of ticks it takes for the vibration
	// to travel from its source to its destination.
	Ticks int32
}

func (p *ParticleVibration) Name() string {
	return "minecraft:vibration"
}

func (p *ParticleVibration) Marshal(io IO) {
	io.Varint32(&p.PositionSourceType)
	if p.PositionSourceType == PositionSourceTypeBlock {
		io.Position(&p.BlockPosition)
	} else {
		io.Varint32(&p.EntityID)
		io.Float32(&p.EntityEyeHeight)
	}
	io.Varint32(&p.Ticks)
}

// ParticleTrail ..
type ParticleTrail struct {
	// Target X.
	X float64
	// Target Y.
	Y float64
	// Target Z.
	Z float64
	// The trail color, encoded as 0xRRGGBB;
	// top bits are ignored.
	Color color.RGBA
	// Duration (in ticks?)
	Duration int32
}

func (p *ParticleTrail) Name() string {
	return "minecraft:trail"
}

func (p *ParticleTrail) Marshal(io IO) {
	io.Float64(&p.X)
	io.Float64(&p.Y)
	io.Float64(&p.Z)
	io.RGB(&p.Color)
	io.Varint32(&p.Duration)
}

// ParticleShriek ..
type ParticleShriek struct {
	// The time in ticks before
	// the particle is displayed.
	Delay int32
}

func (p *ParticleShriek) Name() string {
	return "minecraft:shriek"
}

func (p *ParticleShriek) Marshal(io IO) {
	io.Varint32(&p.Delay)
}

// ParticleDustPillar ..
type ParticleDustPillar struct {
	// The ID of the block state.
	BlockState int32
}

func (p *ParticleDustPillar) Name() string {
	return "minecraft:dust_pillar"
}

func (p *ParticleDustPillar) Marshal(io IO) {
	io.Varint32(&p.BlockState)
}

// ParticleBlockCrumble ..
type ParticleBlockCrumble struct {
	// The ID of the block state.
	BlockState int32
}

func (p *ParticleBlockCrumble) Name() string {
	return "minecraft:block_crumble"
}

func (p *ParticleBlockCrumble) Marshal(io IO) {
	io.Varint32(&p.BlockState)
}
