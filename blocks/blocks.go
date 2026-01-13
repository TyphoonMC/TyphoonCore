package blocks

import (
	"embed"
	"encoding/json"
	"log/slog"
	"strconv"
)

//go:embed data/_all_blocks.json
var blocksData embed.FS

// BlockState represents a single block state with its IDs per protocol version
type BlockState struct {
	Name     string         `json:"name"`
	Fallback string         `json:"fallback,omitempty"`
	Ids      map[string]int `json:"ids"` // protocol version string -> block id
}

var (
	allBlocks        []BlockState
	idsByProtocol    map[uint16]map[int]*BlockState // protocol -> id -> block
	blocksByName     map[string]*BlockState         // name -> block
	protocolVersions = []uint16{
		393, 401, 404, // 1.13.x
		477, 480, 485, 490, 498, // 1.14.x
		573, 575, 578, // 1.15.x
		735, 736, 751, 753, 754, // 1.16.x
		755, 756, // 1.17.x
		757, 758, // 1.18.x
		759, 760, 761, 762, // 1.19.x
		763, 764, 765, 766, // 1.20.x
		767, // 1.21
	}
)

func init() {
	loadBlocks()
}

func loadBlocks() {
	data, err := blocksData.ReadFile("data/_all_blocks.json")
	if err != nil {
		slog.Error("failed to load blocks data", "error", err)
		return
	}

	if err := json.Unmarshal(data, &allBlocks); err != nil {
		slog.Error("failed to parse blocks data", "error", err)
		return
	}

	// Build lookup maps
	idsByProtocol = make(map[uint16]map[int]*BlockState)
	blocksByName = make(map[string]*BlockState)

	// Initialize protocol maps
	for _, proto := range protocolVersions {
		idsByProtocol[proto] = make(map[int]*BlockState)
	}

	for i := range allBlocks {
		block := &allBlocks[i]
		blocksByName[block.Name] = block

		for versionStr, id := range block.Ids {
			version, err := strconv.ParseUint(versionStr, 10, 16)
			if err != nil {
				continue
			}
			proto := uint16(version)
			if idsByProtocol[proto] == nil {
				idsByProtocol[proto] = make(map[int]*BlockState)
			}
			idsByProtocol[proto][id] = block
		}
	}

	slog.Debug("loaded blocks", "count", len(allBlocks), "types", len(blocksByName))
}

// GetGlobalMapping returns all blocks by name (compatibility)
func GetGlobalMapping() map[string]*BlockState {
	result := make(map[string]*BlockState)
	for name, block := range blocksByName {
		result[name] = block
	}
	return result
}

// GetNameFromNewBlocks returns the block name for a given protocol and ID
func GetNameFromNewBlocks(protocol uint16, id int) string {
	if protoMap, ok := idsByProtocol[protocol]; ok {
		if block, ok := protoMap[id]; ok {
			return block.Name
		}
	}
	// Try to find in closest protocol version
	for _, proto := range protocolVersions {
		if proto <= protocol {
			if protoMap, ok := idsByProtocol[proto]; ok {
				if block, ok := protoMap[id]; ok {
					return block.Name
				}
			}
		}
	}
	return "minecraft:air"
}

// GetNewBlockFromName returns the block ID for a given name and protocol
func GetNewBlockFromName(name string, protocol uint16) int {
	block, ok := blocksByName[name]
	if !ok {
		return 0 // air
	}
	return block.GetIdForProtocol(protocol)
}

// GetIdForProtocol returns the block ID for a specific protocol version
func (b *BlockState) GetIdForProtocol(protocol uint16) int {
	// Try exact match first
	protoStr := strconv.FormatUint(uint64(protocol), 10)
	if id, ok := b.Ids[protoStr]; ok {
		return id
	}

	// Find the closest lower version
	var bestVersion uint16
	var bestId int
	for versionStr, id := range b.Ids {
		version, err := strconv.ParseUint(versionStr, 10, 16)
		if err != nil {
			continue
		}
		v := uint16(version)
		if v <= protocol && v > bestVersion {
			bestVersion = v
			bestId = id
		}
	}
	return bestId
}

// NewBlock compatibility wrapper
type NewBlock = BlockState

// NewMapping compatibility type
type NewMapping = map[uint16]int
