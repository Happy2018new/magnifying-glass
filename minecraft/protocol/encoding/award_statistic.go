package encoding

// ConstAwardCategoriesEnum record all the Minecraft Java
// award categories (each name and their id).
// It is refer to the CategoryID in the AwardStatistic struct.
//
// Note that:
// 		- These are namespaced, but with ":" replaced
// 		  with ".".
//		- Blocks, Items, and Entities use block (not block state),
// 		  item, and entity ids.
//
// Dump from (https://minecraft.wiki/w/Java_Edition_protocol#Award_Statistics)
var ConstAwardCategoriesEnum = NewMinecraftEnum(
	0,
	[]string{
		"minecraft.mined",     // 0; Registry: Blocks
		"minecraft.crafted",   // 1; Registry: Items
		"minecraft.used",      // 2; Registry: Items
		"minecraft.broken",    // 3; Registry: Items
		"minecraft.picked_up", // 4; Registry: Items
		"minecraft.dropped",   // 5; Registry: Items
		"minecraft.killed",    // 6; Registry: Entities
		"minecraft.killed_by", // 7; Registry: Entities
		"minecraft.custom",    // 8; Registry: Custom
	},
)

// ConstAwardCustomStatisticEnum record all the Minecraft Java
// custom award (each name and their id).
// It is refer to the StatisticID in the AwardStatistic struct.
//
// Note that uints is refer to:
//		- None: just a normal number (formatted with 0 decimal places)
//		- Damage: value is 10 times the normal amount
//		- Distance: a distance in centimeters (hundredths of blocks)
//		- Time: a time span in ticks
//
// Dump from (https://minecraft.wiki/w/Java_Edition_protocol#Award_Statistics)
var ConstAwardCustomStatisticEnum = NewMinecraftEnum(
	0,
	[]string{
		"minecraft.leave_game",                      // 0; Unit: None
		"minecraft.play_time",                       // 1; Unit: Time
		"minecraft.total_world_time",                // 2; Unit: Time
		"minecraft.time_since_death",                // 3; Unit: Time
		"minecraft.time_since_rest",                 // 4; Unit: Time
		"minecraft.sneak_time",                      // 5; Unit: Time
		"minecraft.walk_one_cm",                     // 6; Unit: Distance
		"minecraft.crouch_one_cm",                   // 7; Unit: Distance
		"minecraft.sprint_one_cm",                   // 8; Unit: Distance
		"minecraft.walk_on_water_one_cm",            // 9; Unit: Distance
		"minecraft.fall_one_cm",                     // 10; Unit: Distance
		"minecraft.climb_one_cm",                    // 11; Unit: Distance
		"minecraft.fly_one_cm",                      // 12; Unit: Distance
		"minecraft.walk_under_water_one_cm",         // 13; Unit: Distance
		"minecraft.minecart_one_cm",                 // 14; Unit: Distance
		"minecraft.boat_one_cm",                     // 15; Unit: Distance
		"minecraft.pig_one_cm",                      // 16; Unit: Distance
		"minecraft.horse_one_cm",                    // 17; Unit: Distance
		"minecraft.aviate_one_cm",                   // 18; Unit: Distance
		"minecraft.swim_one_cm",                     // 19; Unit: Distance
		"minecraft.strider_one_cm",                  // 20; Unit: Distance
		"minecraft.jump",                            // 21; Unit: None
		"minecraft.drop",                            // 22; Unit: None
		"minecraft.damage_dealt",                    // 23; Unit: Damage
		"minecraft.damage_dealt_absorbed",           // 24; Unit: Damage
		"minecraft.damage_dealt_resisted",           // 25; Unit: Damage
		"minecraft.damage_taken",                    // 26; Unit: Damage
		"minecraft.damage_blocked_by_shield",        // 27; Unit: Damage
		"minecraft.damage_absorbed",                 // 28; Unit: Damage
		"minecraft.damage_resisted",                 // 29; Unit: Damage
		"minecraft.deaths",                          // 30; Unit: None
		"minecraft.mob_kills",                       // 31; Unit: None
		"minecraft.animals_bred",                    // 32; Unit: None
		"minecraft.player_kills",                    // 33; Unit: None
		"minecraft.fish_caught",                     // 34; Unit: None
		"minecraft.talked_to_villager",              // 35; Unit: None
		"minecraft.traded_with_villager",            // 36; Unit: None
		"minecraft.eat_cake_slice",                  // 37; Unit: None
		"minecraft.fill_cauldron",                   // 38; Unit: None
		"minecraft.use_cauldron",                    // 39; Unit: None
		"minecraft.clean_armor",                     // 40; Unit: None
		"minecraft.clean_banner",                    // 41; Unit: None
		"minecraft.clean_shulker_box",               // 42; Unit: None
		"minecraft.interact_with_brewingstand",      // 43; Unit: None
		"minecraft.interact_with_beacon",            // 44; Unit: None
		"minecraft.inspect_dropper",                 // 45; Unit: None
		"minecraft.inspect_hopper",                  // 46; Unit: None
		"minecraft.inspect_dispenser",               // 47; Unit: None
		"minecraft.play_noteblock",                  // 48; Unit: None
		"minecraft.tune_noteblock",                  // 49; Unit: None
		"minecraft.pot_flower",                      // 50; Unit: None
		"minecraft.trigger_trapped_chest",           // 51; Unit: None
		"minecraft.open_enderchest",                 // 52; Unit: None
		"minecraft.enchant_item",                    // 53; Unit: None
		"minecraft.play_record",                     // 54; Unit: None
		"minecraft.interact_with_furnace",           // 55; Unit: None
		"minecraft.interact_with_crafting_table",    // 56; Unit: None
		"minecraft.open_chest",                      // 57; Unit: None
		"minecraft.sleep_in_bed",                    // 58; Unit: None
		"minecraft.open_shulker_box",                // 59; Unit: None
		"minecraft.open_barrel",                     // 60; Unit: None
		"minecraft.interact_with_blast_furnace",     // 61; Unit: None
		"minecraft.interact_with_smoker",            // 62; Unit: None
		"minecraft.interact_with_lectern",           // 63; Unit: None
		"minecraft.interact_with_campfire",          // 64; Unit: None
		"minecraft.interact_with_cartography_table", // 65; Unit: None
		"minecraft.interact_with_loom",              // 66; Unit: None
		"minecraft.interact_with_stonecutter",       // 67; Unit: None
		"minecraft.bell_ring",                       // 68; Unit: None
		"minecraft.raid_trigger",                    // 69; Unit: None
		"minecraft.raid_win",                        // 70; Unit: None
		"minecraft.interact_with_anvil",             // 71; Unit: None
		"minecraft.interact_with_grindstone",        // 72; Unit: None
		"minecraft.target_hit",                      // 73; Unit: None
		"minecraft.interact_with_smithing_table",    // 74; Unit: None
	},
)

// AwardStatistic is used in Award Statistics packet,
// which used to statistic the award that client have.
type AwardStatistic struct {
	// See constant enum above.
	CategoryID int32
	// See constant enum above.
	StatisticID int32
	// The amount to set it to.
	Value int32
}

func (a *AwardStatistic) Marshal(io IO) {
	io.Varint32(&a.CategoryID)
	io.Varint32(&a.StatisticID)
	io.Varint32(&a.Value)
}
