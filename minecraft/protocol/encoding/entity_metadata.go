package encoding

import (
	"github.com/google/uuid"
)

// ------------------------- General define -------------------------

const EntityMetadataTagEnd byte = 0xff

const (
	EntityDataTypeByte int32 = iota
	EntityDataTypeVarint32
	EntityDataTypeVarint64
	EntityDataTypeFloat32
	EntityDataTypeString
	EntityDataTypeTextCompound
	EntityDataTypeOptionalTextCompound
	EntityDataTypeItemStack
	EntityDataTypeBoolean
	EntityDataTypeRotations
	EntityDataTypePosition
	EntityDataTypeOptionalPosition
	EntityDataTypeDirection
	EntityDataTypeOptionalUUID
	EntityDataTypeBlockState
	EntityDataTypeOptionalBlockState
	EntityDataTypeTagNBT
	EntityDataTypeParticle
	EntityDataTypeParticles
	EntityDataTypeVillagerData
	EntityDataTypeOptionalVarint32
	EntityDataTypePose
	EntityDataTypeCatVariant
	EntityDataTypeWolfVariant
	EntityDataTypeFrogVariant
	EntityDataTypeOptionalGlobalPosition
	EntityDataTypePaintingVariant
	EntityDataTypeSnifferState
	EntityDataTypeArmadilloState
	EntityDataTypeVec3
	EntityDataTypeQuaternion
)

// EntityMetadata refer to the metadata
// of Minecraft Java entity.
type EntityMetadata map[uint8]any

// ------------------------- Entity metadata data type -------------------------

// Rotation on x,
// rotation on y,
// rotation on z.
type EntityDataRotations Rotation

func (e *EntityDataRotations) Marshal(io IO) {
	io.Float32(&e[0])
	io.Float32(&e[1])
	io.Float32(&e[2])
}

// EntityDataRotations ..
type EntityDataOptionalPosition Optional[BlockPos]

func (e *EntityDataOptionalPosition) Marshal(io IO) {
	OptionalFunc(io, (*Optional[BlockPos])(e), io.Position)
}

const (
	DirectionDown int32 = iota
	DirectionUp
	DirectionNorth
	DirectionSouth
	DirectionWest
	DirectionEast
)

// Down = 0,
// Up = 1,
// North = 2,
// South = 3,
// West = 4,
// East = 5
type EntityDataDirection int32

func (e *EntityDataDirection) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

// EntityDataOptionalUUID ..
type EntityDataOptionalUUID Optional[uuid.UUID]

func (e *EntityDataOptionalUUID) Marshal(io IO) {
	OptionalFunc(io, (*Optional[uuid.UUID])(e), io.UUID)
}

// An ID in the block state registry.
type EntityDataBlockState int32

func (e *EntityDataBlockState) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

const OptionalBlockStateAbsent int32 = 0

// 0 for absent (air is unrepresentable);
// otherwise, an ID in the block state registry.
type EntityDataOptionalBlockState int32

func (e *EntityDataOptionalBlockState) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

// ConstVillagerTypesEnum record all the Minecraft Java
// villager types (each name and their id).
// Dump from (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Entity_metadata#Entity_Metadata_Format).
var ConstVillagerTypesEnum = NewMinecraftEnum(
	0,
	[]string{
		"minecraft:desert",  // 0
		"minecraft:jungle",  // 1
		"minecraft:plains",  // 2
		"minecraft:savanna", // 3
		"minecraft:snow",    // 4
		"minecraft:swamp",   // 5
		"minecraft:taiga",   // 6
	},
)

// ConstVillagerProfessionsEnum record all the Minecraft Java
// villager professions (each name and their id).
// Dump from (https://minecraft.wiki/w/Minecraft_Wiki:Projects/wiki.vg_merge/Entity_metadata#Entity_Metadata_Format).
var ConstVillagerProfessionsEnum = NewMinecraftEnum(
	0,
	[]string{
		"minecraft:none",          // 0
		"minecraft:armorer",       // 1
		"minecraft:butcher",       // 2
		"minecraft:cartographer",  // 3
		"minecraft:cleric",        // 4
		"minecraft:farmer",        // 5
		"minecraft:fisherman",     // 6
		"minecraft:fletcher",      // 7
		"minecraft:leatherworker", // 8
		"minecraft:librarian",     // 9
		"minecraft:mason",         // 10
		"minecraft:nitwit",        // 11
		"minecraft:shepherd",      // 12
		"minecraft:toolsmith",     // 13
		"minecraft:weaponsmith",   // 14
	},
)

