package typhoon

import "github.com/seebs/nbt"

type BlockPalette interface {
	GetId(name string) int
	RecoverName(id int) string
	GetSize() int
	GetContent() []string
}

type ChunkBlockPalette struct {
	Map []string
}

type ChunkSection struct {
	Palette BlockPalette
	Blocks  [16 * 16 * 16]int
}

type Chunk struct {
	ChunkX   int32
	ChunkZ   int32
	Sections [16]*ChunkSection
}

type Map struct {
	Spawn     Location
	Dimension Dimension
	Chunks    []*Chunk
}

func (m *Map) SetBlock(x, y, z int, typ string) {
	chunk := m.GetChunk(int32(x)/16, int32(z)/16)
	chunkSection := chunk.GetSection(y / 16)
	chunkSection.SetBlock(x%16, y%16, z%16, typ)
}

func (m *Map) GetChunk(x int32, z int32) *Chunk {
	for _, c := range m.Chunks {
		if c.ChunkX == x && c.ChunkZ == z {
			return c
		}
	}

	c := &Chunk{
		ChunkX: x,
		ChunkZ: z,
	}
	for i := 0; i < 16; i++ {
		c.Sections[i] = &ChunkSection{
			Palette: &ChunkBlockPalette{
				[]string{"minecraft:air"},
			},
		}
	}
	m.Chunks = append(m.Chunks, c)
	return c
}

func (chunk *Chunk) GetSection(y int) *ChunkSection {
	if chunk.Sections[y] != nil {
		return chunk.Sections[y]
	}

	c := &ChunkSection{
		Palette: &ChunkBlockPalette{
			[]string{"minecraft:air"},
		},
	}
	chunk.Sections[y] = c
	return c
}

func (palette *ChunkBlockPalette) GetId(name string) int {
	for i, v := range palette.Map {
		if v == name {
			return i
		}
	}
	palette.Map = append(palette.Map, name)
	return len(palette.Map) - 1
}

func (palette *ChunkBlockPalette) RecoverName(id int) string {
	if id < 0 || id >= len(palette.Map) {
		return "minecraft:air"
	}
	return palette.Map[id]
}

func (palette *ChunkBlockPalette) GetSize() int {
	return len(palette.Map)
}

func (palette *ChunkBlockPalette) GetContent() []string {
	return palette.Map
}

func (section *ChunkSection) SetBlock(x, y, z int, typ string) {
	section.Blocks[y<<8|z<<4|x] = section.Palette.GetId(typ)
}

func (m *Map) SendSpawnChunks(p *Player) {
	sx := int32(m.Spawn.X) / 16
	sz := int32(m.Spawn.Z) / 16

	for x := -4; x < 4; x++ {
		for z := -4; z < 4; z++ {
			c := m.GetChunk(int32(x)+sx, int32(z)+sz)
			biomes := make([]byte, 256)
			packet := &PacketPlayChunkData{
				c.ChunkX,
				c.ChunkZ,
				m.Dimension,
				true,
				false,
				c.Sections[:],
				&biomes,
				make([]nbt.Compound, 0),
			}
			p.WritePacket(packet)
		}
	}
}
