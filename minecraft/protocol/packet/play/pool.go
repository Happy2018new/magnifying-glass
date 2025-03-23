package packet_play

import packet_interface "github.com/Happy2018new/magnifying-glass/minecraft/protocol/packet/interface"

// NewClientPool returns a new pool containing packets sent by a client.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewClientPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{}
}

// NewServerPool returns a new pool containing packets sent by a server.
// Packets may be retrieved from it simply by indexing it with the packet ID.
func NewServerPool() packet_interface.Pool {
	return map[int32]packet_interface.Packet{
		IDClientBoundBundleDelimiter:            &BundleDelimiter{},
		IDClientBoundSpawnEntity:                &SpawnEntity{},
		IDClientBoundSpawnExperienceOrb:         &SpawnExperienceOrb{},
		IDClientBoundEntityAnimation:            &EntityAnimation{},
		IDClientBoundAwardStatistics:            &AwardStatistics{},
		IDClientBoundAcknowledgeBlockChange:     &AcknowledgeBlockChange{},
		IDClientBoundSetBlockDestroyStage:       &SetBlockDestroyStage{},
		IDClientBoundBlockEntityData:            &BlockEntityData{},
		IDClientBoundBlockAction:                &BlockAction{},
		IDClientBoundBlockUpdate:                &BlockUpdate{},
		IDClientBoundBossBar:                    &BossBar{},
		IDClientBoundChangeDifficulty:           &ChangeDifficulty{},
		IDClientBoundChunkBatchFinished:         &ChunkBatchFinished{},
		IDClientBoundChunkBatchStart:            &ChunkBatchStart{},
		IDClientBoundChunkBiomes:                &ChunkBiomes{},
		IDClientBoundClearTitles:                &ClearTitles{},
		IDClientBoundCommandSuggestionsResponse: &CommandSuggestionsResponse{},
		IDClientBoundCommands:                   &Commands{},
		IDClientBoundCloseContainer:             &CloseContainer{},
		IDClientBoundSetContainerContent:        &SetContainerContent{},
		IDClientBoundSetContainerProperty:       &SetContainerProperty{},
	}
}