// EntityDataVillagerData ..
type EntityDataVillagerData [3]int32

// VillagerType return both name and id
// of the villager types that included in
// VillagerData.
// If failed to match, then returned id will
// be -1.
func (e EntityDataVillagerData) VillagerType() (name string, id int32) {
	name, ok := ConstVillagerTypesEnum.Value(e[0])
	if !ok {
		return "", -1
	}
	return name, e[0]
}

// VillagerProfession return both name and id
// of the villager profession that included in
// VillagerData.
// If failed to match, then returned id will
// be -1.
func (e EntityDataVillagerData) VillagerProfession() (name string, id int32) {
	name, ok := ConstVillagerProfessionsEnum.Value(e[1])
	if !ok {
		return "", -1
	}
	return name, e[1]
}

// Level return the level of this villager.
// It is record in the villager data.
func (e EntityDataVillagerData) Level() int32 {
	return e[2]
}

func (e *EntityDataVillagerData) Marshal(io IO) {
	io.Varint32(&e[0])
	io.Varint32(&e[1])
	io.Varint32(&e[2])
}

const OptionalVarint32Absent int32 = 0

// Used for entity IDs.
//   - 0 for absent;
//   - 1 + actual value otherwise.
type EntityDataOptionalVarint32 int32

// GetValue return the real value of a optional
// varint32 data type in entity metadata type.
// If this optional varint32 is not exist,
// then will return 0 and false.
func (e EntityDataOptionalVarint32) GetValue() (result int32, exist bool) {
	if e == 0 {
		return 0, false
	} else {
		return int32(e) - 1, true
	}
}

// SetValue set a optional varint32 data type in entity metadata type.
// If exist is false, then it refer the data record is not exist.
// Otherwise, we will set the record data to input.
func (e *EntityDataOptionalVarint32) SetValue(input int32, exist bool) {
	if !exist {
		*e = 0
	} else {
		*e = EntityDataOptionalVarint32(input + 1)
	}
}

func (e *EntityDataOptionalVarint32) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

const (
	PoseStanding int32 = iota
	PoseFallFlying
	PoseSleeping
	PoseSwimming
	PoseSpinAttack
	PoseSneaking
	PoseLongJumping
	PoseDying
	PoseCroaking
	PoseUsingTongue
	PoseSitting
	PoseRoaring
	PoseSniffing
	PoseEmerging
	PoseDigging
	PoseSliding
	PoseShooting
	PoseInhaling
)

// See enum above for more information.
type EntityDataPose int32

func (e *EntityDataPose) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

// An ID in the minecraft:cat_variant registry.
type EntityDataCatVariant int32

func (e *EntityDataCatVariant) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

// An ID in the minecraft:wolf_variant registry,
// or an inline definition.
type EntityDataWolfVariant IDOrX[WolfVariant]

func (e *EntityDataWolfVariant) Marshal(io IO) {
	IDOrXMarshaler(io, (*IDOrX[WolfVariant])(e))
}

// An ID in the minecraft:frog_variant registry.
type EntityDataForgVariant int32

func (e *EntityDataForgVariant) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

// EntityDataOptionalGlobalPosition ..
type EntityDataOptionalGlobalPosition Optional[GlobalBlockPos]

func (e *EntityDataOptionalGlobalPosition) Marshal(io IO) {
	OptionalMarshaler(io, (*Optional[GlobalBlockPos])(e))
}

// An ID in the minecraft:painting_variant registry,
// or an inline definition.
type EntityDataPaintingVariant IDOrX[PaintingVariant]

func (e *EntityDataPaintingVariant) Marshal(io IO) {
	IDOrXMarshaler(io, (*IDOrX[PaintingVariant])(e))
}

const (
	SnifferStateIdling int32 = iota
	SnifferStateFeelingHappy
	SnifferStateScenting
	SnifferStateSniffing
	SnifferStateSearching
	SnifferStateDigging
	SnifferStateRishing
)

// See enum above for more information.
type EntityDataSnifferState int32

func (e *EntityDataSnifferState) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

const (
	ArmadilloStateIdle int32 = iota
	ArmadilloStateRolling
	ArmadilloStateScared
	ArmadilloStateUnrolling
)

// See enum above for more information.
type EntityDataArmadilloState int32

func (e *EntityDataArmadilloState) Marshal(io IO) {
	io.Varint32((*int32)(e))
}

// ------------------------- End -------------------------
