package typhoon

var(
	BLOCK_REGISTRY = &BlockRegistry{
		counter:    0,
		NameToGuid: make(map[string]uint16),
		GuidToName: make(map[uint16]string),
	}
)

type BlockState struct {

}

type BlockRegistry struct {
	counter uint16
	NameToGuid map[string]uint16
	GuidToName map[uint16]string
}

func (registry *BlockRegistry) GetGuid(name string) uint16 {
	if val, ok := registry.NameToGuid[name]; ok {
		return val
	} else {
		nid := registry.counter
		registry.counter++
		registry.NameToGuid[name] = nid
		registry.GuidToName[nid] = name
		return nid
	}
}

func (registry *BlockRegistry) GetName(guid uint16) string {
	if val, ok := registry.GuidToName[guid]; ok {
		return val
	} else {
		return "minecraft:stone"
	}
}

func (registry *BlockRegistry) GetBlockId(name string, proto Protocol) int {
	if name == "minecraft:air" {
		return 0 << 4 | (0 & 0xF)
	}
	//return int(registry.GetGuid(name)) << 4 | (0 & 0xF)
	return 1 << 4 | (0 & 0xF)
}